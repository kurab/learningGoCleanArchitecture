package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/julienschmidt/httprouter"

    "api/domain/model"
    "api/interface/controllers"
)

type UserHandler interface {
    CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
    GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
    GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type userHandler struct {
    userController controllers.UserController
}

func NewUserHandler(uc controllers.UserController) UserHandler {
    return &userHandler{userController: uc}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    data := &model.User{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        fmt.Fprintln(w, "Bad request: "+err.Error())
        return
    }
    err := uh.userController.CreateUser(data)
    if err == false {
        fmt.Fprintln(w, "registerd")
    }
}

func (uh *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    u, _ := uh.userController.GetAllUsers()
    json.NewEncoder(w).Encode(u)
}

func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, _ := strconv.Atoi(ps.ByName("id"))
    u, _ := uh.userController.GetUser(id)
    json.NewEncoder(w).Encode(u)
}
