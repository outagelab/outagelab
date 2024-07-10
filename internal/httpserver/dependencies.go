package httpserver

import (
	"context"

	"github.com/outagelab/outagelab/internal/models"
	"github.com/outagelab/outagelab/internal/services/authservice"
	"github.com/outagelab/outagelab/internal/services/datapageservice"
)

type accountService interface {
	GetAccount(ctx context.Context, userID string) (*models.Account, error)
	UpdateAccount(ctx context.Context, userID string, account *models.Account) (*models.Account, error)
}

type datapageService interface {
	GenerateDatapageDeprecated(
		ctx context.Context,
		apiKeyStr string,
		request *datapageservice.GenerateDatapageRequest,
	) (*models.Account, error)

	GenerateDatapage(
		ctx context.Context,
		apiKeyStr string,
		request *datapageservice.GenerateDatapageRequest,
	) (*datapageservice.GenerateDatapageResponse, error)
}

type authService interface {
	ValidateToken(string) (*models.AuthClaims, error)
	Login(ctx context.Context, request *authservice.LoginRequest) (*authservice.LoginResponse, error)
}
