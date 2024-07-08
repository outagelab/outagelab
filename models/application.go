package models

type Application struct {
	ID           string         `json:"id"`
	Environments []*Environment `json:"environments"`
	Rules        []*OutageRule  `json:"rules"`
}
