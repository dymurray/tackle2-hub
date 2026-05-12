package api

// Agent REST resource.
type Agent struct {
	Resource    `yaml:",inline"`
	Name        string            `json:"name" binding:"required"`
	Description string            `json:"description,omitempty" yaml:",omitempty"`
	Recipes     []Ref             `json:"recipes"`
	ModelConfig *AgentModelConfig `json:"modelConfig,omitempty" yaml:"modelConfig,omitempty"`
}

// AgentModelConfig is the model-provider configuration for an Agent.
// The credential (API key) lives on the referenced Identity, not inline.
// By convention, the Identity Kind for LLM credentials is "llm" — this is
// not enforced by the hub since Identity.Kind is a free-form string.
type AgentModelConfig struct {
	ProviderType string `json:"provider_type,omitempty" yaml:"provider_type,omitempty"`
	URL          string `json:"url,omitempty" yaml:"url,omitempty"`
	Model        string `json:"model,omitempty" yaml:"model,omitempty"`
	Identity     *Ref   `json:"identity,omitempty" yaml:"identity,omitempty"`
}

// AgentPlan REST resource.
type AgentPlan struct {
	Resource `yaml:",inline"`
	Name     string `json:"name" binding:"required"`
	Markdown string `json:"markdown" binding:"required"`
}

// AgentRecipe REST resource.
type AgentRecipe struct {
	Resource    `yaml:",inline"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description,omitempty" yaml:",omitempty"`
	YAML        string `json:"yaml" binding:"required"`
}
