package model

// Agent stores an agentic-migration agent configuration.
//
// The four ModelConfig fields (ProviderType, URL, Model, IdentityID) are
// flattened onto the row; the resource converter re-groups them into the
// nested modelConfig object on the wire. The credential (API key) lives on
// the referenced Identity, never inline.
type Agent struct {
	Model
	Name          string `gorm:"uniqueIndex;not null"`
	Description   string
	Recipes       []AgentRecipe `gorm:"many2many:AgentRecipeMap;constraint:OnDelete:CASCADE"`
	ProviderType  string
	URL           string
	ProviderModel string `gorm:"column:model"`
	IdentityID    *uint
	Identity      *Identity
}

// AgentPlan stores the markdown content of an agentic-migration plan.
type AgentPlan struct {
	Model
	Name     string `gorm:"uniqueIndex;not null"`
	Markdown string `gorm:"type:text"`
}

// AgentRecipe stores the raw pallet YAML of an agentic-migration recipe.
type AgentRecipe struct {
	Model
	Name        string `gorm:"uniqueIndex;not null"`
	Description string
	YAML        string `gorm:"column:yaml;type:text"`
}
