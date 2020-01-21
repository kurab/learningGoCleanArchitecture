package main

import (
    "net/http"

    "github.com/julienschmidt/httprouter"

    "api/infrastructure/api/router"
    "api/infrastructure/datastore"
    "api/registry"
)

func main() {
    db := datastore.NewMySQL()
    r := httprouter.New()
    rg := registry.NewInteractor(db)
    h := rg.NewAppHandler()
    router.NewRouter(r, h)

    defer db.Close()

    http.ListenAndServe(":8080", r)
}
