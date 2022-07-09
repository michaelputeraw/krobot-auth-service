package health

import (
	ut "github.com/go-playground/universal-translator"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

const (
	tag = "[HealthController]"

	tracingGetServiceHealth = "getServiceHealth"
)

type Controller struct {
	fx.In

	Logger     *zap.Logger
	Translator *ut.UniversalTranslator
}
