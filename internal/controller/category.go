package controller

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Nilesh-Coherent/baston-eventos/internal/model"
	"github.com/Nilesh-Coherent/baston-eventos/internal/service"
	"github.com/Nilesh-Coherent/baston-eventos/internal/util"
	"github.com/Nilesh-Coherent/baston-eventos/internal/web"
)

type CategoryController struct{}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

func (categorycont *CategoryController) AddCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	log.Println("Add Category Service Called.....")
	err := web.RequestParse(r, &category)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	_, category.ID = util.CreateID()
	err = service.AddCategory(&category)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "New Category Added Successfully"})
}

func (categorycont *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	err := web.RequestParse(r, &category)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	params := mux.Vars(r)
	err = util.ValidateIDFormat(params["categoryid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	category.ID, err = util.ParseID(params["categoryid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.UpdateCategory(&category)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "Category Updated Successfully"})
}

func (categorycont *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	params := mux.Vars(r)
	err := util.ValidateIDFormat(params["categoryid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	category.ID, err = util.ParseID(params["categoryid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	err = service.DeleteCategory(&category)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "Category Deleted Successfully"})
}

func (categorycont *CategoryController) GetCategory(w http.ResponseWriter, r *http.Request) {
	var category model.Category
	params := mux.Vars(r)
	categoryid, err := util.ParseID(params["categoryid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.GetCategory(&category, categoryid)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &category)
}

func (categorycont *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
	var categories []model.Category
	err := service.GetCategories(&categories)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &categories)
}
