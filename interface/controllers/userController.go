package controllers

import (
    "api/domain/model"
    "api/usecase/service"
)

type UserController interface {
    CreateUser(data *model.User) bool
    GetAllUsers() (model.Users, error)
    GetUser(id int) (model.User, error)
}

type userController struct {
    UserService service.UserService
}

func NewUserController(us service.UserService) UserController {
    return &userController{UserService: us}
}

func (uc *userController) CreateUser(data *model.User) bool {
    return uc.UserService.Create(data)
}

func (uc *userController) GetAllUsers() (model.Users, error) {
    us, err := uc.UserService.GetAll()
    if err != nil {
        return nil, err
    }
    return us, nil
}

func (uc *userController) GetUser(id int) (model.User, error) {
    u, err := uc.UserService.Get(id)
    return u, err
}
