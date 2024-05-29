package service

import (
	"context"

	"api.kalbe.crm/apps/domain"
)

type IAuth interface {
	SignIn(ctx context.Context, payload domain.AuthRequest) (token string, err error)
}
