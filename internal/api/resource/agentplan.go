package resource

import (
	"github.com/konveyor/tackle2-hub/internal/model"
	"github.com/konveyor/tackle2-hub/shared/api"
)

// AgentPlan REST resource.
type AgentPlan api.AgentPlan

// With updates the resource with the model.
func (r *AgentPlan) With(m *model.AgentPlan) {
	baseWith(&r.Resource, &m.Model)
	r.Name = m.Name
	r.Markdown = m.Markdown
}

// Model builds a model.
func (r *AgentPlan) Model() (m *model.AgentPlan) {
	m = &model.AgentPlan{
		Name:     r.Name,
		Markdown: r.Markdown,
	}
	m.ID = r.ID
	return
}
