package infrastructure

import (
	"github.com/go-playground/universal-translator"
	"github.com/go-redis/redis/v9"
	"github.com/labstack/echo/v4"
	"github.com/newrelic/go-agent/v3/newrelic"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type Infrastructure struct {
	fx.In

	Config      *Config
	Logger      *zap.Logger
	Database    *Database
	NewRelicAPM *newrelic.Application
	Redis       *redis.Client
	Translator  *ut.UniversalTranslator
	Cv          echo.Validator
	Echo        *echo.Echo
}

func NewInfrastructure() *Infrastructure {
	return &Infrastructure{}
}

var Module = fx.Options(
	fx.Provide(NewConfig),
	fx.Provide(NewLogger),
	fx.Provide(NewRelicAPM),
	fx.Provide(NewDatabase),
	fx.Provide(NewRedis),
	fx.Provide(NewTranslator),
	fx.Provide(NewValidator),
	fx.Provide(NewRouter),

	fx.Populate(NewInfrastructure()),
)
