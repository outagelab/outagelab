package apikeyrepo

import "github.com/outagelab/outagelab/models"

type apiKeyData struct {
	Version int            `json:"version"`
	V1      *models.ApiKey `json:"v1"`
}
