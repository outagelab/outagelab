package accountservice

type accountService struct {
	accountRepo accountRepo
	apiKeyRepo  apiKeyRepo
}

func New(accountRepo accountRepo, apiKeyRepo apiKeyRepo) *accountService {
	return &accountService{
		accountRepo: accountRepo,
		apiKeyRepo:  apiKeyRepo,
	}
}
