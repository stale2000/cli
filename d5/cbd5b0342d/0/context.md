# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Replace SQLite access with `opencode session delete`

## Context

`sqlite.go` runs raw SQL against OpenCode's database. OpenCode has `opencode session delete <id>` for this. The current SQLite approach hardcodes the DB path, schema, and requires `sqlite3` installed.

## Changes

### `sqlite.go`

Delete everything except `runOpenCodeImport()`. Replace the 5 SQLite functions with one:

```go
func runOpenCodeSessionDelete(sessionID string) error {
    // runs: openc...

