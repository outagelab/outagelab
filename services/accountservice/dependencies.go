package accountservice

import (
	"context"

	"github.com/outagelab/outagelab/models"
)

type accountRepo interface {
	Get(context.Context, string) (*models.Account, error)
	Create(context.Context, *models.Account) (*models.Account, error)
	Update(context.Context, *models.Account) (*models.Account, error)
}

type apiKeyRepo interface {
	Create(context.Context, *models.ApiKey) error
}
