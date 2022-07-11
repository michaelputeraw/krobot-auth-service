package bootstrap

import (
	"context"
	"fmt"
	"net/http"

	"github.com/common-nighthawk/go-figure"
	"github.com/michaelputeraw/krobot-auth-service/api/controller"
	"github.com/michaelputeraw/krobot-auth-service/api/repository"
	"github.com/michaelputeraw/krobot-auth-service/api/route"
	"github.com/michaelputeraw/krobot-auth-service/api/service"
	"github.com/michaelputeraw/krobot-auth-service/infrastructure"
	"go.uber.org/fx"
)

var AppModule = fx.Options(
	infrastructure.Module,
	repository.Module,
	service.Module,
	controller.Module,
	route.Module,
	fx.Invoke(appBootstrap),
)

func appBootstrap(
	lifecycle fx.Lifecycle,
	infra infrastructure.Infrastructure,
) {

	appStop := func(context.Context) error {
		infra.Logger.Info("Stopping Application")
		conn := infra.Database.DB
		conn.Close()
		return nil
	}

	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			infra.Logger.Info("Starting Application")
			figure.NewColorFigure(infra.Config.AppName, "", "purple", true).Print()
			go func() {

				err := infra.Database.DB.Ping()
				if err != nil {
					infra.Logger.Panic(err.Error())
					panic(err)
				} else {
					infra.Logger.Info("Database connected")
				}

				_, err = infra.Redis.Ping(ctx).Result()
				if err != nil {
					infra.Logger.Panic(err.Error())
					panic(err)
				} else {
					infra.Logger.Info("Redis connected")
				}

				PORT := infra.Config.AppPort

				infra.Logger.Info(fmt.Sprintf("APP RUNNING ON http://0.0.0.0:%s", PORT))
				if err := infra.Echo.Start(fmt.Sprintf(":%s", PORT)); err != nil && err != http.ErrServerClosed {
					infra.Echo.Logger.Fatal("shutting down the server")
				}
			}()
			return nil
		},
		OnStop: appStop,
	})
}
