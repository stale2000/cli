# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Fix: Avoid redundant []rune allocations in UTF-8 truncation

## Context
PR #444 review comments point out that both `TruncateDescription` and the condensation call site allocate `[]rune` unnecessarily — once to check length, then again inside `TruncateRunes`. This adds overhead especially for the common case where no truncation is needed.

## Changes

### 1. `cmd/entire/cli/stringutil/stringutil.go` — `TruncateRunes`
Add a fast path using `utf8.RuneCountInStri...

### Prompt 2

commit this

