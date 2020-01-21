package main

import (
    "net/http"

    "github.com/julienschmidt/httprouter"

    "api/infrastructure/api/handler"
    "api/infrastructure/api/router"
    "api/infrastructure/datastore"
    "api/interface/controllers"
    "api/usecase/service"
)

func main() {
    db := datastore.NewMySQL()
    r := httprouter.New()
    rp := datastore.NewUserRepository(db)
    s := service.NewUserService(rp)
    c := controllers.NewUserController(s)
    h := handler.NewUserHandler(c)
    router.NewRouter(r, h)

    defer db.Close()

    http.ListenAndServe(":8080", r)
}
