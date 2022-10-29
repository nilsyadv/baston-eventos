package service

import (
	"log"
	"net/http"

	uuid "github.com/satori/go.uuid"

	custerror "github.com/nilsyadv/baston-eventos/error"
	"github.com/nilsyadv/baston-eventos/internal/db"
	"github.com/nilsyadv/baston-eventos/internal/model"
	"github.com/nilsyadv/common-service-evnt/pkg/repository"
)

func GetPayment(payment *model.Payment, id uuid.UUID) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.Get(uow, payment, id, []string{"PaymentHistory"})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get Payment from db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		uow.RollBack()
		return &er
	}
	uow.Commit()
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

func GetPayments(payments *[]model.Payment, conditions ...repository.ConditionalClause) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	conditions = append(conditions, repository.PreloadAssociations([]string{"PaymentHistory"}))
	err := repository.GetAll(uow, payments, conditions)
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

	if err := payment.PaymentCalculation(); err != nil {
		er := custerror.CreateCustomeError("invalid payment detail", err,
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

	if err := payment.PaymentCalculation(); err != nil {
		er := custerror.CreateCustomeError("invalid payment detail", err,
			http.StatusBadRequest)
		return &er
	}

	err := repository.Update(uow, payment)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to update Payment in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		uow.RollBack()
		return &er
	}
	uow.Commit()
	return nil
}

func DeletePayment(payment *model.Payment) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Delete(uow, payment, repository.Filter("id = ?", payment.ID))
	if err != nil {
		er := custerror.CreateCustomeError("Failed to delete Payment in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		uow.RollBack()
		return &er
	}
	uow.Commit()
	return nil
}
