package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/konveyor/tackle2-hub/internal/api/resource"
	"github.com/konveyor/tackle2-hub/internal/model"
	"github.com/konveyor/tackle2-hub/shared/api"
)

// AgentRecipeHandler handles AgentRecipe resource routes.
type AgentRecipeHandler struct {
	BaseHandler
}

// AddRoutes adds routes.
func (h AgentRecipeHandler) AddRoutes(e *gin.Engine) {
	routeGroup := e.Group("/")
	routeGroup.Use(Required("agentrecipes"))
	routeGroup.GET(api.AgentRecipeRoute, h.Get)
	routeGroup.GET(api.AgentRecipesRoute, h.List)
	routeGroup.GET(api.AgentRecipesRoute+"/", h.List)
	routeGroup.POST(api.AgentRecipesRoute, h.Create)
	routeGroup.PUT(api.AgentRecipeRoute, h.Update)
	routeGroup.DELETE(api.AgentRecipeRoute, h.Delete)
}

// Get godoc
// @summary Get an AgentRecipe by ID.
// @description Get an AgentRecipe by ID.
// @tags agentrecipes
// @produce json
// @success 200 {object} AgentRecipe
// @router /agent-recipes/{id} [get]
// @param id path int true "AgentRecipe ID"
func (h AgentRecipeHandler) Get(ctx *gin.Context) {
	r := AgentRecipe{}
	id := h.pk(ctx)
	m := &model.AgentRecipe{}
	err := h.DB(ctx).First(m, id).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	r.With(m)

	h.Respond(ctx, http.StatusOK, r)
}

// List godoc
// @summary List all agent recipes.
// @description List all agent recipes.
// @tags agentrecipes
// @produce json
// @success 200 {object} []AgentRecipe
// @router /agent-recipes [get]
func (h AgentRecipeHandler) List(ctx *gin.Context) {
	resources := []AgentRecipe{}
	var list []model.AgentRecipe
	err := h.DB(ctx).Find(&list).Error
	if err != nil {
		_ = ctx.Error(err)
		return
	}
	for i := range list {
		m := &list[i]
		r := AgentRecipe{}
		r.With(m)
		resources = append(resources, r)
	}

	h.Respond(ctx, http.StatusOK, resources)
}

// Create godoc
// @summary Create an agent recipe.
// @description Create an agent recipe.
// @tags agentrecipes
// @accept json
// @produce json
// @success 201 {object} AgentRecipe
// @router /agent-recipes [post]
// @param recipe body AgentRecipe true "AgentRecipe data"
func (h AgentRecipeHandler) Create(ctx *gin.Context) {
	r := &AgentRecipe{}
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
// @summary Delete an agent recipe.
// @description Delete an agent recipe.
// @tags agentrecipes
// @success 204
// @router /agent-recipes/{id} [delete]
// @param id path int true "AgentRecipe ID"
func (h AgentRecipeHandler) Delete(ctx *gin.Context) {
	id := h.pk(ctx)
	m := &model.AgentRecipe{}
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
// @summary Update an agent recipe.
// @description Update an agent recipe.
// @tags agentrecipes
// @accept json
// @success 204
// @router /agent-recipes/{id} [put]
// @param id path int true "AgentRecipe ID"
// @param recipe body AgentRecipe true "AgentRecipe data"
func (h AgentRecipeHandler) Update(ctx *gin.Context) {
	id := h.pk(ctx)
	r := &AgentRecipe{}
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

// AgentRecipe REST resource.
type AgentRecipe = resource.AgentRecipe
