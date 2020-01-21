package presenter

import "api/domain/model"

type PropertyPresenter interface {
    ResponseProperties(props model.Properties) model.Properties
}
