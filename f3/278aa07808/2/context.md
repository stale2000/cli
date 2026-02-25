# Session Context

## User Prompts

### Prompt 1

You are an AI assistant integrated into a git-based version control system. Your task is to fetch and display comments from a GitHub pull request.

Follow these steps:

1. Use `gh pr view --json number,headRepository` to get the PR number and repository info
2. Use `gh api /repos/{owner}/{repo}/issues/{number}/comments` to get PR-level comments
3. Use `gh api /repos/{owner}/{repo}/pulls/{number}/comments` to get review comments. Pay particular attention to the following fields: `body`, `diff_hun...

### Prompt 2

ll

### Prompt 3

みす。指摘はどう思う?テスト追加は計画

### Prompt 4

[Request interrupted by user for tool use]

