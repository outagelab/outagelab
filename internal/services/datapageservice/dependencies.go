package datapageservice

import (
	"context"

	"github.com/outagelab/outagelab/internal/models"
)

type accountService interface {
	GetAccount(context.Context, string) (*models.Account, error)
	UpdateAccount(context.Context, string, *models.Account) (*models.Account, error)
}

type apiKeyRepo interface {
	Get(context.Context, string) (*models.ApiKey, error)
}
