package userservice

import (
	"context"

	"github.com/outagelab/outagelab/models"
)

type userRepo interface {
	Create(context.Context, *models.User) (*models.User, error)
	Get(context.Context, string) (*models.User, error)
}
