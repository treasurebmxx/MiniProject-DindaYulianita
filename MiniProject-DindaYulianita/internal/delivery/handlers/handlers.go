package handlers

import (
	"MiniProject-DindaYulianita/internal/usecase"
)


type Handler struct {
	AuthHandler *AuthHandler
}


func NewHandler(uc *usecase.Usecase) *Handler {
	return &Handler{
		AuthHandler: NewAuthHandler(uc),
	}
}
