package service

import (
    "api/domain/model"
    "api/usecase/repository"
    "api/usecase/presenter"
)

type PropertyService interface {
    GetAll() (model.Properties, error)
    Get(id int) (model.Property, error)
}

type propertyService struct {
    PropertyRepository repository.PropertyRepository
    PropertyPresenter presenter.PropertyPresenter
}

func NewPropertyService(pr repository.PropertyRepository, pp presenter.PropertyPresenter) PropertyService {
    return &propertyService{PropertyRepository: pr, PropertyPresenter: pp}
}

func (ps *propertyService) GetAll() (model.Properties, error) {
    props, err := ps.PropertyRepository.FindAll()
    if err != nil {
        return nil, err
    }
    return ps.PropertyPresenter.ResponseProperties(props), nil
}

func (ps *propertyService) Get(id int) (model.Property, error) {
    prop, err := ps.PropertyRepository.FindById(id)
    return prop, err
}
