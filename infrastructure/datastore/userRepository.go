package datastore

import (
    "github.com/jinzhu/gorm"

    "api/domain/model"
)

type UserRepository interface {
    Store(data *model.User) bool
    FindAll() (model.Users, error)
    FindById(id int) (model.User, error)
}

type userRepository struct {
    db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
    return &userRepository{db: db}
}

func (ur *userRepository) Store(data *model.User) bool {
    ur.db.NewRecord(data)
    ur.db.Create(&data)
    return ur.db.NewRecord(data)
}

func (ur *userRepository) FindAll() (model.Users, error) {
    users := model.Users{}
    err := ur.db.Find(&users).Error
    if err != nil {
        return nil, err
    }
    return users, nil
}

func (ur *userRepository) FindById(id int) (model.User, error) {
    user := model.User{}
    err := ur.db.Where("Id = ?", id).First(&user).Error
    return user, err
}
