package service

import (
    "github.com/jinzhu/gorm"

    "api/domain/model"
)

type UserService interface {
    Create(data *model.User) bool
    GetAll(users model.Users) (model.Users, error)
    Get(id int) (model.User, error)
}

type userService struct {
    db *gorm.DB
}

func NewUserService(db *gorm.DB) UserService {
    return &userService{db: db}
}

func (us *userService) Create(data *model.User) bool {
    us.db.NewRecord(data)
    us.db.Create(&data)
    return us.db.NewRecord(data)
}

func (us *userService) GetAll(users model.Users) (model.Users, error) {
    err := us.db.Find(&users).Error
    if err != nil {
        return nil, err
    }
    return users, nil
}

func (us *userService) Get(id int) (model.User, error) {
    user := model.User{}
    err := us.db.Where("Id = ?", id).First(&user).Error
    return user, err
}
