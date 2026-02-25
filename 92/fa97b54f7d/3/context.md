# Session Context

## User Prompts

### Prompt 1

can you review this PR and compare it against the cursor docs, also take a look at existing PRs for curso and compare them to. And then please explain to me key design decisions

### Prompt 2

how much work is changing PR uses preToolUse/postToolUse with Task matcher instead of the dedicated subagentStart/subagentStop hooks. Both approaches work, but using the dedicated hooks would be more semantic.

Can you give an overview / estimate?

### Prompt 3

let's do the change and see what it looks like

### Prompt 4

the cursor specific lifecycle.go?

### Prompt 5

ok, so let's ignore this, and then the PR is in general good to merge, right?

### Prompt 6

Looking at this again, i'm confused. What is the difference between:

HookSupport.GetSupportedHooks()
HookHandler.GetHookNames() (guessing this one is replaced by Agent.HookNames())


One returns []string, one returns []HookType

