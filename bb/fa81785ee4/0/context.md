# Session Context

## User Prompts

### Prompt 1

can you review the current opencode implementation against the doc/requirements/agent-integration-checklist.md

### Prompt 2

can you help me understand this one:   3. Major: Decide which format is "native": If export JSON is native per the checklist, store that in NativeData and derive JSONL for Entire's internal use. If JSONL is preferred for
   Entire, update the checklist.

The idea of creating a jsonl format should violate:

**don't:**
- create a "universal transcript format" in the cli
- transform logs to match what the web ui expects
- strip or restructure data to simplify backend processing

### Prompt 3

yes, do you think the checklist doc needs an update so you would have better catched this on the first pass?

### Prompt 4

maybe let's remove the agent specific, since in theory we would add agents there that are implemented already, so not sure there is really a again.

### Prompt 5

remove the example

