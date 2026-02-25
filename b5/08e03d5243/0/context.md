# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Fix: OpenCode mid-turn commits don't create checkpoints

## Context

When OpenCode's agent runs `git commit` mid-turn (e.g., "add a script and commit"), the commit gets an `Entire-Checkpoint` trailer but no checkpoint data is written to `entire/checkpoints/v1`. The root cause: OpenCode's transcript file (`.entire/tmp/<sessionID>.json`) only gets created at turn-end via `opencode export`, but condensation tries to read it at commit time.

Claude Code doesn't have ...

