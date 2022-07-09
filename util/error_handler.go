package util

import (
	"fmt"
	"net/http"
	"strings"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"github.com/michaelputeraw/krobot-auth-service/constant"
	"github.com/michaelputeraw/krobot-auth-service/model"
)

func TranslatorFromRequestHeader(ctx echo.Context, trans *ut.UniversalTranslator) ut.Translator {
	lang := ctx.Request().Header.Get("Accept-Language")
	translator, found := trans.GetTranslator(lang)

	if !found {
		translator, _ = trans.GetTranslator(constant.LANG_DEFAULT)
	}
	return translator
}

func BuildValidationErrors(err error, trans ut.Translator) []model.ValidationError {
	errors := []model.ValidationError{}
	for _, err := range err.(validator.ValidationErrors) {
		field := strings.Join(strings.Split(err.Namespace(), ".")[1:], ".")
		buildMessage := err.Translate(trans)
		if lang, ok := constant.Fields[field]; ok {
			if fieldName, ok := lang[trans.Locale()]; ok {
				buildMessage = strings.Replace(buildMessage, field, fieldName, 1)
			}
		}
		errors = append(errors, model.ValidationError{
			Field:   field,
			Message: buildMessage,
		})
	}
	return errors
}

func CustomHTTPErrorHandler(err error, c echo.Context) {
	var (
		response model.RequestError
	)

	// default
	statusCode := http.StatusInternalServerError
	message := "Internal Server Error"

	if e, ok := err.(*model.HttpCustomError); ok {
		statusCode = e.StatusCode
		response.Error = model.HttpCustomError{
			StatusCode: e.StatusCode,
			Message:    e.Message,
		}
	} else if e, ok := err.(*echo.HTTPError); ok {
		statusCode = e.Code
		if m, ok := e.Message.([]model.ValidationError); ok {
			response.Error = model.HttpCustomErrors{
				Message: "validation errors",
				Errors:  m,
			}
		} else {
			response.Error = model.HttpCustomError{
				StatusCode: statusCode,
				Message:    fmt.Sprintf("%s", e.Message),
			}
		}
	} else {
		response.Error = model.HttpCustomError{
			StatusCode: http.StatusInternalServerError,
			Message:    message,
		}
	}

	c.JSON(statusCode, response)
}
