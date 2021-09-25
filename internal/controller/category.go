package controller

import "net/http"

type CategoryController struct{}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (category *CategoryController) AddCategory(w http.ResponseWriter, r *http.Request) {

}
