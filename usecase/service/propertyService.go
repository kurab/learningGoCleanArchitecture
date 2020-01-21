package service

import (
    "api/domain/model"
    "api/usecase/repository"
)

type PropertyService interface {
    GetAll() (model.Properties, error)
    Get(id int) (model.Property, error)
}

type propertyService struct {
    PropertyRepository repository.PropertyRepository
}

func NewPropertyService(pr repository.PropertyRepository) PropertyService {
    return &propertyService{PropertyRepository: pr}
}

func (ps *propertyService) GetAll() (model.Properties, error) {
    props, err := ps.PropertyRepository.FindAll()
    if err != nil {
        return nil, err
    }
    return props, nil
}

func (ps *propertyService) Get(id int) (model.Property, error) {
    prop, err := ps.PropertyRepository.FindById(id)
    return prop, err
}
