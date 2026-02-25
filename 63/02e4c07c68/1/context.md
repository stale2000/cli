# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Test: Skip transcript flush wait for idle/ended sessions

## Context

We fixed a performance bug where `waitForTranscriptFlush` (3s timeout) was called during prepare-commit-msg and post-commit hooks for sessions that had already flushed their transcript (IDLE/ENDED phase). The fix guards `PrepareTranscript` calls with `state.Phase.IsActive()`.

We need a test that verifies: when a session is IDLE or ENDED and PostCommit runs, it completes quickly without hitt...

