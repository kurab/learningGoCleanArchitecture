package controllers

import (
    "api/domain/model"
    "api/usecase/service"
)

type PropertyController interface {
    GetAllProperties() (model.Properties, error)
    GetProperty(id int) (model.Property, error)
}

type propertyController struct {
    PropertyService service.PropertyService
}

func NewPropertyController(ps service.PropertyService) PropertyController {
    return &propertyController{PropertyService: ps}
}

func (pc *propertyController) GetAllProperties() (model.Properties, error) {
    props, err := pc.PropertyService.GetAll()
    if err != nil {
        return nil, err
    }
    return props, nil
}

func (pc *propertyController) GetProperty(id int) (model.Property, error) {
    prop, err := pc.PropertyService.Get(id)
    return prop, err
}
