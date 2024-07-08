package models

type User struct {
	ID       string         `json:"id"`
	Email    string         `json:"email"`
	Accounts []*UserAccount `json:"accounts"`
}

type UserAccount struct {
	AccountID string `json:"accountId"`
}
