package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"test_avito/internal/helper"
	"time"
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

// Проверяем, был ли пользователь ранее добавлен в сегмент и впоследствии удален (чтобы не создавать дубликат)
func (us *UserSegmentRepoImpl) SegmentExistsAndInactive(ctx context.Context, idUser, idSegment int) bool {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	existingQuery := "SELECT 1 FROM users_segments WHERE id_user=? AND id_segment=? AND is_active=FALSE"
	var existing int
	err = tx.QueryRowContext(ctx, existingQuery, idUser, idSegment).Scan(&existing)

	if err == nil {
		return true
	} else {
		return false
	}
}

// Проверяем, добавлен ли сегмент пользователю уже (чтобы не создавать дубликат)
func (us *UserSegmentRepoImpl) SegmentExistAndActive(ctx context.Context, idUser, idSegment int) bool {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	var existing int
	query := "SELECT 1 FROM users_segments WHERE id_user=? AND id_segment=? AND is_active=TRUE"
	err = tx.QueryRowContext(ctx, query, idUser, idSegment).Scan(&existing)

	if err == nil {
		return true
	} else {
		return false
	}
}

// Добавление сегмента пользователю с TTL, когда пользователь ранее был в сегменте
func (us *UserSegmentRepoImpl) UpdateSegmentsTTL(ctx context.Context, idUser, idSegment int, ttl string) error {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE users_segments SET is_active=TRUE, enter_date=NOW(), operation='insertion', ttl=? WHERE id_user=? and id_segment=?"
	_, err = tx.ExecContext(ctx, query, ttl /*.Format("2006-01-02 15:04:05")*/, idUser, idSegment)
	helper.PanicIfError(err)
	return nil
}

// Добавление сегмента пользователю без TTL, когда пользователь ранее был в сегменте
func (us *UserSegmentRepoImpl) UpdateSegments(ctx context.Context, idUser, idSegment int) error {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE users_segments SET is_active=TRUE, enter_date=NOW(), operation='insertion' WHERE id_user=? and id_segment=?"
	_, err = tx.ExecContext(ctx, query, idUser, idSegment)
	helper.PanicIfError(err)
	return nil
}

// Сегмента у пользователя еще нет и мы можем его добавить (с TTL)
func (us *UserSegmentRepoImpl) AddSegmentsTTL(ctx context.Context, idUser, idSegment int, ttl string) error {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	insertQuery := "INSERT INTO users_segments (id_user, id_segment, enter_date, is_active, operation, ttl) VALUES (?, ?, NOW(), TRUE, 'insertion', ?)"
	_, err = tx.ExecContext(ctx, insertQuery, idUser, idSegment, ttl /*.Format("2006-01-02 15:04:05")*/)
	helper.PanicIfError(err)

	return nil
}

// Сегмента у пользователя еще нет и мы можем его добавить (без TTL)
func (us *UserSegmentRepoImpl) AddSegments(ctx context.Context, idUser, idSegment int) error {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	insertQuery := "INSERT INTO users_segments (id_user, id_segment, enter_date, is_active, operation) VALUES (?, ?, NOW(), TRUE, 'insertion')"
	_, err = tx.ExecContext(ctx, insertQuery, idUser, idSegment)
	helper.PanicIfError(err)

	return nil
}

func (us *UserSegmentRepoImpl) CheckTTL() {
	currTime := time.Now().Format("2006-01-02 15:04:05")
	log.Print("Current Time:", currTime)
	//ttlTime, err := time.Parse("2006-01-02 15:04:05", ttl)
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE users_segments SET is_active=FALSE, operation='deletion' WHERE ttl<?"
	_, err = tx.Exec(query, currTime)
	if err != nil {
		log.Error().Msgf("Err in CheckTTL: ", err)
	}
	//helper.PanicIfError(err)
}

func (us *UserSegmentRepoImpl) DeleteSegments(ctx context.Context, idUser, idSegment int) error {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE users_segments SET is_active=FALSE, operation='deletion', enter_date=NOW() WHERE id_user=? AND id_segment=?"
	result, err := tx.ExecContext(ctx, query, idUser, idSegment)

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

func (us *UserSegmentRepoImpl) GetReport(ctx context.Context, year, month string) ([]string, error) {
	tx, err := us.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := `
		SELECT
    		us.id_user,
    		s.name AS segment_name,
    		us.operation,
    		us.enter_date
		FROM
    		users_segments us
		JOIN
    		segments s ON us.id_segment = s.id_segment
		WHERE
    		MONTH(us.enter_date) = ? AND
    		YEAR(us.enter_date) = ?;`
	result, err := tx.QueryContext(ctx, query, month, year)
	helper.PanicIfError(err)
	defer result.Close()

	var report []string
	for result.Next() {
		var idUser int
		var segmentName string
		var operation string
		var enterDate string

		err = result.Scan(&idUser, &segmentName, &operation, &enterDate)
		helper.PanicIfError(err)
		report = append(report, fmt.Sprintf("%d; %s; %s; %s", idUser, segmentName, operation, enterDate))
	}

	return report, nil
}
