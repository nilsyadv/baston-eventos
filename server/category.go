package server

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/lin-sel/baston-eventos/internal/controller"
)

func CreateCategoryRoute(route *mux.Router) {
	category := controller.NewCategoryController()
	route.HandleFunc("/category/", category.GetCategory).Methods(http.MethodGet)
	route.HandleFunc("/category/add", category.AddCategory).Methods(http.MethodPost)
	route.HandleFunc("/category/get", category.GetCategories).Methods(http.MethodGet)
	route.HandleFunc("/category/update", category.UpdateCategory).Methods(http.MethodPut)
	route.HandleFunc("/category/delete", category.DeleteCategory).Methods(http.MethodDelete)
}
