package service

import (
	"context"
	"test_avito/internal/data/request"
	"test_avito/internal/data/response"
)

type SegmentService interface {
	FindById(ctx context.Context, idSegment int) response.SegmentResponse
	FindByName(cxt context.Context, name string) response.SegmentResponse
	FindAll(ctx context.Context) []response.SegmentResponse
	Save(ctx context.Context, request request.SegmentCreateRequest)
	Update(ctx context.Context, request request.SegmentUpdateRequest)
	Delete(ctx context.Context, request request.SegmentDeleteRequest)
}
