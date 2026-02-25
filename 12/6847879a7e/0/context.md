# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Plan: Auto-delete stale session state files

## Context

Session state files (`.git/entire-sessions/<session-id>.json`) accumulate over time. Sessions that ended more than 24 hours ago are no longer useful but remain on disk. This change makes stale sessions invisible to callers by deleting them transparently during load/list operations.

## Approach

Add an `IsStale()` method to `session.State` and integrate cleanup into the two code paths that read session stat...

### Prompt 2

[Request interrupted by user for tool use]

### Prompt 3

commit

