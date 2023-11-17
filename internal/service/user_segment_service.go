package service

import (
	"context"
	"test_avito/internal/data/response"
)

type UserSegmentService interface {
	AddRemoveSegments(ctx context.Context, idUser int, segmentsToAdd []string, segmentsToDelete []string)
	GetActiveUserSegments(ctx context.Context, idUser int) response.UserSegmentResponse
}
