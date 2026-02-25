# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Plan: Handle `droid exec` missing UserPromptSubmit

## Context

`droid exec` fires `SessionStart → Stop → SessionEnd` without `UserPromptSubmit`. Since `UserPromptSubmit` maps to `TurnStart` (which creates session state, transitions IDLE→ACTIVE, captures pre-prompt state), the session lifecycle is broken:

- **Phase stays IDLE** — PrepareCommitMsg/PostCommit hooks won't properly link mid-exec commits (Droid CAN auto-commit with `--auto medium/high`)
- **No ses...

