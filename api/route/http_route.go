package route

import (
	"github.com/michaelputeraw/krobot-auth-service/api/controller"
	"github.com/michaelputeraw/krobot-auth-service/infrastructure"
	"go.uber.org/fx"
)

func NewRoutes(
	infra infrastructure.Infrastructure,
	handler controller.Handler,
) {
	e := infra.Echo

	internal := e.Group("/_internal/auth")
	internal.GET("/health", handler.HealthController.GetServiceHealth)
	internal.POST("/register", handler.UserController.HandleRegisterUser)

}

var Module = fx.Options(
	fx.Invoke(NewRoutes),
)
