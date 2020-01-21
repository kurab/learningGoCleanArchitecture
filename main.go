package main

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/mysql"
    "github.com/julienschmidt/httprouter"

    "api/domain/model"
    "api/infrastructure/datastore"
)

func setRouter(router *httprouter.Router, db *gorm.DB) {
    // USER API
    router.POST("/api/user/register", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        data := model.User{}
        if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
            fmt.Fprintln(w, "Bad request: "+err.Error())
            return
        }
        db.NewRecord(data)
        db.Create(&data)
        if db.NewRecord(data) == false {
            fmt.Fprintln(w, "registerd")
        }
    })
    router.GET("/api/user/get", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
        users := model.Users{}
        db.Find(&users)
        json.NewEncoder(w).Encode(users)
    })
    router.GET("/api/user/get/:id", func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
        id, _ := strconv.Atoi(ps.ByName("id"))
        user := model.User{}
        db.Where("Id = ?", id).First(&user)
        json.NewEncoder(w).Encode(user)
    })
}

func main() {
    db := datastore.NewMySQL()
    r := httprouter.New()
    setRouter(r, db)

    defer db.Close()

    http.ListenAndServe(":8080", r)
}
