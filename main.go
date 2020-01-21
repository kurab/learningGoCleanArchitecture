package main

import (
    "net/http"

    "github.com/julienschmidt/httprouter"

    "api/infrastructure/api/handler"
    "api/infrastructure/api/router"
    "api/infrastructure/datastore"
    "api/interface/controllers"
)

func main() {
    db := datastore.NewMySQL()
    r := httprouter.New()
    c := controllers.NewUserController(db)
    h := handler.NewUserHandler(c)
    router.NewRouter(r, h)

    defer db.Close()

    http.ListenAndServe(":8080", r)
}
