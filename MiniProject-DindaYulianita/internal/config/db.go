package config

import (
    "fmt"
    "log"

    "gorm.io/driver/mysql"
    "gorm.io/gorm"
    "MiniProject-DindaYulianita/internal/models"
)

func NewGormDB() *gorm.DB {
    dsn := "root:@tcp(127.0.0.1:3306)/mini_db?charset=utf8mb4&parseTime=True&loc=Local"
    db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("❌ Gagal konek ke database:", err)
    }

   
    err = db.AutoMigrate(&models.User{})
    if err != nil {
        log.Fatal("❌ Gagal migrate tabel:", err)
    }

    fmt.Println("✅ Database connected & migrated successfully!")
    return db
}
