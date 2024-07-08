package models

type AuthClaims struct {
	UserID    string `json:"userId"`
	UserEmail string `json:"userEmail"`
	AccountId string `json:"accountId"`
}
