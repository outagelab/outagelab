package models

type OutageRule struct {
	Type     string `json:"type"`
	Host     string `json:"host"`
	Status   *int   `json:"status"`
	Duration *int   `json:"duration"`
	Enabled  bool   `json:"enabled"`
}
