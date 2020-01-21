package registry

import (
    "github.com/jinzhu/gorm"

    "api/infrastructure/api/handler"
    "api/infrastructure/datastore"
    "api/interface/controllers"
    "api/usecase/service"
)

type interactor struct {
    db *gorm.DB
}

type Interactor interface {
    NewAppHandler() handler.AppHandler
}

func NewInteractor(db *gorm.DB) Interactor {
    return &interactor{db: db}
}

func (i *interactor) NewAppHandler() handler.AppHandler {
    return handler.AppHandler{UserHandler: i.NewUserHandler()}
}

func (i *interactor) NewUserHandler() handler.UserHandler {
    return handler.NewUserHandler(i.NewUserController())
}

func (i *interactor) NewUserController() controllers.UserController {
    return controllers.NewUserController(i.NewUserService())
}

func (i *interactor) NewUserService() service.UserService {
    return service.NewUserService(i.NewUserRepository())
}

func (i *interactor) NewUserRepository() datastore.UserRepository {
    return datastore.NewUserRepository(i.db)
}
