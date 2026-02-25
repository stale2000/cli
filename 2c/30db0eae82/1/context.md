# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Plan: Add Cursor Agent Support on New Agent Interface

## Context

The `soph/agent-refactor` branch refactored the agent system from a callback-driven model (agents call framework functions) to a dispatcher-driven model (agents are passive data providers, framework calls agent methods). PR #260 attempted to add Cursor support on the *old* interface. We need to implement Cursor support on the *new* interface.

The new architecture has a single `DispatchLifecycleEv...

### Prompt 2

commit this and create a draft pr

