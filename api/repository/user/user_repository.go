package user

import (
	"context"

	"github.com/michaelputeraw/krobot-auth-service/infrastructure"
	"github.com/michaelputeraw/krobot-auth-service/model/database"
	"go.uber.org/zap"
)

const (
	tag = `[UserRepository]`

	tracingStore = "StoreUser"
)

type (
	UserRepository interface {
		GetTableName() string
		Store(ctx context.Context, db infrastructure.Querier, input *database.User) error
	}
	repository struct {
		logger *zap.Logger
	}
)

func New(infrastructure infrastructure.Infrastructure) UserRepository {
	return &repository{
		logger: infrastructure.Logger,
	}
}

func (r *repository) GetTableName() string {
	return "users"
}
