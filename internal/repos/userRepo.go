package repos

import (
	"context"
	"test_avito/internal/model"
)

type UserRepo interface {
	FindById(ctx context.Context, idUser int) (model.User, error)
	FindAll(ctx context.Context) []model.User
	Save(ctx context.Context, user model.User)
	Update(ctx context.Context, user model.User)
	Delete(ctx context.Context, idUser int)
}
