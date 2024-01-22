package repository

import (
	"RestAPIJWT/helper"
	"RestAPIJWT/model"
	"context"
	"database/sql"
	"errors"
)

type UserRepositoryImpl struct {
}

func NewUserRepositoryImpl() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "INSERT INTO users (id, username, email, password, create_at, update_at) VALUES (?,?,?,?,?,?)"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
		user.Username,
		user.Email,
		user.Password,
		user.CreateAt,
		user.UpdateAt,
	)
	helper.PanicError(err)
	return user
}

func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user model.User) model.User {
	SQL := "UPDATE users SET username=?, update_at=? WHERE id=?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Username,
		user.UpdateAt,
		user.Id,
	)
	helper.PanicError(err)
	return user
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, user model.User) {
	SQL := "delete from users WHERE id=?"
	_, err := tx.ExecContext(
		ctx,
		SQL,
		user.Id,
	)
	helper.PanicError(err)
}

func (u *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []model.User {
	SQL := "SELECT id, username, email, create_at, update_at FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicError(err)
	defer rows.Close()

	var users []model.User
	for rows.Next() {
		user := model.User{}
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.CreateAt,
			&user.UpdateAt,
		)
		helper.PanicError(err)
		users = append(users, user)
	}

	return users
}

func (u *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId string) (model.User, error) {
	SQL := "SELECT id, username, email, create_at, update_at FROM users WHERE id=?"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		userId,
	)
	helper.PanicError(err)
	defer rows.Close()

	user := model.User{}
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.CreateAt,
			&user.UpdateAt,
		)
		helper.PanicError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u *UserRepositoryImpl) FindByEmail(ctx context.Context, tx *sql.Tx, email string) (model.User, error) {
	SQL := "SELECT id, username, email, password, create_at, update_at FROM users WHERE email=?"
	rows, err := tx.QueryContext(
		ctx,
		SQL,
		email,
	)
	helper.PanicError(err)
	defer rows.Close()

	user := model.User{}
	if rows.Next() {
		err := rows.Scan(
			&user.Id,
			&user.Username,
			&user.Email,
			&user.Password,
			&user.CreateAt,
			&user.UpdateAt,
		)
		helper.PanicError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}
