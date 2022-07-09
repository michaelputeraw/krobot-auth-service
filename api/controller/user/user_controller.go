package user

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/michaelputeraw/krobot-auth-service/api/service/user"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	tag = "[UserController]"

	tracingHandleRegisterUser = "handleRegisterUser"
)

type Controller struct {
	fx.In

	Logger      *zap.Logger
	Translator  *ut.UniversalTranslator
	UserService user.UserService
}
