package repos

import (
	"context"
)

type UserSegmentRepo interface {
	GetIdSegmentByName(ctx context.Context, name string) (idSegment int, err error)
	DeleteSegments(ctx context.Context, idUser, idSegment int) error
	AddSegments(ctx context.Context, idUser, idSegment int) error
	GetActiveUserSegments(ctx context.Context, idUser int) ([]string, error)
	GetReport(ctx context.Context, year, month string) ([]string, error)
}
