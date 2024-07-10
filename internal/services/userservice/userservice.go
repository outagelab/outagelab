package userservice

type userService struct {
	userRepo userRepo
}

func New(userRepo userRepo) *userService {
	return &userService{
		userRepo: userRepo,
	}
}
