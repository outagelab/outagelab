package authservice

type authService struct {
	userService    userService
	accountService accountService
}

func New(userService userService, accountService accountService) *authService {
	return &authService{
		userService:    userService,
		accountService: accountService,
	}
}
