package service

import (
	"log"
	"net/http"

	custerror "github.com/lin-sel/baston-eventos/error"
	"github.com/lin-sel/baston-eventos/internal/db"
	"github.com/lin-sel/baston-eventos/internal/model"
	"github.com/lin-sel/baston-eventos/internal/repository"
	uuid "github.com/satori/go.uuid"
)

func GetCategory(category *model.ECategory, id uuid.UUID) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.Get(uow, category, id, []string{})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get Category from db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}

func GetCategories(categories *[]model.ECategory) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.GetAll(uow, categories, []repository.ConditionalClause{})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get all Category from db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}

func AddCategory(category *model.ECategory) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Add(uow, category)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to add New Category in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}

func UpdateCategory(category *model.ECategory) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Update(uow, category)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to update Category in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}

func DeleteCategory(category *model.ECategory) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Delete(uow, category, repository.Filter("category_id = ?", category.CategoryID))
	if err != nil {
		er := custerror.CreateCustomeError("Failed to delete Category in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}
