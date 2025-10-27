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
	// 1️⃣ Inisialisasi konfigurasi dan database
	cfg := &config.Config{} // kalau kamu punya isi struct config (misalnya ENV), isi di sini
	db := config.NewGormDB()

	// 2️⃣ Inisialisasi repository dan usecase
	repo := repository.NewRepository(db)
	uc := usecase.NewUsecase(repo)

	// 3️⃣ Setup router pakai gorilla/mux
	r := mux.NewRouter()
	delivery.NewRouter(r, uc, cfg)

	// 4️⃣ Tambahkan CORS supaya bisa diakses dari React (localhost:5173)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:5173"}, // asal frontend kamu
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	// 5️⃣ Jalankan server dengan middleware CORS
	handler := c.Handler(r)

	fmt.Println("🚀 Server running at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
