# Go Style Guide

This document outlines the coding conventions and style guidelines for this Go project.

## Formatting

- All Go code should be formatted with `gofmt`.

## Naming

- Use camelCase for variable and function names.
- Interface names should end with `er` (e.g., `Reader`, `Writer`).
- Acronyms should be all uppercase (e.g., `HTTP`, `URL`).

## Comments

- All public functions and types should have a comment explaining their purpose.
- Use `//` for line comments and `/* */` for block comments.

## Error Handling

- Errors should be handled explicitly and not ignored.
- Use the `errors` package to create new error types.

## Testing

- All new features should be accompanied by unit tests.
- Use the `testing` package for writing tests.
