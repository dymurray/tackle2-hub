package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konveyor/tackle2-hub/internal/api/resource"
	"github.com/konveyor/tackle2-hub/internal/model"
	"github.com/konveyor/tackle2-hub/shared/api"
	"gorm.io/gorm/clause"
)

// AgentHandler handles Agent resource routes.
//
// An Agent's modelConfig.identity references an Identity that holds the
// LLM provider credential. By convention the referenced Identity should
// have Kind "llm", but this is not enforced by the hub.
type AgentHandler struct {
	BaseHandler
}

// AddRoutes adds routes.
func (h AgentHandler) AddRoutes(e *gin.Engine) {
	routeGroup := e.Group("/")
	routeGroup.Use(Required("agents"), Transaction)
	routeGroup.GET(api.AgentRoute, h.Get)
	routeGroup.GET(api.AgentsRoute, h.List)
	routeGroup.GET(api.AgentsRoute+"/", h.List)
	routeGroup.POST(api.AgentsRoute, h.Create)
	routeGroup.PUT(api.AgentRoute, h.Update)
	routeGroup.DELETE(api.AgentRoute, h.Delete)
}

// Get godoc
// @summary Get an Agent by ID.
// @description Get an Agent by ID.
// @tags agents
// @produce json
// @success 200 {object} Agent
// @router /agents/{id} [get]
// @param id path int true "Agent ID"
func (h AgentHandler) Get(ctx *gin.Context) {
	r := Agent{}
	id := h.pk(ctx)
	m := &model.Agent{}
	db := h.preLoad(h.DB(ctx), clause.Associations)
	err := db.First(m, id).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	r.With(m)

	h.Respond(ctx, http.StatusOK, r)
}

// List godoc
// @summary List all agents.
// @description List all agents.
// @tags agents
// @produce json
// @success 200 {object} []Agent
// @router /agents [get]
func (h AgentHandler) List(ctx *gin.Context) {
	resources := []Agent{}
	var list []model.Agent
	db := h.preLoad(h.DB(ctx), clause.Associations)
	err := db.Find(&list).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	for i := range list {
		m := &list[i]
		r := Agent{}
		r.With(m)
		resources = append(resources, r)
	}

	h.Respond(ctx, http.StatusOK, resources)
}

// Create godoc
// @summary Create an agent.
// @description Create an agent.
// @tags agents
// @accept json
// @produce json
// @success 201 {object} Agent
// @router /agents [post]
// @param agent body Agent true "Agent data"
func (h AgentHandler) Create(ctx *gin.Context) {
	r := &Agent{}
	err := h.Bind(ctx, r)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	m := r.Model()
	m.CreateUser = h.CurrentUser(ctx)
	db := h.DB(ctx)
	err = db.Omit(clause.Associations).Create(m).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = h.DB(ctx).Model(m).Association("Recipes").Replace(m.Recipes)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	db = h.preLoad(h.DB(ctx), clause.Associations)
	err = db.First(m, m.ID).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	r.With(m)

	h.Respond(ctx, http.StatusCreated, r)
}

// Delete godoc
// @summary Delete an agent.
// @description Delete an agent.
// @tags agents
// @success 204
// @router /agents/{id} [delete]
// @param id path int true "Agent ID"
func (h AgentHandler) Delete(ctx *gin.Context) {
	id := h.pk(ctx)
	m := &model.Agent{}
	db := h.DB(ctx)
	err := db.First(m, id).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = db.Delete(m).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	h.Status(ctx, http.StatusNoContent)
}

// Update godoc
// @summary Update an agent.
// @description Update an agent.
// @tags agents
// @accept json
// @success 204
// @router /agents/{id} [put]
// @param id path int true "Agent ID"
// @param agent body Agent true "Agent data"
func (h AgentHandler) Update(ctx *gin.Context) {
	id := h.pk(ctx)
	r := &Agent{}
	err := h.Bind(ctx, r)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	m := r.Model()
	m.ID = id
	m.UpdateUser = h.CurrentUser(ctx)
	db := h.DB(ctx).Model(m)
	db = db.Omit(clause.Associations)
	err = db.Save(m).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	err = h.DB(ctx).Model(m).Association("Recipes").Replace(m.Recipes)
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	h.Status(ctx, http.StatusNoContent)
}

// Agent REST resource.
type Agent = resource.Agent
