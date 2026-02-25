# Session Context

## User Prompts

### Prompt 1

The stale session cleanup logic is only implemented in this package-level LoadSessionState function, but not in StateStore.Load which is called directly by ManualCommitStrategy.loadSessionState(). This creates an inconsistency where stale sessions are cleaned up in some code paths but not others.

Consider adding the stale session check to StateStore.Load() in session/state.go instead, so all callers automatically benefit from stale session cleanup. This would make the behavior consistent across...

### Prompt 2

[Request interrupted by user for tool use]

### Prompt 3

Commit with thorough message.

