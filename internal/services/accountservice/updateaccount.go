package accountservice

import (
	"context"

	"github.com/outagelab/outagelab/internal/models"
)

func (as *accountService) UpdateAccount(ctx context.Context, accountID string, account *models.Account) (*models.Account, error) {
	account.ID = accountID
	account, err := as.accountRepo.Update(ctx, account)
	if err != nil {
		return nil, err
	}

	return account, nil
}
