import unittest
from unittest.mock import patch, MagicMock
import os
import sys

# Add the parent directory to the sys.path to allow imports from code_whisperer
sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), '..')))

from code_whisperer import (
    get_line_for_comment,
    generate_review_comment,
    main,
)

class TestCodeWhisperer(unittest.TestCase):

    def test_get_line_for_comment(self):
        """Test that the line number is correctly parsed from a diff hunk."""
        patch_with_hunk = '@@ -15,4 +15,5 @@ some context\n+ an added line\n'
        patch_without_hunk = 'Just some text without a hunk header.'
        self.assertEqual(get_line_for_comment(patch_with_hunk), 15)
        self.assertEqual(get_line_for_comment(patch_without_hunk), 1)

    @patch("code_whisperer.requests.post")
    @patch("builtins.open", new_callable=unittest.mock.mock_open, read_data="Review this: {diff_hunk}")
    def test_generate_review_comment_success(self, mock_open, mock_post):
        """Test successful generation of a review comment."""
        mock_response = MagicMock()
        mock_response.status_code = 200
        mock_response.json.return_value = {
            "candidates": [{"content": {"parts": [{"text": "This is a test comment."}]}}]
        }
        mock_post.return_value = mock_response

        comment = generate_review_comment("a-diff-hunk", "a-filename.py")
        self.assertEqual(comment, "This is a test comment.")
        mock_open.assert_called_with(os.path.join(os.path.dirname(__file__), '..', 'prompt.md'), 'r')

    @patch("code_whisperer.requests.post")
    def test_generate_review_comment_api_error(self, mock_post):
        """Test handling of a Gemini API error."""
        mock_response = MagicMock()
        mock_response.status_code = 500
        mock_response.raise_for_status.side_effect = requests.exceptions.HTTPError
        mock_post.return_value = mock_response

        with patch('builtins.open', unittest.mock.mock_open(read_data="prompt")):
            comment = generate_review_comment("a-diff-hunk", "a-filename.py")
            self.assertEqual(comment, "")

    @patch.dict(os.environ, {"GEMINI_API_KEY": "", "GITHUB_TOKEN": "", "REPO": "", "PR_NUMBER": ""})
    def test_main_missing_env_vars(self):
        """Test that main exits if environment variables are missing."""
        with self.assertRaises(SystemExit) as cm:
            main()
        self.assertEqual(cm.exception.code, 1)

    @patch("code_whisperer.get_pr_commit_id", return_value="test_commit_id")
    @patch("code_whisperer.fetch_pull_request_files", return_value=[
        {"filename": "test.py", "status": "modified", "patch": "@@ -1,1 +1,1 @@"}
    ])
    @patch("code_whisperer.generate_review_comment", return_value="A test comment.")
    @patch("code_whisperer.post_inline_comment")
    @patch.dict(os.environ, {"GEMINI_API_KEY": "key", "GITHUB_TOKEN": "token", "REPO": "owner/repo", "PR_NUMBER": "1"})
    def test_main_full_run(self, mock_post_comment, mock_gen_comment, mock_fetch_files, mock_get_commit):
        """Test a full, successful run of the main function."""
        main()
        mock_get_commit.assert_called_once()
        mock_fetch_files.assert_called_once()
        mock_gen_comment.assert_called_with("@@ -1,1 +1,1 @@", "test.py")
        mock_post_comment.assert_called_with("A test comment.", "test.py", "test_commit_id", 1)

if __name__ == '__main__':
    unittest.main()
