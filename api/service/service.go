package service

import (
	"github.com/michaelputeraw/krobot-auth-service/api/service/user"
	"go.uber.org/fx"
)

type Service struct {
	fx.In

	UserService user.UserService
}

var Module = fx.Options(
	fx.Provide(user.New),
)
