package agent

// CapabilityDeclarer is implemented by agents that declare their capabilities
// at registration time (e.g., external plugin agents). The As* helper functions
// below use this interface to gate capability access: an agent must both implement
// the optional interface AND declare the capability as true.
//
// Built-in agents (Claude Code, Gemini CLI, etc.) do NOT implement this interface.
// For those agents, the As* helpers fall through to a direct type assertion,
// preserving existing behavior.
type CapabilityDeclarer interface {
	DeclaredCapabilities() DeclaredCaps
}

// DeclaredCaps enumerates the optional interfaces an agent claims to support.
// JSON tags match the external agent protocol schema so external.InfoResponse
// can deserialize directly into this type.
//
// Not every optional interface appears here: built-in-only capabilities that
// have no external-protocol equivalent (SessionBaseDirProvider, ModelExtractor)
// are intentionally excluded — their As* helpers resolve by type assertion
// alone, with no DeclaredCaps gate.
type DeclaredCaps struct {
	Hooks                  bool `json:"hooks"`
	TranscriptAnalyzer     bool `json:"transcript_analyzer"`
	TranscriptPreparer     bool `json:"transcript_preparer"`
	TokenCalculator        bool `json:"token_calculator"`
	CompactTranscript      bool `json:"compact_transcript"`
	TextGenerator          bool `json:"text_generator"`
	HookResponseWriter     bool `json:"hook_response_writer"`
	SubagentAwareExtractor bool `json:"subagent_aware_extractor"`
}

func hasDeclaredCapability(ag Agent, declared func(DeclaredCaps) bool) bool {
	if cd, ok := ag.(CapabilityDeclarer); ok {
		return declared(cd.DeclaredCapabilities())
	}
	return true
}

// AsHookSupport returns the agent as HookSupport if it both implements the
// interface and (for CapabilityDeclarer agents) has declared the capability.
func AsHookSupport(ag Agent) (HookSupport, bool) {
	hs, ok := ag.(HookSupport)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.Hooks }) {
		return nil, false
	}
	return hs, true
}

// AsTranscriptAnalyzer returns the agent as TranscriptAnalyzer if it both
// implements the interface and (for CapabilityDeclarer agents) has declared the capability.
func AsTranscriptAnalyzer(ag Agent) (TranscriptAnalyzer, bool) {
	ta, ok := ag.(TranscriptAnalyzer)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.TranscriptAnalyzer }) {
		return nil, false
	}
	return ta, true
}

// AsTranscriptPreparer returns the agent as TranscriptPreparer if it both
// implements the interface and (for CapabilityDeclarer agents) has declared the capability.
func AsTranscriptPreparer(ag Agent) (TranscriptPreparer, bool) {
	tp, ok := ag.(TranscriptPreparer)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.TranscriptPreparer }) {
		return nil, false
	}
	return tp, true
}

// AsTokenCalculator returns the agent as TokenCalculator if it both
// implements the interface and (for CapabilityDeclarer agents) has declared the capability.
func AsTokenCalculator(ag Agent) (TokenCalculator, bool) {
	tc, ok := ag.(TokenCalculator)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.TokenCalculator }) {
		return nil, false
	}
	return tc, true
}

// AsTextGenerator returns the agent as TextGenerator if it both
// implements the interface and (for CapabilityDeclarer agents) has declared the capability.
func AsTextGenerator(ag Agent) (TextGenerator, bool) {
	tg, ok := ag.(TextGenerator)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.TextGenerator }) {
		return nil, false
	}
	return tg, true
}

// AsTranscriptCompactor returns the agent as TranscriptCompactor if it both
// implements the interface and (for CapabilityDeclarer agents) has declared the capability.
func AsTranscriptCompactor(ag Agent) (TranscriptCompactor, bool) {
	tc, ok := ag.(TranscriptCompactor)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.CompactTranscript }) {
		return nil, false
	}
	return tc, true
}

// AsHookResponseWriter returns the agent as HookResponseWriter if it both
// implements the interface and (for CapabilityDeclarer agents) has declared the capability.
func AsHookResponseWriter(ag Agent) (HookResponseWriter, bool) {
	hrw, ok := ag.(HookResponseWriter)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.HookResponseWriter }) {
		return nil, false
	}
	return hrw, true
}

// AsPromptExtractor returns the agent as PromptExtractor if it both implements
// the interface and (for CapabilityDeclarer agents) has declared TranscriptAnalyzer.
// ExtractPrompts is conceptually part of transcript analysis, so it shares the same
// capability gate — this prevents calling extract-prompts on external agent binaries
// that never declared transcript_analyzer support.
func AsPromptExtractor(ag Agent) (PromptExtractor, bool) {
	pe, ok := ag.(PromptExtractor)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.TranscriptAnalyzer }) {
		return nil, false
	}
	return pe, true
}

// AsSessionBaseDirProvider returns the agent as SessionBaseDirProvider if it implements
// the interface. No capability declaration is needed since this is a built-in-only feature
// (external agents use the agent binary's own session resolution).
func AsSessionBaseDirProvider(ag Agent) (SessionBaseDirProvider, bool) {
	if ag == nil {
		return nil, false
	}
	sbp, ok := ag.(SessionBaseDirProvider)
	if !ok {
		return nil, false
	}
	return sbp, true
}

// AsModelExtractor returns the agent as ModelExtractor if it implements the
// interface. No capability declaration is needed: transcript-based model
// extraction is a built-in-only fallback for agents whose hooks omit the model
// (e.g., Pi). External agents report the model through their own hook protocol.
func AsModelExtractor(ag Agent) (ModelExtractor, bool) {
	if ag == nil {
		return nil, false
	}
	me, ok := ag.(ModelExtractor)
	if !ok {
		return nil, false
	}
	return me, true
}

// AsSubagentAwareExtractor returns the agent as SubagentAwareExtractor if it both
// implements the interface and (for CapabilityDeclarer agents) has declared the capability.
func AsSubagentAwareExtractor(ag Agent) (SubagentAwareExtractor, bool) {
	sae, ok := ag.(SubagentAwareExtractor)
	if !ok || !hasDeclaredCapability(ag, func(caps DeclaredCaps) bool { return caps.SubagentAwareExtractor }) {
		return nil, false
	}
	return sae, true
}

// AsSkillEventExtractor returns the agent as SkillEventExtractor if it implements
// the interface. Skill-event extraction is currently built-in only; external
// agents do not expose this optional interface through declared capabilities.
func AsSkillEventExtractor(ag Agent) (SkillEventExtractor, bool) {
	if ag == nil {
		return nil, false
	}
	see, ok := ag.(SkillEventExtractor)
	return see, ok
}
