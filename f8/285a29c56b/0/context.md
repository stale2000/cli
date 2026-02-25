# Session Context

## User Prompts

### Prompt 1

is Agent.GetHookConfigPath used anywhere?

### Prompt 2

lets remove it

### Prompt 3

commit and create a pr referencing issue 424

### Prompt 4

now lets remove HookHandler interface, which only contains GetHookNames(). Call sites should call Agent.HookNames() instead. Check that this refactor does not cause any changes.

### Prompt 5

commit as 'agent: remove HookHandler'

### Prompt 6

what is the difference between HookSupport.GetSupportedHooks and Agent.HookNames

### Prompt 7

lets remove HookSuport.GetSupportedHooks

