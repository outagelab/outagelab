package accountrepo

import "github.com/outagelab/outagelab/models"

type accountData struct {
	Version int             `json:"version"`
	V1      *models.Account `json:"v1"`
}
