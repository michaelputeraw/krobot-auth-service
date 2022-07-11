package infrastructure

import (
	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	"github.com/go-playground/universal-translator"
)

func NewTranslator() *ut.UniversalTranslator {
	langEN := en.New()
	langID := id.New()
	uni := ut.New(langEN, langEN, langID)
	return uni
}
