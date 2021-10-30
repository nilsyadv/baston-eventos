package controller

type CategoryController struct{}

func NewCategoryController() *CategoryController {
	return &CategoryController{}
}

// func (categorycont *CategoryController) AddCategory(w http.ResponseWriter, r *http.Request) {
// 	var category model.ECategory
// 	err := web.RequestParse(r, &category)
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}
// 	category.CategoryID = util.CreateID()
// 	err = service.AddCategory(&category)
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}
// 	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "New Category Added Successfully"})
// }

// func (categorycont *CategoryController) UpdateCategory(w http.ResponseWriter, r *http.Request) {
// 	var category model.ECategory
// 	err := web.RequestParse(r, &category)
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}

// 	params := mux.Vars(r)
// 	err = util.ValidateIDFormat(params["categoryid"])
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}

// 	category.CategoryID = params["categoryid"]
// 	err = service.UpdateCategory(&category)
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}
// 	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "Category Updated Successfully"})
// }

// func (categorycont *CategoryController) DeleteCategory(w http.ResponseWriter, r *http.Request) {
// 	var category model.ECategory
// 	params := mux.Vars(r)
// 	err := util.ValidateIDFormat(params["categoryid"])
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}

// 	category.CategoryID = params["categoryid"]
// 	err = service.DeleteCategory(&category)
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}
// 	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "Category Deleted Successfully"})
// }

// func (categorycont *CategoryController) GetCategory(w http.ResponseWriter, r *http.Request) {
// 	var category model.ECategory
// 	params := mux.Vars(r)
// 	categoryid, err := util.ParseID(params["categoryid"])
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}

// 	err = service.GetCategory(&category, categoryid)
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}
// 	web.RespondJSON(w, http.StatusOK, &category)
// }

// func (categorycont *CategoryController) GetCategories(w http.ResponseWriter, r *http.Request) {
// 	var categories []model.ECategory
// 	err := service.GetCategories(&categories)
// 	if err != nil {
// 		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
// 		return
// 	}
// 	web.RespondJSON(w, http.StatusOK, &categories)
// }
