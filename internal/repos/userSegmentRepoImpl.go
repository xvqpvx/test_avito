package repos

import (
	"context"
	"database/sql"
	"errors"
	"github.com/rs/zerolog/log"
	"test_avito/internal/helper"
)

type UserSegmentRepoImpl struct {
	Db *sql.DB
}

func NewUserSegmentRepository(Db *sql.DB) UserSegmentRepo {
	return &UserSegmentRepoImpl{Db: Db}
}

func (us *UserSegmentRepoImpl) GetIdSegmentByName(ctx context.Context, name string) (idSegment int, err error) {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT id_segment FROM segments WHERE name=?"
	result, errQuery := tx.QueryContext(ctx, query, name)
	helper.PanicIfError(errQuery)
	defer result.Close()

	if result.Next() {
		err := result.Scan(&idSegment)
		helper.PanicIfError(err)
		return idSegment, nil
	} else {
		return 0, nil
	}
}

func (us *UserSegmentRepoImpl) AddSegments(ctx context.Context, idUser, idSegment int) error {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	existingQuery := "SELECT 1 FROM users_segments WHERE id_user=? AND id_segment=? AND is_active=FALSE"
	var existing int
	err = tx.QueryRowContext(ctx, existingQuery, idUser, idSegment).Scan(&existing)
	if err == nil {
		query := "UPDATE users_segments SET is_active=TRUE WHERE id_user=? and id_segment=?"
		_, err := tx.ExecContext(ctx, query, idUser, idSegment)
		helper.PanicIfError(err)
		return nil
	} else if err != sql.ErrNoRows {
		//Произошла ошибка при выполнении запроса
		return err
	}

	query := "SELECT 1 FROM users_segments WHERE id_user=? AND id_segment=? AND is_active=TRUE"
	err = tx.QueryRowContext(ctx, query, idUser, idSegment).Scan(&existing)
	if err == nil {
		//Сегмент уже добавлен пользователю (не добавляем дубликат)
		return nil
	} else {
		//Связи еще нет, и мы можем её добавить
		insertQuery := "INSERT INTO users_segments (id_user, id_segment, enter_date, is_active) VALUES (?, ?, NOW(), TRUE)"
		_, err = tx.ExecContext(ctx, insertQuery, idUser, idSegment)
		helper.PanicIfError(err)
	}

	return nil
}

func (us *UserSegmentRepoImpl) DeleteSegments(ctx context.Context, idUser, idSegment int) error {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE users_segments SET is_active=FALSE WHERE id_user=? AND id_segment=?"
	result, err := tx.ExecContext(ctx, query, idUser, idSegment)
	log.Info().Msgf("idSegment - %d, idUser - %d", idSegment, idUser)
	if err != nil {
		return err
	}

	rowsAffected, err := result.RowsAffected()
	helper.PanicIfError(err)
	if rowsAffected == 0 {
		return errors.New("no rows were updated")
	}
	return nil
}

func (us *UserSegmentRepoImpl) GetActiveUserSegments(ctx context.Context, idUser int) ([]string, error) {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT s.name FROM users_segments us JOIN segments s ON us.id_segment = s.id_segment WHERE us.id_user = ? AND us.is_active = TRUE"
	result, err := tx.QueryContext(ctx, query, idUser)
	helper.PanicIfError(err)

	var activeSegments []string

	for result.Next() {
		var tmp string
		err := result.Scan(&tmp)
		helper.PanicIfError(err)

		activeSegments = append(activeSegments, tmp)
	}

	return activeSegments, nil
}
