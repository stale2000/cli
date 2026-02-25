# Session Context

## User Prompts

### Prompt 1

can you do a full review of the changes in this branch?

### Prompt 2

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

### Prompt 3

can you show me the suggestions for 1-3?

### Prompt 4

yes do all three

### Prompt 5

and can you show me these:  Suggestions (5 items):
  - Minor code quality improvements around error handling and documentation

### Prompt 6

can you do me a summary for the changes in the branch? maybe 4-6 sentences

