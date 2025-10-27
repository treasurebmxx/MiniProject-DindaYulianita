package handlers

import (
    "encoding/json"
    "net/http"
    "MiniProject-DindaYulianita/internal/usecase"
    "MiniProject-DindaYulianita/internal/models"
)

type AuthHandler struct {
    uc *usecase.Usecase
}

func NewAuthHandler(uc *usecase.Usecase) *AuthHandler {
    return &AuthHandler{uc: uc}
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    if err := h.uc.Register(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Registrasi berhasil",
        "user":    user,
    })
}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
    var data struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    user, err := h.uc.Login(data.Email, data.Password)
    if err != nil {
        http.Error(w, err.Error(), http.StatusUnauthorized)
        return
    }

    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "Login berhasil",
        "user":    user,
    })
}
