package userrepo

import "github.com/outagelab/outagelab/internal/models"

type userData struct {
	Version int          `json:"version"`
	V1      *models.User `json:"v1"`
}
