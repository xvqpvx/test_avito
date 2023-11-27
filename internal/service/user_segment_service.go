package service

import (
	"context"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
)

type UserSegmentService interface {
	AddRemoveSegments(ctx context.Context, request request.AddSegmentsRequest)
	GetActiveUserSegments(ctx context.Context, idUser int) response.UserSegmentResponse
	GetReport(ctx context.Context, request request.GetReport) (string, error)
	StartTTLChecker()
}
