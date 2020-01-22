package repository

import "api/domain/model"

type PropertyRepository interface {
    FindAll() (model.Properties, error)
    FindById(id int) (model.Property, error)
    FindByPref(pref_cd int) (model.Properties, error)
}
