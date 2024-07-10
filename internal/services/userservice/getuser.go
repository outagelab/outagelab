package userservice

import (
	"context"

	"github.com/outagelab/outagelab/internal/models"
)

func (s *userService) GetUser(ctx context.Context, userEmail string) (*models.User, error) {
	user, err := s.userRepo.Get(ctx, userEmail)
	if err != nil {
		return nil, err
	}
	if user == nil {
		return nil, nil
	}

	return user, nil
}
