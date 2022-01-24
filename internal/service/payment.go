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

func GetPayment(payment *model.Payment, id uuid.UUID) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.Get(uow, payment, id, []string{})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get Payment from db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}

func IsPaymentValid(payment *model.Payment, id uuid.UUID) (bool, *custerror.CustomeError) {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.Get(uow, payment, id, []string{})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get Payment from db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return false, &er
	}

	isempty := &model.Payment{} == payment
	if isempty {
		return false, nil
	}

	return true, nil
}

func GetPayments(payments *[]model.Payment) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.GetAll(uow, payments, []repository.ConditionalClause{})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get all Payment from db", err,
			http.StatusInternalServerError)
		uow.RollBack()
		log.Println(er.Error(), er.Message())
		return &er
	}
	uow.Commit()
	return nil
}

func AddPayment(payment *model.Payment) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)

	// validate category
	if ok, err := IsEventValid(&model.Event{}, payment.EventID); !ok {
		er := custerror.CreateCustomeError("invalid category", err,
			http.StatusBadRequest)
		return &er
	}
	err := repository.Add(uow, payment)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to add New Payment in db", err,
			http.StatusInternalServerError)
		uow.RollBack()
		log.Println(er.Error(), er.Message())
		return &er
	}
	uow.Commit()
	return nil
}

func UpdatePayment(payment *model.Payment) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)

	// validate category
	if ok, err := IsEventValid(&model.Event{}, payment.EventID); !ok {
		er := custerror.CreateCustomeError("invalid event", err,
			http.StatusBadRequest)
		return &er
	}
	err := repository.Update(uow, payment)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to update Payment in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}

func DeletePayment(payment *model.Payment) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Delete(uow, payment, repository.Filter("id = ?", payment.ID))
	if err != nil {
		er := custerror.CreateCustomeError("Failed to delete Payment in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		return &er
	}
	return nil
}
