package infrastructure

import (
	"reflect"
	"strings"

	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTranslations "github.com/go-playground/validator/v10/translations/en"
	idTranslations "github.com/go-playground/validator/v10/translations/id"
	"github.com/labstack/echo/v4"
	"github.com/michaelputeraw/krobot-auth-service/constant"
	cv "github.com/michaelputeraw/krobot-auth-service/infrastructure/custom-validator"
)

type (
	ValidationUtil struct {
		validator *validator.Validate
	}
)

func registerTagNameWithLabel(validate *validator.Validate) {
	validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("label"), ",", 2)[0]
		if name == "-" {
			return ""
		}
		return name
	})
}

func NewValidator(trans *ut.UniversalTranslator, db *Database) echo.Validator {
	validate := validator.New()
	registerTagNameWithLabel(validate)

	langEN, _ := trans.GetTranslator(constant.LANG_DEFAULT)
	langID, _ := trans.GetTranslator(constant.LANG_ID)
	_ = enTranslations.RegisterDefaultTranslations(validate, langEN)
	_ = idTranslations.RegisterDefaultTranslations(validate, langID)
	customValidation := cv.New(db.SqlxDB)
	RegisterCustomValidation(customValidation, validate, trans)

	return &ValidationUtil{validator: validate}
}

func (v *ValidationUtil) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

func RegisterCustomValidation(customValidation cv.CustomValidator, val *validator.Validate, trans *ut.UniversalTranslator) {
	langEN, _ := trans.GetTranslator(constant.LANG_DEFAULT)
	langID, _ := trans.GetTranslator(constant.LANG_ID)

	val.RegisterValidation("uniquedb", customValidation.UniqueValidator)
	registerTranslation(val, langEN, "uniquedb", "{0} already exist")
	registerTranslation(val, langID, "uniquedb", "{0} sudah digunakan")

	val.RegisterValidation("existdb", customValidation.ExistValidator)
	registerTranslation(val, langEN, "existdb", "{0} not exist")
	registerTranslation(val, langID, "existdb", "{0} tidak ditemukan")

}

func registerTranslation(v *validator.Validate, trans ut.Translator, tag string, message string) {
	_ = v.RegisterTranslation(tag, trans,
		func(ut ut.Translator) error {
			return ut.Add(tag, message, true)
		},
		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T(tag, fe.Field())
			return t
		},
	)
}
