package repository

import (
	"github.com/michaelputeraw/krobot-auth-service/api/repository/user"
	"go.uber.org/fx"
)

type Repository struct {
	fx.In

	UserRepository user.UserRepository
}

var Module = fx.Options(
	fx.Provide(user.New),
)
