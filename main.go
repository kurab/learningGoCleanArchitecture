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
)

func connectDB() *gorm.DB {
    connectString := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s%s",
        "user",
        "secret",
        "0.0.0.0",
        "3306",
        "goSample",
        "?charset=utf8&parseTime=True&loc=Local",
    )
    db, err := gorm.Open("mysql", connectString)
    if err != nil {
        panic(err.Error())
    }
    db.LogMode(true)
    db.Set("gorm:table_options", "ENGIN=InnoDB")

    return db
}

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
    db := connectDB()
    r := httprouter.New()
    setRouter(r, db)

    defer db.Close()

    http.ListenAndServe(":8080", r)
}
