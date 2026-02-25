# Session Context

## User Prompts

### Prompt 1

Implement the following plan:

# Plan: Update Droid E2E Tests to Use Anthropic API Key (BYOK)

## Context

Droid E2E tests currently require `FACTORY_API_KEY` for authentication. We're switching to use `ANTHROPIC_API_KEY` via Factory AI's [BYOK feature](https://docs.factory.ai/cli/byok/overview), which configures Droid to call the Anthropic API directly with a custom model entry in `.factory/settings.json`.

## Changes

### 1. `cmd/entire/cli/e2e_test/agent_runner.go`

**Update `NewFactoryAID...

