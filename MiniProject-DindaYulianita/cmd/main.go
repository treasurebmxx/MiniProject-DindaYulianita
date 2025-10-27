package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"MiniProject-DindaYulianita/internal/config"
	"MiniProject-DindaYulianita/internal/delivery"
	"MiniProject-DindaYulianita/internal/repository"
	"MiniProject-DindaYulianita/internal/usecase"
)

func main() {
	
	cfg := &config.Config{}
	db := config.NewGormDB()


	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)

	
	r := mux.NewRouter()
	delivery.NewRouter(r, uc, cfg)

	
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // alamat frontend apabila memakai frontend
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	
	handler := c.Handler(r)

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
