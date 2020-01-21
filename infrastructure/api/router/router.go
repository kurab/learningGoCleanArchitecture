package router

import (
    "github.com/julienschmidt/httprouter"

    "api/infrastructure/api/handler"
)

func NewRouter(router *httprouter.Router, handler handler.UserHandler) {
    // USER API
    router.POST("/api/user/register", handler.CreateUser)
    router.GET("/api/user/get", handler.GetAllUsers)
    router.GET("/api/user/get/:id", handler.GetUser)
}
