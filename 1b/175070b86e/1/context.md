# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Preflight dependency checks for E2E tests

## Context

When E2E tests encounter missing dependencies (tmux, agent binaries, entire CLI), they fail deep inside test execution with unhelpful errors and no artifact reports. The artifact capture is registered as a `t.Cleanup` inside `SetupRepo`, so if something fails before or during setup, there's no diagnostic output.

## Approach

Add preflight checks in `TestMain` (e2e/tests/main_test.go) that verify all requi...

### Prompt 2

aren't we doubling up on what should be one definition?

### Prompt 3

are there lint issues?

### Prompt 4

commit this, then let's fix the lint

### Prompt 5

any other lint failures?

### Prompt 6

can we clean that up please

### Prompt 7

oh wait where are we? we should be working in the cli repo :|

### Prompt 8

yes, reset the e2e-tests repo and commit in cli

