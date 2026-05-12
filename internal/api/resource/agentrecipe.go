package resource

import (
	"github.com/konveyor/tackle2-hub/internal/model"
	"github.com/konveyor/tackle2-hub/shared/api"
)

// AgentRecipe REST resource.
type AgentRecipe api.AgentRecipe

// With updates the resource with the model.
func (r *AgentRecipe) With(m *model.AgentRecipe) {
	baseWith(&r.Resource, &m.Model)
	r.Name = m.Name
	r.Description = m.Description
	r.YAML = m.YAML
}

// Model builds a model.
func (r *AgentRecipe) Model() (m *model.AgentRecipe) {
	m = &model.AgentRecipe{
		Name:        r.Name,
		Description: r.Description,
		YAML:        r.YAML,
	}
	m.ID = r.ID
	return
}
