package response

import "main/models"

type Files struct {
	Files []models.File `json:"files"`
}
