package accountservice

import (
	"context"

	"github.com/outagelab/outagelab/models"
)

func (s *accountService) GetAccount(ctx context.Context, accountID string) (*models.Account, error) {
	account, err := s.accountRepo.Get(ctx, accountID)
	if err != nil {
		return nil, err
	}

	return account, nil
}
