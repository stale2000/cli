# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Fix: Separate OpenCode tool detail extraction in summarize.go

## Context

`extractGenericToolDetail` in summarize.go is shared between OpenCode (camelCase `"filePath"`) and Gemini (snake_case `"file_path"`). Adding `"filePath"` to the shared function mixes format concerns from two different agents.

## Changes

### `cmd/entire/cli/summarize/summarize.go`

1. Revert `extractGenericToolDetail` — remove `"filePath"`, keep original Gemini/generic keys
2. Add `extr...

### Prompt 2

can we update the files directly attached to opencode to only have the right filePath format?

