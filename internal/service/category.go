package service

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	custerror "github.com/Nilesh-Coherent/baston-eventos/error"
	"github.com/Nilesh-Coherent/baston-eventos/internal/db"
	"github.com/Nilesh-Coherent/baston-eventos/internal/model"
	"github.com/Nilesh-Coherent/common-service-evnt/pkg/repository"
)

func GetCategory(category *model.Category, id uuid.UUID) *custerror.CustomeError {
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

func GetCategories(categories *[]model.Category) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.GetAll(uow, categories, []repository.ConditionalClause{})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get all Category from db", err,
			http.StatusInternalServerError)
		uow.RollBack()
		log.Println(er.Error(), er.Message())
		return &er
	}
	uow.Commit()
	return nil
}

func AddCategory(category *model.Category) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Add(uow, category)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to add New Category in db", err,
			http.StatusInternalServerError)
		uow.RollBack()
		log.Println(er.Error(), er.Message())
		return &er
	}
	uow.Commit()
	return nil
}

func UpdateCategory(category *model.Category) *custerror.CustomeError {
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

func DeleteCategory(category *model.Category) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Delete(uow, category, repository.Filter("id = ?", category.ID))
	if err != nil {
		er := custerror.CreateCustomeError("Failed to delete Category in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}
