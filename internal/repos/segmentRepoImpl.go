package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"test_avito/internal/helper"
	"test_avito/internal/model"
)

type SegmentRepoImpl struct {
	Db *sql.DB
}

func NewSegmentRepository(Db *sql.DB) SegmentRepo {
	return &SegmentRepoImpl{Db: Db}
}

func (s *SegmentRepoImpl) FindById(ctx context.Context, idSegment int) (model.Segment, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT * FROM segments WHERE id_segment=?"
	result, errQuery := tx.QueryContext(ctx, query, idSegment)
	helper.PanicIfError(errQuery)
	defer result.Close()

	segment := model.Segment{}

	if result.Next() {
		err := result.Scan(&segment.IdSegment, &segment.Name, &segment.IsActive)
		helper.PanicIfError(err)
		return segment, nil
	} else {
		return segment, errors.New(fmt.Sprintf("User with id %d not found", idSegment))
	}
}

func (s *SegmentRepoImpl) FindByName(ctx context.Context, name string) (model.Segment, error) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT * FROM segments WHERE name=?"
	result, errQuery := tx.QueryContext(ctx, query, name)
	helper.PanicIfError(errQuery)

	segment := model.Segment{}

	if result.Next() {
		err := result.Scan(&segment.IdSegment, &segment.Name, &segment.IsActive)
		helper.PanicIfError(err)
		return segment, nil
	} else {
		return segment, errors.New(fmt.Sprintf("Segment with name %s not found", name))
	}
}

func (s *SegmentRepoImpl) FindAll(ctx context.Context) []model.Segment {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT * FROM segments"
	result, errQuery := tx.QueryContext(ctx, query)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var segments []model.Segment

	for result.Next() {
		segment := model.Segment{}
		err := result.Scan(&segment.IdSegment, &segment.Name, &segment.IsActive)
		helper.PanicIfError(err)

		segments = append(segments, segment)
	}

	return segments
}

func (s *SegmentRepoImpl) Save(ctx context.Context, segment model.Segment) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "INSERT INTO segments(name) VALUES (?)"
	_, err = tx.ExecContext(ctx, query, segment.Name)
	helper.PanicIfError(err)

}

func (s *SegmentRepoImpl) Update(ctx context.Context, segment model.Segment) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE segments SET name=? WHERE id_segment=?"
	_, err = tx.ExecContext(ctx, query, segment.Name, segment.IdSegment)
	helper.PanicIfError(err)
}

func (s *SegmentRepoImpl) Delete(ctx context.Context, segment model.Segment) {
	tx, err := s.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE segments SET is_active=FALSE WHERE name=?"
	_, err = tx.ExecContext(ctx, query, segment.Name)
	helper.PanicIfError(err)
}
