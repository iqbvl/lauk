package usecase

import (
	"context"

	"github.com/iqbvl/lauk/internal/model"
)

type Auth interface {
	GetUser(ctx context.Context, args model.User) (model.User, error)
	GetUserExistence(ctx context.Context, args model.User) (bool, error)
	SetUser(ctx context.Context, args model.User) error
}
