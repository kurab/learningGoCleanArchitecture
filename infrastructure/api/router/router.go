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

    // PROPERTY API
    router.GET("/api/property/get", handler.PropertyHandler.GetAllProperties)
    router.GET("/api/property/get/:id", handler.PropertyHandler.GetProperty)
    router.GET("/api/property/pref/get/:pref_cd", handler.PropertyHandler.GetPropertiesByPref)
}
