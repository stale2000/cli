# Session Context

## User Prompts

### Prompt 1

**ACTION REQUIRED: Spawn a subagent using the Task tool.**

Do NOT review code directly. Instead, immediately call the Task tool with:

```
Task(
  subagent_type: "general-purpose",
  description: "Reviewer checking [feature]",
  prompt: "
    Read and follow the instructions in .claude/agents/reviewer.md

    Requirements folder: 

    Your task:
    1. Read .claude/agents/reviewer.md for your role and process
    2. Read /README.md for requirements context
    3. Read any existing review-NN.md...

### Prompt 2

lets implement the suggestion from copilot in cursor/hooks.go InstallHooks unmarshals the existing .cursor/hooks.json into CursorHooksFile (a fixed struct) and later marshals it back. Any unknown top-level fields or unmodeled hook sections present in a user’s hooks.json will be dropped on write, which can unintentionally delete user/Cursor config. Consider preserving unknown JSON by parsing into a raw map (e.g., map[string]json.RawMessage for the hooks object) and only mutating the specific ho...

### Prompt 3

lets try and make the tests in cmd/entire/cli/agent/cursor/hooks_test.go parallelizable

### Prompt 4

no

