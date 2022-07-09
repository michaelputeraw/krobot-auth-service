package util

import (
	"context"
	"fmt"

	"github.com/newrelic/go-agent/v3/newrelic"
)

func StartTracer(ctx context.Context, tag, trackerName string) *newrelic.Segment {
	trx := newrelic.FromContext(ctx)
	segment := newrelic.StartSegment(trx, fmt.Sprintf("%s %s", tag, trackerName))
	return segment
}
