package presenters

import (
    "api/domain/model"
    "strconv"
)

type propertyPresenter struct {
}

type PropertyPresenter interface {
    ResponseProperties(props model.Properties) model.Properties
}

func NewPropertyPresenter() PropertyPresenter {
    return &propertyPresenter{}
}

func (pp *propertyPresenter) ResponseProperties(props model.Properties) model.Properties {
    for n, p := range props {
        props[n].WalkTime = p.WalkTime + "分"
        props[n].Area = p.Area + "m2"
        price, _ := strconv.Atoi(p.Price)
        props[n].Price = strconv.Itoa(price/10000) + "万円"
    }
    return props
}
