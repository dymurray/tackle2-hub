package resource

import (
	"github.com/konveyor/tackle2-hub/internal/model"
	"github.com/konveyor/tackle2-hub/shared/api"
)

// Agent REST resource.
type Agent api.Agent

// With updates the resource with the model.
func (r *Agent) With(m *model.Agent) {
	baseWith(&r.Resource, &m.Model)
	r.Name = m.Name
	r.Description = m.Description
	r.Recipes = make([]Ref, 0, len(m.Recipes))
	for i := range m.Recipes {
		p := &m.Recipes[i]
		r.Recipes = append(r.Recipes, ref(p.ID, p))
	}
	if m.ProviderType != "" || m.URL != "" || m.ProviderModel != "" || m.IdentityID != nil {
		r.ModelConfig = &api.AgentModelConfig{
			ProviderType: m.ProviderType,
			URL:          m.URL,
			Model:        m.ProviderModel,
			Identity:     refPtr(m.IdentityID, m.Identity),
		}
	}
}

// Model builds a model.
func (r *Agent) Model() (m *model.Agent) {
	m = &model.Agent{}
	m.ID = r.ID
	m.Name = r.Name
	m.Description = r.Description
	m.Recipes = make([]model.AgentRecipe, 0, len(r.Recipes))
	for _, rf := range r.Recipes {
		ar := model.AgentRecipe{}
		ar.ID = rf.ID
		m.Recipes = append(m.Recipes, ar)
	}
	if r.ModelConfig != nil {
		m.ProviderType = r.ModelConfig.ProviderType
		m.URL = r.ModelConfig.URL
		m.ProviderModel = r.ModelConfig.Model
		if r.ModelConfig.Identity != nil {
			m.IdentityID = &r.ModelConfig.Identity.ID
		}
	}
	return
}
