import os
import json
import re
import requests
from typing import List, Dict, Optional

try:
    from dotenv import load_dotenv
    load_dotenv() # Load .env file for local development
except ImportError:
    print("[INFO] python-dotenv not found, skipping. This is normal for GitHub Actions.")

# --- Configuration ---
# Populated by GitHub Actions environment or local .env file.
GITHUB_TOKEN = os.getenv("GITHUB_TOKEN")
REPO = os.getenv("REPO")
PR_NUMBER = os.getenv("PR_NUMBER")
GEMINI_API_KEY = os.getenv("GEMINI_API_KEY")
DEBUG = os.getenv("DEBUG", "false").lower() == "true"

HEADERS = {
    "Authorization": f"token {GITHUB_TOKEN}",
    "Accept": "application/vnd.github.v3+json",
}

# --- Helper Functions ---

def debug_log(message: str) -> None:
    """Prints a debug message if DEBUG is enabled."""
    if DEBUG:
        print(f"[DEBUG] {message}")

def get_line_for_comment(patch: str) -> int:
    """
    Parses a diff patch to find the first added line number in the first hunk.
    This provides a more accurate location for the inline comment.
    Defaults to 1 if no hunk header is found.
    """
    match = re.search(r"@@ -\d+,\d+ \+(\d+),\d+ @@", patch)
    return int(match.group(1)) if match else 1

def post_inline_comment(comment: str, filename: str, commit_id: str, line: int) -> None:
    """Posts an inline review comment to the GitHub pull request."""
    url = f"https://api.github.com/repos/{REPO}/pulls/{PR_NUMBER}/comments"
    data = {
        "body": comment,
        "commit_id": commit_id,
        "path": filename,
        "line": line,
    }
    debug_log(f"Posting comment to {url} with data: {json.dumps(data)}")
    try:
        response = requests.post(url, headers=HEADERS, data=json.dumps(data), timeout=10)
        response.raise_for_status() # Raise an exception for bad status codes (4xx or 5xx)
        print(f"[INFO] Successfully posted comment to {filename} on line {line}")
    except requests.exceptions.RequestException as e:
        print(f"[ERROR] Failed to post comment to {filename}: {e}")

def get_pr_commit_id() -> Optional[str]:
    """Gets the commit ID of the head of the pull request."""
    url = f"https://api.github.com/repos/{REPO}/pulls/{PR_NUMBER}"
    try:
        response = requests.get(url, headers=HEADERS, timeout=10)
        response.raise_for_status()
        return response.json()["head"]["sha"]
    except requests.exceptions.RequestException as e:
        print(f"[ERROR] GitHub API error getting PR info: {e}")
        return None

def fetch_pull_request_files() -> List[Dict]:
    """Fetches the list of changed files in the pull request."""
    url = f"https://api.github.com/repos/{REPO}/pulls/{PR_NUMBER}/files"
    print(f"[INFO] Fetching changed files from: {url}")
    try:
        response = requests.get(url, headers=HEADERS, timeout=10)
        response.raise_for_status()
        return response.json()
    except requests.exceptions.RequestException as e:
        print(f"[ERROR] GitHub API error fetching files: {e}")
        return []

def generate_review_comment(diff_hunk: str, filename: str) -> str:
    """Generates a review comment using the Gemini API."""
    prompt_path = os.path.join(os.path.dirname(__file__), "prompt.md")
    try:
        with open(prompt_path, "r") as f:
            prompt_template = f.read()
    except FileNotFoundError:
        print(f"[ERROR] prompt.md not found at {prompt_path}. Please create it.")
        return ""

    prompt = prompt_template.format(filename=filename, diff_hunk=diff_hunk)
    api_url = f"https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key={GEMINI_API_KEY}"
    api_headers = {"Content-Type": "application/json"}
    data = {"contents": [{"parts": [{"text": prompt}]}]}

    debug_log(f"Gemini API URL: {api_url}")
    debug_log(f"Request data: {json.dumps(data)}")

    try:
        response = requests.post(api_url, headers=api_headers, data=json.dumps(data), timeout=30)
        response.raise_for_status()
        result = response.json()
        return result["candidates"][0]["content"]["parts"][0]["text"].strip()
    except requests.exceptions.RequestException as e:
        print(f"[ERROR] Gemini API error: {e}")
    except (KeyError, IndexError) as e:
        print(f"[ERROR] Failed to parse Gemini response: {e}")
        debug_log(f"Full Gemini Response: {response.text}")
    return ""

def main() -> None:
    """Main function to run the review process."""
    if not all([GITHUB_TOKEN, REPO, PR_NUMBER, GEMINI_API_KEY]):
        print("[ERROR] Missing one or more required environment variables.")
        print("[INFO] For local use, create a .env file. For Actions, set repository secrets.")
        return

    commit_id = get_pr_commit_id()
    if not commit_id:
        return

    pr_files = fetch_pull_request_files()
    if not pr_files:
        print("[INFO] No files found in the pull request.")
        return

    for file in pr_files:
        filename = file["filename"]
        status = file["status"]

        if status == "removed":
            print(f"[INFO] Skipping removed file: {filename}")
            continue
        if "patch" not in file:
            print(f"[INFO] Skipping file with no diff: {filename}")
            continue

        diff_hunk = file["patch"]
        comment = generate_review_comment(diff_hunk, filename)
        if comment:
            line_to_comment = get_line_for_comment(diff_hunk)
            post_inline_comment(comment, filename, commit_id, line_to_comment)

if __name__ == "__main__":
    main()
