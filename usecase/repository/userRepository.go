package repository

import "api/domain/model"

type UserRepository interface {
    Store(data *model.User) bool
    FindAll() (model.Users, error)
    FindById(id int) (model.User, error)
}
