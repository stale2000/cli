# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Plan: Remove trail enable/disable commands and related functionality

## Context
The user wants to simplify the trails feature by removing the ability to enable/disable trails. Trails should always be active — only `create` and `update` commands remain. The `list` and `show` (default) commands also stay since they're read-only.

## Changes

### 1. `cmd/entire/cli/trail_cmd.go`
- Remove `newTrailEnableCmd()` and `newTrailDisableCmd()` function definitions
- Remo...

### Prompt 2

Create a branch, commit the changes

