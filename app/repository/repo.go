package app

import "context"

type (
	UserRepo interface {
		FindByID(ctx context.Context, id uuid.UUID) (model.User, error)
	}
)
