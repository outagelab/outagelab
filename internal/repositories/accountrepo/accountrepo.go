package accountrepo

import (
	"context"

	"github.com/Masterminds/squirrel"
	"github.com/outagelab/outagelab/internal/db"
	"github.com/outagelab/outagelab/internal/models"
)

type accountRepo struct {
	db *db.DB
}

func New(db *db.DB) *accountRepo {
	return &accountRepo{
		db: db,
	}
}

func (ar *accountRepo) Get(ctx context.Context, accountID string) (*models.Account, error) {
	var accountData accountData

	query, args := squirrel.
		Select("data").
		From("accounts").
		Where("id = ?", accountID).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	err := ar.db.QueryRow(ctx, query, args...).Scan(&accountData)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return accountData.V1, nil
}

func (ar *accountRepo) Create(ctx context.Context, account *models.Account) (*models.Account, error) {
	accountData := &accountData{
		Version: 1,
		V1:      account,
	}

	query, args := squirrel.
		Insert("accounts").
		Columns("id", "data").
		Values(account.ID, accountData).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	_, err := ar.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return account, nil
}

func (ar *accountRepo) Update(ctx context.Context, account *models.Account) (*models.Account, error) {
	accountData := &accountData{
		Version: 1,
		V1:      account,
	}

	query, args := squirrel.
		Update("accounts").
		Set("data", accountData).
		Where("id = ?", account.ID).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	_, err := ar.db.Exec(ctx, query, args...)
	if err != nil {
		return nil, err
	}

	return account, nil
}
