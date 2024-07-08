package authservice

import (
	"context"

	"github.com/outagelab/outagelab/models"
)

type userService interface {
	GetUser(context.Context, string) (*models.User, error)
	CreateUser(ctx context.Context, userEmail, accountID string) (*models.User, error)
}

type accountService interface {
	GetAccount(ctx context.Context, userEmail string) (*models.Account, error)
	CreateAccount(ctx context.Context) (*models.Account, error)
}
