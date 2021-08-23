package controllers

import (
	"0x_mt109/application/models"
	"0x_mt109/application/services"
	"0x_mt109/helpers/api"
	"encoding/json"
	"github.com/valyala/fasthttp"
	"strconv"
)

type ActorHandler struct {
	service services.IActorService
}

func NewActorHandler(service services.IActorService) *ActorHandler {
	return &ActorHandler{
		service: service,
	}
}

func (controller *ActorHandler) GetAll(ctx *fasthttp.RequestCtx) {
	actors, err := controller.service.FindAll()
	if err != nil {
		api.RenderError(ctx, err.Error(), fasthttp.StatusNotFound)
		return
	}
	api.RenderSuccess(ctx, actors, fasthttp.StatusOK)
}

func (controller *ActorHandler) Update(ctx *fasthttp.RequestCtx) {
	var request models.Actor
	err := json.Unmarshal(ctx.PostBody(), &request)
	if err != nil {
		api.RenderError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	request.Id = id
	err = controller.service.Update(request)
	if err != nil {
		api.RenderError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	api.RenderSuccess(ctx, nil, fasthttp.StatusAccepted)
}

func (controller *ActorHandler) Create(ctx *fasthttp.RequestCtx) {
	var request models.Actor
	err := json.Unmarshal(ctx.PostBody(), &request)
	if err != nil {
		api.RenderError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	err = controller.service.Create(request)
	if err != nil {
		api.RenderError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	api.RenderSuccess(ctx, nil, fasthttp.StatusCreated)
}

func (controller *ActorHandler) Delete(ctx *fasthttp.RequestCtx) {
	id, _ := strconv.Atoi(ctx.UserValue("id").(string))
	err := controller.service.Delete(id)
	if err != nil {
		api.RenderError(ctx, err.Error(), fasthttp.StatusBadRequest)
		return
	}
	api.RenderSuccess(ctx, nil, fasthttp.StatusNoContent)
}