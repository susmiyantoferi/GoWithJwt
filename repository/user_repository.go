package repository

import (
	"RestAPIJWT/model"
	"context"
	"database/sql"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Update(ctx context.Context, tx *sql.Tx, user model.User) model.User
	Delete(ctx context.Context, tx *sql.Tx, user model.User)
	FindAll(ctx context.Context, tx *sql.Tx) []model.User
	FindById(ctx context.Context, tx *sql.Tx, userId string) (model.User, error)
	FindByEmail(ctx context.Context, tx *sql.Tx, email string) (model.User, error)
}
