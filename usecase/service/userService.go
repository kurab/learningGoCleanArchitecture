package service

import (
    "api/domain/model"
    "api/usecase/repository"
)

type UserService interface {
    Create(data *model.User) bool
    GetAll() (model.Users, error)
    Get(id int) (model.User, error)
}

type userService struct {
    UserRepository repository.UserRepository
}

func NewUserService(ur repository.UserRepository) UserService {
    return &userService{UserRepository: ur}
}

func (us *userService) Create(data *model.User) bool {
    return us.UserRepository.Store(data)
}

func (us *userService) GetAll() (model.Users, error) {
    u, err := us.UserRepository.FindAll()
    if err != nil {
        return nil, err
    }
    return u, nil
}

func (us *userService) Get(id int) (model.User, error) {
    u, err := us.UserRepository.FindById(id)
    return u, err
}
