package response

import "main/models"

type Files struct {
	Files []models.FileDto `json:"files"`
}
