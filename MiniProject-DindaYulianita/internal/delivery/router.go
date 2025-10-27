package delivery

import (
	"MiniProject-DindaYulianita/internal/usecase"
	"MiniProject-DindaYulianita/internal/config"
	"MiniProject-DindaYulianita/internal/delivery/handlers"
	"github.com/gorilla/mux"
)

func NewRouter(r *mux.Router, uc *usecase.Usecase, cfg *config.Config) {
	authHandler := handlers.NewAuthHandler(uc)

	// ✅ endpoint auth
	r.HandleFunc("/auth/register", authHandler.RegisterHandler).Methods("POST")
	r.HandleFunc("/auth/login", authHandler.LoginHandler).Methods("POST")

}
