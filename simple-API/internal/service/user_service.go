package service

import "simple-api/internal/model"

type UserService struct {
    users []model.User
}

func NewUserService() *UserService {
    return &UserService{
        users: []model.User{
            {ID: 1, Name: "John Doe", Email: "john@example.com"},
            {ID: 2, Name: "Jane Smith", Email: "jane@example.com"},
        },
    }
}

func (s *UserService) GetAllUsers() []model.User {
    return s.users
}

func (s *UserService) AddUser(user model.User) model.User {
    user.ID = len(s.users) + 1
    s.users = append(s.users, user)
    return user
}