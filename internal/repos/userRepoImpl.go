package repos

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"test_avito/internal/helper"
	"test_avito/internal/model"
	"time"
)

type UserRepoImpl struct {
	Db *sql.DB
}

func NewUserRepository(Db *sql.DB) UserRepo {
	return &UserRepoImpl{Db: Db}
}

func (u *UserRepoImpl) FindById(ctx context.Context, idUser int) (model.User, error) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT * FROM users WHERE id_user=?"
	result, errQuery := tx.QueryContext(ctx, query, idUser)
	helper.PanicIfError(errQuery)
	defer result.Close()

	user := model.User{}

	if result.Next() {
		err := result.Scan(&user.IdUser, &user.Firstname, &user.Lastname,
			&user.Created, &user.Updated, &user.IsActive)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New(fmt.Sprintf("User with id %d not found", idUser))
	}
}

func (u *UserRepoImpl) FindAll(ctx context.Context) []model.User {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "SELECT * FROM users"
	result, errQuery := tx.QueryContext(ctx, query)
	helper.PanicIfError(errQuery)
	defer result.Close()

	var users []model.User

	for result.Next() {
		user := model.User{}
		err := result.Scan(&user.IdUser, &user.Firstname, &user.Lastname, &user.Created, &user.Updated)
		helper.PanicIfError(err)

		users = append(users, user)
	}

	return users
}

func (u *UserRepoImpl) Save(ctx context.Context, user model.User) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "INSERT INTO users(firstname, lastname, created, updated) VALUES (?,?,?,?)"
	_, err = tx.ExecContext(ctx, query, user.Firstname, user.Lastname, user.Created, user.Updated)
	helper.PanicIfError(err)
}

func (u *UserRepoImpl) Update(ctx context.Context, user model.User) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	query := "UPDATE users SET firstname=?, lastname=?, updated=? WHERE id_user=?"
	_, err = tx.ExecContext(ctx, query, user.Firstname, user.Lastname, user.Updated, user.IdUser)
	helper.PanicIfError(err)
}

func (u *UserRepoImpl) Delete(ctx context.Context, idUser int) {
	tx, err := u.Db.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	now := time.Now().Format("2006-01-02 15:04:05")
	query := "UPDATE users SET is_active=FALSE, updated=? WHERE id_user=?"
	_, err = tx.ExecContext(ctx, query, now, idUser)
	helper.PanicIfError(err)
}
