package user

import (
	"context"
	"fmt"

	"github.com/michaelputeraw/krobot-auth-service/model"
	"github.com/michaelputeraw/krobot-auth-service/model/database"
	"github.com/michaelputeraw/krobot-auth-service/util"
	"github.com/microcosm-cc/bluemonday"
)

func (svc *service) RegisterUser(ctx context.Context, payload *model.UserRegisterPayload) error {
	segment := util.StartTracer(ctx, tag, tracingRegister)
	defer segment.End()

	p := bluemonday.UGCPolicy()

	pass, err := util.HashPassword(payload.Password)
	if err != nil {
		svc.logger.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRegister, err))
		return err
	}
	input := &database.User{
		ID:       util.NewUUID(),
		FullName: p.Sanitize(payload.FullName),
		Gender:   payload.Gender,
		Email:    p.Sanitize(payload.Email),
		Password: pass,
	}

	err = svc.repository.UserRepository.Store(ctx, svc.db, input)
	if err != nil {
		svc.logger.Error(fmt.Sprintf("%s %s with error: %v", tag, tracingRegister, err))
		return err
	}

	return nil
}
