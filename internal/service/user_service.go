package service

import (
	"context"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
)

type UserService interface {
	FindById(ctx context.Context, idUser int) response.UserResponse
	FindAll(ctx context.Context) []response.UserResponse
	Save(ctx context.Context, request request.UserCreateRequest)
	Update(ctx context.Context, request request.UserUpdateRequest)
	Delete(ctx context.Context, idUser int)
}
