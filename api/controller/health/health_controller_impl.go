package health

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/michaelputeraw/krobot-auth-service/model"
	"github.com/newrelic/go-agent/v3/integrations/nrecho-v4"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func (c *Controller) GetServiceHealth(eCtx echo.Context) error {
	txn := nrecho.FromContext(eCtx)
	segment := newrelic.StartSegment(txn, fmt.Sprintf("%s %s", tag, tracingGetServiceHealth))
	defer segment.End()

	resp := model.NewResponse("OK", nil)
	return eCtx.JSON(http.StatusOK, resp)
}
