package user

import (
	"context"

	"github.com/jmoiron/sqlx"
	"github.com/michaelputeraw/krobot-auth-service/api/repository"
	"github.com/michaelputeraw/krobot-auth-service/infrastructure"
	"github.com/michaelputeraw/krobot-auth-service/model"
	"go.uber.org/zap"
)

const (
	tag = `[UserService]`

	tracingRegister = "RegisterUser"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, payload *model.UserRegisterPayload) error
	}
	service struct {
		logger     *zap.Logger
		db         *sqlx.DB
		repository repository.Repository
	}
)

func New(
	infrastructure infrastructure.Infrastructure,
	repository repository.Repository,
) UserService {
	return &service{
		logger:     infrastructure.Logger,
		db:         infrastructure.Database.SqlxDB,
		repository: repository,
	}
}
