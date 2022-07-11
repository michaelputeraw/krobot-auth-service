package user

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michaelputeraw/krobot-auth-service/model"
	"github.com/michaelputeraw/krobot-auth-service/util"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (c *Controller) HandleRegisterUser(eCtx echo.Context) error {
	txn := nrecho.FromContext(eCtx)
	segment := newrelic.StartSegment(txn, fmt.Sprintf("%s %s", tag, tracingHandleRegisterUser))
	defer segment.End()

	ctx := eCtx.Request().Context()

	payload := new(model.UserRegisterPayload)
	if err := eCtx.Bind(payload); err != nil {
		return err
	}

	if err := eCtx.Validate(payload); err != nil {
		trans := util.TranslatorFromRequestHeader(eCtx, c.Translator)
		return echo.NewHTTPError(http.StatusBadRequest, util.BuildValidationErrors(err, trans))
	}

	err := c.UserService.RegisterUser(ctx, payload)
	if err != nil {
		return err
	}

	resp := model.NewResponse("OK", nil)

	return eCtx.JSON(http.StatusOK, resp)
}
