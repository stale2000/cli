# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Plan: Deprecation warning for `strategy` field in settings

## Context

The `auto-commit` strategy was removed and `manual-commit` is the only strategy. Users may still have `"strategy": "auto-commit"` in their settings files. Because `DisallowUnknownFields()` is used, this causes a hard error — breaking the CLI.

## Changes (~20 lines of production code)

### 1. `cmd/entire/cli/settings/settings.go`

Add deprecated `Strategy` field to struct:
```go
// Deprecat...

### Prompt 2

I do want a warning message more user friendly, and entire doctor should display the warning at the end of the prompt

