package main

import (
    "net/http"

    "github.com/julienschmidt/httprouter"

    "api/infrastructure/api/router"
    "api/infrastructure/datastore"
)

func main() {
    db := datastore.NewMySQL()
    r := httprouter.New()
    router.NewRouter(r, db)

    defer db.Close()

    http.ListenAndServe(":8080", r)
}
