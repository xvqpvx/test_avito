package repos

import (
	"context"
	"test_avito/internal/model"
)

type SegmentRepo interface {
	FindById(ctx context.Context, idSegment int) (model.Segment, error)
	FindByName(ctx context.Context, name string) (model.Segment, error)
	FindAll(ctx context.Context) []model.Segment
	Save(ctx context.Context, segment model.Segment)
	Update(ctx context.Context, segment model.Segment)
	Delete(ctx context.Context, segment model.Segment)
}
