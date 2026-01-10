package handler

import (
    "encoding/json"
    "net/http"
    "simple-api/internal/model"
    "simple-api/internal/service"
    "simple-api/pkg/response"
)

type UserHandler struct {
    service *service.UserService
}

func NewUserHandler(service *service.UserService) *UserHandler {
    return &UserHandler{service: service}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        response.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    users := h.service.GetAllUsers()
    response.JSON(w, users, http.StatusOK)
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        response.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
        return
    }

    var user model.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        response.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    newUser := h.service.AddUser(user)
    response.JSON(w, newUser, http.StatusCreated)
}