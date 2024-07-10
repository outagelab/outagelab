package accountservice

import (
	"context"

	"github.com/outagelab/outagelab/internal/models"
)

type accountService struct {
	accountRepo accountRepo
	apiKeyRepo  apiKeyRepo
}

func New(accountRepo accountRepo, apiKeyRepo apiKeyRepo) *accountService {
	return &accountService{
		accountRepo: accountRepo,
		apiKeyRepo:  apiKeyRepo,
	}
}

func UpdateAccount(ctx context.Context, userID string, account *models.Account) (*models.Account, error) {
	return nil, nil
}

func GetAccountByApiKey(ctx context.Context, apiKey string) (*models.Account, error) {
	return nil, nil
}
