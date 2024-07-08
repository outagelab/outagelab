package authservice

import (
	"context"
)

func (as *authService) Login(ctx context.Context, request *LoginRequest) (*LoginResponse, error) {
	userEmail, err := getUserEmailFromAuthProvider(ctx, request)
	if err != nil {
		return nil, err
	}
	if userEmail == "" {
		// TODO: need proper error type signaling
		return nil, nil
	}

	user, err := as.userService.GetUser(ctx, userEmail)
	if err != nil {
		return nil, err
	}

	if user == nil {
		account, err := as.accountService.CreateAccount(ctx)
		if err != nil {
			return nil, err
		}

		user, err = as.userService.CreateUser(ctx, userEmail, account.ID)
		if err != nil {
			return nil, err
		}
	}

	token, err := generateToken(user)
	if err != nil {
		return nil, err
	}

	response := &LoginResponse{
		AuthToken: token,
	}

	return response, nil
}
