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

### Prompt 9

if we need to set a global env var, where would we do it? (e.g. GIT_CONFIG_GLOBAL - we are seeing some interference)

### Prompt 10

yes, point it at /dev/null

### Prompt 11

I bailed on that, we're fixing that on the other side

### Prompt 12

gemini is fritzing again

### Prompt 13

run #83

### Prompt 14

93 sorry!

### Prompt 15

let's find out for sure

### Prompt 16

2, poll for the file

### Prompt 17

commit this, pull-rebase then push

