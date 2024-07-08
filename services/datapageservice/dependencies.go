package datapageservice

import (
	"context"

	"github.com/outagelab/outagelab/models"
)

type accountRepo interface {
	Get(context.Context, string) (*models.Account, error)
}

type apiKeyRepo interface {
	Get(context.Context, string) (*models.ApiKey, error)
}
