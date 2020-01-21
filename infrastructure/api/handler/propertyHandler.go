package handler

import (
    "encoding/json"
    "net/http"
    "strconv"

    "github.com/julienschmidt/httprouter"

    "api/interface/controllers"
)

type PropertyHandler interface {
    GetAllProperties(w http.ResponseWriter, r *http.Request, _ httprouter.Params)
    GetProperty(w http.ResponseWriter, r *http.Request, ps httprouter.Params)
}

type propertyHandler struct {
    PropertyController controllers.PropertyController
}

func NewPropertyHandler(pc controllers.PropertyController) PropertyHandler {
    return &propertyHandler{PropertyController: pc}
}

func (ph *propertyHandler) GetAllProperties(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
    props, _ := ph.PropertyController.GetAllProperties()
    json.NewEncoder(w).Encode(props)
}

func (ph *propertyHandler) GetProperty(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    id, _ := strconv.Atoi(ps.ByName("id"))
    prop, _ := ph.PropertyController.GetProperty(id)
    json.NewEncoder(w).Encode(prop)
}
