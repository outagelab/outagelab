package authservice

import (
	"context"
	"errors"
	"os"

	"google.golang.org/api/idtoken"
)

func getUserEmailFromAuthProvider(ctx context.Context, request *LoginRequest) (string, error) {
	var userID string
	switch request.AuthProvider {
	case "google":
		userID = getGoogleSignInUserID(ctx, request.AuthToken)
	default:
		return "", errors.New("invalid login provider")
	}

	return userID, nil
}

func getGoogleSignInUserID(ctx context.Context, token string) string {
	googleTokenValidator, _ := idtoken.NewValidator(ctx)

	audience := os.Getenv("GOOGLE_TOKEN_AUDIENCE")
	if audience == "" {
		return ""
	}

	p, e := googleTokenValidator.Validate(ctx, token, audience)
	if e != nil {
		return ""
	}

	return p.Claims["email"].(string)
}
