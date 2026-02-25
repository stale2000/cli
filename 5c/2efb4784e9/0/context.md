# Session Context

## User Prompts

### Prompt 1

can you review these changes and also compare them against the other open OpenCode related PRs and the official OpenCode docs

### Prompt 2

The plugin stores transcripts at ~/.opencode/sessions/entire/<sanitized-path>/<sessionId>.json. This is a custom location (not OpenCode's native ~/.local/share/opencode/storage/), which is fine since the plugin controls the writes. The Go side mirrors this with SanitizePathForOpenCode(). PR #220 took a different
  approach (reading from OpenCode's native storage), but this PR's approach is more reliable since it controls the format.

so ~/.opencode/sessions/entire/<sanitized-path>/<sessionId>.js...

### Prompt 3

entire isn't using any path outside of the repo yet, right?

### Prompt 4

hmm, but thinking more now: in general session logs are always bound to the path, right? So having them in the repo might be right anyway? so we could just add them to .git/entire-sessions and maybe it's enough to just have them there? How does resume work then, we feed data back into sqlite!?

### Prompt 5

ok, then I agree we should go with os.TempDir()

