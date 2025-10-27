package handlers

import (
	"MiniProject-DindaYulianita/internal/usecase"
)

// Handler utama yang menyimpan semua handler spesifik
type Handler struct {
	AuthHandler *AuthHandler
}

// NewHandler membuat instance handler baru
func NewHandler(uc *usecase.Usecase) *Handler {
	return &Handler{
		AuthHandler: NewAuthHandler(uc),
	}
}
