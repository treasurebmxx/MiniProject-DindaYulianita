package repository

import (
    "MiniProject-DindaYulianita/internal/models"
    "gorm.io/gorm"
)

type Repository struct {
    DB *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
    return &Repository{DB: db}
}

func (r *Repository) CreateUser(user *models.User) error {
    return r.DB.Create(user).Error
}

func (r *Repository) GetUserByEmail(email string) (*models.User, error) {
    var user models.User
    if err := r.DB.Where("email = ?", email).First(&user).Error; err != nil {
        return nil, err
    }
    return &user, nil
}
