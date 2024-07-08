package models

type Account struct {
	ID           string          `json:"id"`
	ApiKeys      []*AcountApiKey `json:"apiKeys"`
	Applications []*Application  `json:"applications"`
}

type AcountApiKey struct {
	Key string `json:"key"`
}
