package repos

import (
	"context"
)

type UserSegmentRepo interface {
	GetIdSegmentByName(ctx context.Context, name string) (idSegment int, err error)
	GetActiveUserSegments(ctx context.Context, idUser int) ([]string, error)
	GetReport(ctx context.Context, year, month string) ([]string, error)

	DeleteSegments(ctx context.Context, idUser, idSegment int) error

	SegmentExistAndActive(ctx context.Context, idUser, idSegment int) bool
	SegmentExistsAndInactive(ctx context.Context, idUser, idSegment int) bool

	AddSegments(ctx context.Context, idUser, idSegment int) error
	UpdateSegments(ctx context.Context, idUser, idSegment int) error

	AddSegmentsTTL(ctx context.Context, idUser, idSegment int, ttl string) error
	UpdateSegmentsTTL(ctx context.Context, idUser, idSegment int, ttl string) error

	CheckTTL()
}
