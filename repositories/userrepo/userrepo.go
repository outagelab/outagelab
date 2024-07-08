package userrepo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/outagelab/outagelab/db"
	"github.com/outagelab/outagelab/models"
)

type userRepo struct {
	db *db.DB
}

func New(db *db.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

func (r *userRepo) Create(ctx context.Context, user *models.User) (*models.User, error) {
	userData := &userData{
		Version: 1,
		V1:      user,
	}

	query, args := squirrel.
		Insert("users").
		Columns("id", "email", "data").
		Values(user.ID, user.Email, userData).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	_, err := r.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (r *userRepo) Get(ctx context.Context, userEmail string) (*models.User, error) {
	var userData userData

	query, args := squirrel.
		Select("data").
		From("users").
		Where("email = ?", userEmail).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	err := r.db.QueryRow(ctx, query, args...).Scan(&userData)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return userData.V1, nil
}
