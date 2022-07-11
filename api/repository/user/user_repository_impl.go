package user

import (
	"context"
	"fmt"

	"github.com/michaelputeraw/krobot-auth-service/infrastructure"
	"github.com/michaelputeraw/krobot-auth-service/model/database"
	"github.com/michaelputeraw/krobot-auth-service/util"
)

func (r *repository) Store(ctx context.Context, db infrastructure.Querier, input *database.User) error {
	segment := util.StartTracer(ctx, tag, tracingStore)

	defer segment.End()

	query, args, err := r.buildInsertQuery(input).ToSql()

	if err != nil {
		r.logger.Error(fmt.Sprintf("%s %s with: %v", tag, tracingStore, err))
		return err
	}

	_, err = db.ExecContext(ctx, query, args...)
	if err != nil {
		r.logger.Error(fmt.Sprintf("%s %s with: %v", tag, tracingStore, err))
		return err
	}

	return nil
}
