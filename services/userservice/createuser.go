package userservice

import (
	"context"

	"github.com/google/uuid"
	"github.com/outagelab/outagelab/models"
)

func (s *userService) CreateUser(ctx context.Context, email string, accountID string) (*models.User, error) {
	user, err := s.userRepo.Create(ctx, &models.User{
		ID:    uuid.NewString(),
		Email: email,
		Accounts: []*models.UserAccount{
			{
				AccountID: accountID,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return user, nil
}
