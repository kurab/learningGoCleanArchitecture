package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/jinzhu/gorm"
    "github.com/julienschmidt/httprouter"

    "api/domain/model"
)

type UserHandler interface {
    CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
    GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
    GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type userHandler struct {
    db *gorm.DB
}

func NewUserHandler(db *gorm.DB) UserHandler {
    return &userHandler{db: db}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    data := model.User{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        fmt.Fprintln(w, "Bad request: "+err.Error())
        return
    }
    uh.db.NewRecord(data)
    uh.db.Create(&data)
    if uh.db.NewRecord(data) == false {
        fmt.Fprintln(w, "registerd")
    }
}

func (uh *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    users := model.Users{}
    uh.db.Find(&users)
    json.NewEncoder(w).Encode(users)
}

func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, _ := strconv.Atoi(ps.ByName("id"))
    user := model.User{}
    uh.db.Where("Id = ?", id).First(&user)
    json.NewEncoder(w).Encode(user)
}
