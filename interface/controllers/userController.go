package controllers

import (
    "github.com/jinzhu/gorm"

    "api/domain/model"
)

type UserController interface {
    CreateUser(data *model.User) bool
    GetAllUsers() (model.Users, error)
    GetUser(id int) (model.User, error)
}

type userController struct {
    db *gorm.DB
}

func NewUserController(db *gorm.DB) UserController {
    return &userController{db: db}
}

func (uc *userController) CreateUser(data *model.User) bool {
    uc.db.NewRecord(data)
    uc.db.Create(&data)
    return uc.db.NewRecord(data)
}

func (uc *userController) GetAllUsers() (model.Users, error) {
    users := model.Users{}
    err := uc.db.Find(&users).Error
    if err != nil {
        return nil, err
    }
    return users, nil
}

func (uc *userController) GetUser(id int) (model.User, error) {
    user := model.User{}
    err := uc.db.Where("Id = ?", id).First(&user).Error
    return user, err
}
