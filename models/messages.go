package models

type MessageResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}
