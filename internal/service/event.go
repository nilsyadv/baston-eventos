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

func GetEvent(event *model.Event, id uuid.UUID) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.Get(uow, event, id, []string{"PaymentDetails", "PaymentDetails.PaymentHistory"})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get Event from db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		uow.RollBack()
		return &er
	}
	uow.Commit()
	return nil
}

func IsEventValid(category *model.Event, id uuid.UUID) (bool, *custerror.CustomeError) {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.Get(uow, category, id, []string{})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get Event from db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return false, &er
	}

	isempty := &model.Event{} == category
	if isempty {
		return false, nil
	}

	return true, nil
}

func GetEvents(events *[]model.Event) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.GetAll(uow, events, []repository.ConditionalClause{
		repository.PreloadAssociations([]string{"PaymentDetails", "PaymentDetails.PaymentHistory"}),
	})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get all Event from db", err,
			http.StatusInternalServerError)
		uow.RollBack()
		log.Println(er.Error(), er.Message())
		return &er
	}
	uow.Commit()
	return nil
}

func AddEvent(event *model.Event) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)

	// validate category
	if ok, err := IsCategoryValid(&model.Category{}, event.CategoryID); !ok {
		er := custerror.CreateCustomeError("invalid category", err,
			http.StatusBadRequest)
		return &er
	}

	if err := event.PaymentDetails.PaymentCalculation(); err != nil {
		er := custerror.CreateCustomeError("invalid event detail", err,
			http.StatusBadRequest)
		return &er
	}

	err := repository.Add(uow, event)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to add New Event in db", err,
			http.StatusInternalServerError)
		uow.RollBack()
		log.Println(er.Error(), er.Message())
		return &er
	}
	uow.Commit()
	return nil
}

func UpdateEvent(event *model.Event) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)

	// validate category
	if ok, err := IsCategoryValid(&model.Category{}, event.CategoryID); !ok {
		er := custerror.CreateCustomeError("invalid category", err,
			http.StatusBadRequest)
		return &er
	}

	if err := event.PaymentDetails.PaymentCalculation(); err != nil {
		er := custerror.CreateCustomeError("invalid event detail", err,
			http.StatusBadRequest)
		return &er
	}

	err := repository.Update(uow, event)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to update Event in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		uow.RollBack()
		return &er
	}
	uow.Commit()
	return nil
}

func DeleteEvent(event *model.Event) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Delete(uow, event, repository.Filter("id = ?", event.ID))
	if err != nil {
		er := custerror.CreateCustomeError("Failed to delete Event in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		uow.RollBack()
		return &er
	}
	uow.Commit()
	return nil
}
