package usecase

import (
    "errors"
    "MiniProject-DindaYulianita/internal/repository"
    "MiniProject-DindaYulianita/internal/models"
    "golang.org/x/crypto/bcrypt"
)

type Usecase struct {
    repo *repository.Repository
}

func NewUsecase(r *repository.Repository) *Usecase {
    return &Usecase{repo: r}
}

// =================== REGISTER ===================
func (u *Usecase) Register(user *models.User) error {
    // untuk mengecek apakah email sudah digunakan
    existingUser, _ := u.repo.GetUserByEmail(user.Email)
    if existingUser != nil {
        return errors.New("email sudah digunakan")
    }

    // hash password
    hashed, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
    if err != nil {
        return err
    }
    user.Password = string(hashed)

    return u.repo.CreateUser(user)
}

// =================== LOGIN ===================
func (u *Usecase) Login(email, password string) (*models.User, error) {
    user, err := u.repo.GetUserByEmail(email)
    if err != nil {
        return nil, errors.New("user tidak ditemukan")
    }

    // cek password
    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
        return nil, errors.New("password salah")
    }

    return user, nil
}
