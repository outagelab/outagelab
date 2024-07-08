package apikeyrepo

import (
	"context"
	"crypto/sha256"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/outagelab/outagelab/db"
	"github.com/outagelab/outagelab/models"
)

type apiKeyRepo struct {
	db *db.DB
}

func New(db *db.DB) *apiKeyRepo {
	return &apiKeyRepo{
		db: db,
	}
}

func (akr *apiKeyRepo) Create(ctx context.Context, apiKey *models.ApiKey) error {
	keyID := hashString(apiKey.ApiKey)
	apiKeyData := &apiKeyData{
		Version: 1,
		V1:      apiKey,
	}

	query, args := squirrel.
		Insert("api_keys").
		Columns("id", "data").
		Values(keyID, apiKeyData).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	_, err := akr.db.Exec(ctx, query, args...)
	if err != nil {
		return err
	}

	return nil
}

func (akr *apiKeyRepo) Get(ctx context.Context, apiKey string) (*models.ApiKey, error) {
	var apiKeyData apiKeyData

	keyID := hashString(apiKey)
	query, args := squirrel.
		Select("data").
		From("api_keys").
		Where("id = ?", keyID).
		PlaceholderFormat(squirrel.Dollar).
		MustSql()

	err := akr.db.QueryRow(ctx, query, args...).Scan(&apiKeyData)
	if err != nil {
		if err.Error() == "no rows in result set" {
			return nil, nil
		}
		return nil, err
	}

	return apiKeyData.V1, nil
}

func hashString(s string) string {
	h := sha256.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
	return fmt.Sprintf("%x", bs)
}
