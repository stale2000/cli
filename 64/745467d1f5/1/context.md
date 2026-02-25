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

[Request interrupted by user]

### Prompt 3

**ACTION REQUIRED: Spawn a subagent using the Task tool.**

Do NOT review code directly. Instead, immediately call the Task tool with:

```
Task(
  subagent_type: "general-purpose",
  description: "Reviewer checking [feature]",
  prompt: "
    Read and follow the instructions in .claude/agents/reviewer.md

    Requirements folder: 367

    Your task:
    1. Read .claude/agents/reviewer.md for your role and process
    2. Read 367/README.md for requirements context
    3. Read any existing review...

### Prompt 4

there has been another commit since, what do you think now?

### Prompt 5

[Request interrupted by user]

