package router

import (
    "github.com/julienschmidt/httprouter"

    "api/infrastructure/api/handler"
)

func NewRouter(router *httprouter.Router, handler handler.AppHandler) {
    // USER API
    router.POST("/api/user/register", handler.UserHandler.CreateUser)
    router.GET("/api/user/get", handler.UserHandler.GetAllUsers)
    router.GET("/api/user/get/:id", handler.UserHandler.GetUser)
}
