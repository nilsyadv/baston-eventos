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

func GetPaymentHistory(payhist *model.PaymentHistory, id uuid.UUID) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.Get(uow, payhist, id, []string{})
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get Payment History from db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		uow.RollBack()
		return &er
	}
	uow.Commit()
	return nil
}

func GetPaymentHistories(payhists *[]model.PaymentHistory, conditions ...repository.ConditionalClause) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, true)
	err := repository.GetAll(uow, payhists, conditions)
	if err != nil {
		er := custerror.CreateCustomeError("Failed to get all Payment History from db", err,
			http.StatusInternalServerError)
		uow.RollBack()
		log.Println(er.Error(), er.Message())
		return &er
	}
	uow.Commit()
	return nil
}

func AddPaymentHistory(payhist *model.PaymentHistory) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)

	// validate category
	if ok, err := IsEventValid(&model.Event{}, payhist.PaymentID); !ok {
		er := custerror.CreateCustomeError("invalid payment", err,
			http.StatusBadRequest)
		return &er
	}

	var payment model.Payment
	if err := GetPayment(&payment, payhist.PaymentID); err != nil {
		return err
	}
	payment.PaymentHistory = append(payment.PaymentHistory, *payhist)
	if err := payment.PaymentCalculation(); err != nil {
		er := custerror.CreateCustomeError("invalid payment detail", err,
			http.StatusBadRequest)
		return &er
	}
	if err := UpdatePayment(&payment); err != nil {
		return err
	}

	err := repository.Add(uow, payhist)
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

func UpdatePaymentHistory(payhist *model.PaymentHistory) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)

	// validate category
	if ok, err := IsEventValid(&model.Event{}, payhist.PaymentID); !ok {
		er := custerror.CreateCustomeError("invalid payment", err,
			http.StatusBadRequest)
		return &er
	}

	var payment model.Payment
	if err := GetPayment(&payment, payhist.PaymentID); err != nil {
		return err
	}
	payment.PaymentHistory = append(payment.PaymentHistory, *payhist)
	if err := payment.PaymentCalculation(); err != nil {
		er := custerror.CreateCustomeError("invalid payment detail", err,
			http.StatusBadRequest)
		return &er
	}
	if err := UpdatePayment(&payment); err != nil {
		return err
	}

	err := repository.Update(uow, payhist)
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

func DeletePaymentHistory(payhist *model.PaymentHistory) *custerror.CustomeError {
	uow := repository.NewUnitOfWork(db.DB, false)
	err := repository.Delete(uow, payhist, repository.Filter("id = ?", payhist.ID))
	if err != nil {
		er := custerror.CreateCustomeError("Failed to delete Payment History in db", err,
			http.StatusInternalServerError)
		log.Println(er.Error(), er.Message())
		uow.RollBack()
		return &er
	}
	uow.Commit()
	return nil
}
