package authservice

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/outagelab/outagelab/internal/models"
)

func generateToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss":       os.Getenv("AUTH_ISSUER"),
		"sub":       user.Email,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
		"accountid": user.Accounts[0].AccountID,
		"userid":    user.ID,
	})
	tokenStr, err := token.SignedString([]byte(os.Getenv("AUTH_SYMMETRIC_KEY")))
	if err != nil {
		// TODO
		return "", err
	}

	return tokenStr, nil
}

func (as *authService) ValidateToken(tokenStr string) (*models.AuthClaims, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("AUTH_SYMMETRIC_KEY")), nil
	})
	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, nil
	}

	tokenClaims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, nil
	}

	userEmail, ok := tokenClaims["sub"].(string)
	if !ok {
		// TODO: how do i consolidate this?
		return nil, errors.New("invalid token: missing claim sub")
	}

	userID, ok := tokenClaims["userid"].(string)
	if !ok {
		// TODO: how do i consolidate this?
		return nil, errors.New("invalid token: missing claim userid")
	}

	accountID, ok := tokenClaims["accountid"].(string)
	if !ok {
		// TODO: how do i consolidate this?
		return nil, errors.New("invalid token: missing claim accountid")
	}

	claims := &models.AuthClaims{
		UserEmail: userEmail,
		UserID:    userID,
		AccountId: accountID,
	}

	return claims, nil
}
