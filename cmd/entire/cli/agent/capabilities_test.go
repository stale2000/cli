package agent

import (
	"context"
	"io"
	"testing"

	"github.com/entireio/cli/cmd/entire/cli/agent/types"
)

// --- Mock types for testing ---

// mockBaseAgent implements only agent.Agent (no optional interfaces).
type mockBaseAgent struct{}

func (m *mockBaseAgent) Name() types.AgentName                        { return "mock" }
func (m *mockBaseAgent) Type() types.AgentType                        { return "Mock" }
func (m *mockBaseAgent) Description() string                          { return "mock agent" }
func (m *mockBaseAgent) IsPreview() bool                              { return false }
func (m *mockBaseAgent) DetectPresence(context.Context) (bool, error) { return false, nil }
func (m *mockBaseAgent) ProtectedDirs() []string                      { return nil }
func (m *mockBaseAgent) ReadTranscript(string) ([]byte, error)        { return nil, nil }
func (m *mockBaseAgent) ChunkTranscript(context.Context, []byte, int) ([][]byte, error) {
	return nil, nil
}
func (m *mockBaseAgent) ReassembleTranscript([][]byte) ([]byte, error)     { return nil, nil }
func (m *mockBaseAgent) GetSessionID(*HookInput) string                    { return "" }
func (m *mockBaseAgent) GetSessionDir(string) (string, error)              { return "", nil }
func (m *mockBaseAgent) ResolveSessionFile(string, string) string          { return "" }
func (m *mockBaseAgent) ReadSession(*HookInput) (*AgentSession, error)     { return nil, nil } //nolint:nilnil // test mock
func (m *mockBaseAgent) WriteSession(context.Context, *AgentSession) error { return nil }
func (m *mockBaseAgent) FormatResumeCommand(string) string                 { return "" }

// mockBuiltinHookAgent is a built-in agent that implements HookSupport but NOT CapabilityDeclarer.
type mockBuiltinHookAgent struct {
	mockBaseAgent
}

func (m *mockBuiltinHookAgent) HookNames() []string { return nil }
func (m *mockBuiltinHookAgent) ParseHookEvent(context.Context, string, io.Reader) (*Event, error) {
	return nil, nil //nolint:nilnil // test mock
}
func (m *mockBuiltinHookAgent) InstallHooks(context.Context, bool, bool) (int, error) {
	return 0, nil
}
func (m *mockBuiltinHookAgent) UninstallHooks(context.Context) error   { return nil }
func (m *mockBuiltinHookAgent) AreHooksInstalled(context.Context) bool { return false }

// mockFullAgent implements all optional interfaces AND CapabilityDeclarer.
type mockFullAgent struct {
	mockBaseAgent

	caps DeclaredCaps
}

func (m *mockFullAgent) DeclaredCapabilities() DeclaredCaps { return m.caps }

// HookSupport
func (m *mockFullAgent) HookNames() []string { return nil }
func (m *mockFullAgent) ParseHookEvent(context.Context, string, io.Reader) (*Event, error) {
	return nil, nil //nolint:nilnil // test mock
}
func (m *mockFullAgent) InstallHooks(context.Context, bool, bool) (int, error) { return 0, nil }
func (m *mockFullAgent) UninstallHooks(context.Context) error                  { return nil }
func (m *mockFullAgent) AreHooksInstalled(context.Context) bool                { return false }

// TranscriptAnalyzer
func (m *mockFullAgent) GetTranscriptPosition(string) (int, error) { return 0, nil }
func (m *mockFullAgent) ExtractModifiedFilesFromOffset(string, int) ([]string, int, error) {
	return nil, 0, nil
}
func (m *mockFullAgent) ExtractPrompts(string, int) ([]string, error) { return nil, nil }
func (m *mockFullAgent) ExtractSummary(string) (string, error)        { return "", nil }

// TranscriptPreparer
func (m *mockFullAgent) PrepareTranscript(context.Context, string) error { return nil }

// TokenCalculator
func (m *mockFullAgent) CalculateTokenUsage([]byte, int) (*TokenUsage, error) { return nil, nil } //nolint:nilnil // test mock

// ModelExtractor
func (m *mockFullAgent) ExtractModel([]byte) (string, error) { return "mock-model", nil }

// TextGenerator
func (m *mockFullAgent) GenerateText(context.Context, string, string) (string, error) {
	return "", nil
}

// TranscriptCompactor
func (m *mockFullAgent) CompactTranscript(context.Context, string) (*CompactedTranscript, error) {
	return &CompactedTranscript{}, nil
}

// HookResponseWriter
func (m *mockFullAgent) WriteHookResponse(string) error { return nil }

// SubagentAwareExtractor
func (m *mockFullAgent) ExtractAllModifiedFiles([]byte, int, string) ([]string, error) {
	return nil, nil
}
func (m *mockFullAgent) CalculateTotalTokenUsage([]byte, int, string) (*TokenUsage, error) {
	return nil, nil //nolint:nilnil // test mock
}

// mockBuiltinPromptAgent is a built-in agent that implements PromptExtractor but NOT CapabilityDeclarer.
type mockBuiltinPromptAgent struct {
	mockBaseAgent
}

func (m *mockBuiltinPromptAgent) ExtractPrompts(string, int) ([]string, error) {
	return []string{"test prompt"}, nil
}

// --- Tests ---

func TestDeclaredCapabilityAdapters(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name   string
		enable func(*DeclaredCaps)
		as     func(Agent) (any, bool)
	}{
		{
			name:   "hooks",
			enable: func(caps *DeclaredCaps) { caps.Hooks = true },
			as: func(ag Agent) (any, bool) {
				return AsHookSupport(ag)
			},
		},
		{
			name:   "transcript analyzer",
			enable: func(caps *DeclaredCaps) { caps.TranscriptAnalyzer = true },
			as: func(ag Agent) (any, bool) {
				return AsTranscriptAnalyzer(ag)
			},
		},
		{
			name:   "prompt extractor",
			enable: func(caps *DeclaredCaps) { caps.TranscriptAnalyzer = true },
			as: func(ag Agent) (any, bool) {
				return AsPromptExtractor(ag)
			},
		},
		{
			name:   "transcript preparer",
			enable: func(caps *DeclaredCaps) { caps.TranscriptPreparer = true },
			as: func(ag Agent) (any, bool) {
				return AsTranscriptPreparer(ag)
			},
		},
		{
			name:   "token calculator",
			enable: func(caps *DeclaredCaps) { caps.TokenCalculator = true },
			as: func(ag Agent) (any, bool) {
				return AsTokenCalculator(ag)
			},
		},
		{
			name:   "text generator",
			enable: func(caps *DeclaredCaps) { caps.TextGenerator = true },
			as: func(ag Agent) (any, bool) {
				return AsTextGenerator(ag)
			},
		},
		{
			name:   "transcript compactor",
			enable: func(caps *DeclaredCaps) { caps.CompactTranscript = true },
			as: func(ag Agent) (any, bool) {
				return AsTranscriptCompactor(ag)
			},
		},
		{
			name:   "hook response writer",
			enable: func(caps *DeclaredCaps) { caps.HookResponseWriter = true },
			as: func(ag Agent) (any, bool) {
				return AsHookResponseWriter(ag)
			},
		},
		{
			name:   "subagent-aware extractor",
			enable: func(caps *DeclaredCaps) { caps.SubagentAwareExtractor = true },
			as: func(ag Agent) (any, bool) {
				return AsSubagentAwareExtractor(ag)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			t.Run("nil agent", func(t *testing.T) {
				t.Parallel()
				if capability, ok := tt.as(nil); ok || capability != nil {
					t.Fatalf("got (%T, %v), want nil false", capability, ok)
				}
			})

			t.Run("not implemented", func(t *testing.T) {
				t.Parallel()
				if capability, ok := tt.as(&mockBaseAgent{}); ok || capability != nil {
					t.Fatalf("got (%T, %v), want nil false", capability, ok)
				}
			})

			t.Run("declared", func(t *testing.T) {
				t.Parallel()
				var caps DeclaredCaps
				tt.enable(&caps)
				capability, ok := tt.as(&mockFullAgent{caps: caps})
				if !ok || capability == nil {
					t.Fatalf("got (%T, %v), want capability true", capability, ok)
				}
			})

			t.Run("undeclared", func(t *testing.T) {
				t.Parallel()
				if capability, ok := tt.as(&mockFullAgent{}); ok || capability != nil {
					t.Fatalf("got (%T, %v), want nil false", capability, ok)
				}
			})
		})
	}
}

func TestDeclaredCapabilityAdapters_BuiltinsBypassDeclarations(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name string
		ag   Agent
		as   func(Agent) (any, bool)
	}{
		{
			name: "hook support",
			ag:   &mockBuiltinHookAgent{},
			as: func(ag Agent) (any, bool) {
				return AsHookSupport(ag)
			},
		},
		{
			name: "prompt extractor",
			ag:   &mockBuiltinPromptAgent{},
			as: func(ag Agent) (any, bool) {
				return AsPromptExtractor(ag)
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			capability, ok := tt.as(tt.ag)
			if !ok || capability == nil {
				t.Fatalf("got (%T, %v), want capability true", capability, ok)
			}
		})
	}
}

func TestAsModelExtractor(t *testing.T) {
	t.Parallel()

	// AsModelExtractor is an ungated, built-in-only helper (no DeclaredCaps
	// field): implementing the interface is sufficient.
	t.Run("not implemented", func(t *testing.T) {
		t.Parallel()
		_, ok := AsModelExtractor(&mockBaseAgent{})
		if ok {
			t.Error("expected false")
		}
	})

	t.Run("implemented", func(t *testing.T) {
		t.Parallel()
		me, ok := AsModelExtractor(&mockFullAgent{})
		if !ok || me == nil {
			t.Error("expected true")
		}
	})

	t.Run("nil agent", func(t *testing.T) {
		t.Parallel()
		_, ok := AsModelExtractor(nil)
		if ok {
			t.Error("expected false")
		}
	})
}
