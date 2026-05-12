package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konveyor/tackle2-hub/internal/api/resource"
	"github.com/konveyor/tackle2-hub/internal/model"
	"github.com/konveyor/tackle2-hub/shared/api"
)

// AgentPlanHandler handles AgentPlan resource routes.
type AgentPlanHandler struct {
	BaseHandler
}

// AddRoutes adds routes.
func (h AgentPlanHandler) AddRoutes(e *gin.Engine) {
	routeGroup := e.Group("/")
	routeGroup.Use(Required("agentplans"))
	routeGroup.GET(api.AgentPlanRoute, h.Get)
	routeGroup.GET(api.AgentPlansRoute, h.List)
	routeGroup.GET(api.AgentPlansRoute+"/", h.List)
	routeGroup.POST(api.AgentPlansRoute, h.Create)
	routeGroup.PUT(api.AgentPlanRoute, h.Update)
	routeGroup.DELETE(api.AgentPlanRoute, h.Delete)
}

// Get godoc
// @summary Get an AgentPlan by ID.
// @description Get an AgentPlan by ID.
// @tags agentplans
// @produce json
// @success 200 {object} AgentPlan
// @router /agent-plans/{id} [get]
// @param id path int true "AgentPlan ID"
func (h AgentPlanHandler) Get(ctx *gin.Context) {
	r := AgentPlan{}
	id := h.pk(ctx)
	m := &model.AgentPlan{}
	err := h.DB(ctx).First(m, id).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	r.With(m)

	h.Respond(ctx, http.StatusOK, r)
}

// List godoc
// @summary List all agent plans.
// @description List all agent plans.
// @tags agentplans
// @produce json
// @success 200 {object} []AgentPlan
// @router /agent-plans [get]
func (h AgentPlanHandler) List(ctx *gin.Context) {
	resources := []AgentPlan{}
	var list []model.AgentPlan
	err := h.DB(ctx).Find(&list).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	for i := range list {
		m := &list[i]
		r := AgentPlan{}
		r.With(m)
		resources = append(resources, r)
	}

	h.Respond(ctx, http.StatusOK, resources)
}

// Create godoc
// @summary Create an agent plan.
// @description Create an agent plan.
// @tags agentplans
// @accept json
// @produce json
// @success 201 {object} AgentPlan
// @router /agent-plans [post]
// @param plan body AgentPlan true "AgentPlan data"
func (h AgentPlanHandler) Create(ctx *gin.Context) {
	r := &AgentPlan{}
	err := h.Bind(ctx, r)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	m := r.Model()
	m.CreateUser = h.CurrentUser(ctx)
	err = h.DB(ctx).Create(m).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	r.With(m)

	h.Respond(ctx, http.StatusCreated, r)
}

// Delete godoc
// @summary Delete an agent plan.
// @description Delete an agent plan.
// @tags agentplans
// @success 204
// @router /agent-plans/{id} [delete]
// @param id path int true "AgentPlan ID"
func (h AgentPlanHandler) Delete(ctx *gin.Context) {
	id := h.pk(ctx)
	m := &model.AgentPlan{}
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
// @summary Update an agent plan.
// @description Update an agent plan.
// @tags agentplans
// @accept json
// @success 204
// @router /agent-plans/{id} [put]
// @param id path int true "AgentPlan ID"
// @param plan body AgentPlan true "AgentPlan data"
func (h AgentPlanHandler) Update(ctx *gin.Context) {
	id := h.pk(ctx)
	r := &AgentPlan{}
	err := h.Bind(ctx, r)
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	m := r.Model()
	m.ID = id
	m.UpdateUser = h.CurrentUser(ctx)
	err = h.DB(ctx).Save(m).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}

	h.Status(ctx, http.StatusNoContent)
}

// AgentPlan REST resource.
type AgentPlan = resource.AgentPlan
