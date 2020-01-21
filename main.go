package main

import (
    "net/http"

    "github.com/julienschmidt/httprouter"

    "api/infrastructure/api/handler"
    "api/infrastructure/api/router"
    "api/infrastructure/datastore"
)

func main() {
    db := datastore.NewMySQL()
    r := httprouter.New()
    h := handler.NewUserHandler(db)
    router.NewRouter(r, h)

    defer db.Close()

    http.ListenAndServe(":8080", r)
}
