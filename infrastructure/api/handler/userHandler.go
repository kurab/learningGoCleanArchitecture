package handler

import (
    "encoding/json"
    "fmt"
    "net/http"
    "strconv"

    "github.com/julienschmidt/httprouter"

    "api/domain/model"
    "api/infrastructure/api/validator"
    "api/interface/controllers"
)

type UserHandler interface {
    CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
    GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
    GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type userHandler struct {
    UserController controllers.UserController
    Validator      validation.CustomValidator
}

func NewUserHandler(uc controllers.UserController, v validation.CustomValidator) UserHandler {
    return &userHandler{UserController: uc, Validator: v}
}

func (uh *userHandler) CreateUser(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    data := &model.User{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        fmt.Fprintln(w, "Bad request: "+err.Error())
        return
    }
    if err := uh.Validator.Validate(data); err != nil {
        fmt.Fprintln(w, "Validation error: "+err.Error())
        return
    }

    err := uh.UserController.CreateUser(data)
    if err == false {
        fmt.Fprintln(w, "registerd")
    }
}

func (uh *userHandler) GetAllUsers(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    u, _ := uh.UserController.GetAllUsers()
    json.NewEncoder(w).Encode(u)
}

func (uh *userHandler) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, _ := strconv.Atoi(ps.ByName("id"))
    u, _ := uh.UserController.GetUser(id)
    json.NewEncoder(w).Encode(u)
}
