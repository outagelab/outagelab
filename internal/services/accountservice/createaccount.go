package accountservice

import (
	"context"

	"github.com/google/uuid"
	"github.com/outagelab/outagelab/internal/models"
)

func (as *accountService) CreateAccount(ctx context.Context) (*models.Account, error) {
	apiKeyStr := uuid.NewString()
	accountId := uuid.NewString()

	account, err := as.accountRepo.Create(ctx, &models.Account{
		ID: accountId,
		ApiKeys: []*models.AcountApiKey{
			{
				Key: apiKeyStr,
			},
		},
		Applications: []*models.Application{},
	})
	if err != nil {
		return nil, err
	}

	err = as.apiKeyRepo.Create(ctx, &models.ApiKey{
		AccountID: accountId,
		ApiKey:    apiKeyStr,
	})
	if err != nil {
		return nil, err
	}

	return account, nil
}
