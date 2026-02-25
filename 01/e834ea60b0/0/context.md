# Session Context

## User Prompts

### Prompt 1

Base directory for this skill: /Users/alex/workspace/cli/.claude/skills/debug-e2e

# Debug Entire CLI via E2E Artifacts

Diagnose Entire CLI bugs using captured artifacts from the E2E test suite. Artifacts are written to `e2e/artifacts/` locally or downloaded from CI via GitHub Actions.

## Inputs

The user provides either:
- **A test run directory:** `e2e/artifacts/{timestamp}/` — triage all failures
- **A specific test directory:** `e2e/artifacts/{timestamp}/{TestName}-{agent}/` — debug one...

### Prompt 2

[Request interrupted by user for tool use]

### Prompt 3

try again now the report should be there

### Prompt 4

[Request interrupted by user]

### Prompt 5

you should know I've just run these against a branch build of the cli

the auto commit failure is expected

### Prompt 6

have a look at the test.json - was there any clues in the pane output?

### Prompt 7

in interactive mode? yes it would

Q: does this test use entire 'manual' mode?

### Prompt 8

yes, do both - after each WaitFor and in cleanup

### Prompt 9

commit this

