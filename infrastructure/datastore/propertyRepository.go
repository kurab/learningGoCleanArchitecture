package datastore

import (
    "github.com/jinzhu/gorm"

    "api/domain/model"
)

type PropertyRepository interface {
    FindAll() (model.Properties, error)
    FindById(id int) (model.Property, error)
}

type propertyRepository struct {
    db *gorm.DB
}

func NewPropertyRepository(db *gorm.DB) PropertyRepository {
    return &propertyRepository{db: db}
}

func (pr *propertyRepository) FindAll() (model.Properties, error) {
    props := model.Properties{}
    err := pr.db.Where("flg_open = ?", true).Order("pref_cd").Order("price").Find(&props).Error
    if err != nil {
        return nil, err
    }
    return props, nil
}

func (pr *propertyRepository) FindById(id int) (model.Property, error) {
    prop := model.Property{}
    err := pr.db.Where("id = ?", id).First(&prop).Error
    return prop, err
}
