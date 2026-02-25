# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Detect External Hook Managers and Warn Users

## Context

PR #355 fixed hook installation to use `git rev-parse --git-path hooks`, which correctly resolves `core.hooksPath` (e.g., Husky's `.husky/_/`). However, Husky regenerates `.husky/_/` on every `npm install` (via the `prepare` script), so Entire's hooks get overwritten.

**Goal:** Detect when an external hook manager is in use, warn the user during `entire enable`, and provide copy-paste instructions for a d...

### Prompt 2

commit this

### Prompt 3

[Request interrupted by user for tool use]

### Prompt 4

move to another branch before committing

