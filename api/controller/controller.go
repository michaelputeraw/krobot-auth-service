package controller

import (
	"github.com/michaelputeraw/krobot-auth-service/api/controller/health"
	"github.com/michaelputeraw/krobot-auth-service/api/controller/user"
	"go.uber.org/fx"
)

type Handler struct {
	fx.In

	HealthController health.Controller
	UserController   user.Controller
}

func NewHandler() *Handler {
	return &Handler{}
}

var Module = fx.Options(
	fx.Populate(NewHandler()),
)
