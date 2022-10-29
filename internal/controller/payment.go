package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nilsyadv/baston-eventos/internal/model"
	"github.com/nilsyadv/baston-eventos/internal/service"
	"github.com/nilsyadv/baston-eventos/internal/util"
	"github.com/nilsyadv/baston-eventos/internal/web"
	"github.com/nilsyadv/baston-eventos/log"
	"github.com/nilsyadv/common-service-evnt/pkg/repository"
)

type PaymentController struct{}

func NewPaymentController() *PaymentController {
	return &PaymentController{}
}

func (paymentcont *PaymentController) AddPayment(w http.ResponseWriter, r *http.Request) {
	var payment model.Payment
	log.Log.Println("Add Payment Service Called.....")
	err := web.RequestParse(r, &payment)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	_, payment.ID = util.CreateID()
	err = service.AddPayment(&payment)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "New Payment Added Successfully"})
}

func (paymentcont *PaymentController) UpdatePayment(w http.ResponseWriter, r *http.Request) {
	var payment model.Payment
	err := web.RequestParse(r, &payment)
	if err != nil {
		log.Log.Printf("Error: %s", err.Error())
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	params := mux.Vars(r)
	err = util.ValidateIDFormat(params["paymentid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	payment.ID, err = util.ParseID(params["paymentid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.UpdatePayment(&payment)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "Payment Updated Successfully"})
}

func (paymentcont *PaymentController) DeletePayment(w http.ResponseWriter, r *http.Request) {
	var payment model.Payment
	params := mux.Vars(r)
	err := util.ValidateIDFormat(params["paymentid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	payment.ID, err = util.ParseID(params["paymentid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	err = service.DeletePayment(&payment)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "Payment Deleted Successfully"})
}

func (paymentcont *PaymentController) GetPayment(w http.ResponseWriter, r *http.Request) {
	var payment model.Payment
	params := mux.Vars(r)
	paymentid, err := util.ParseID(params["paymentid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.GetPayment(&payment, paymentid)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &payment)
}

func (paymentcont *PaymentController) GetPaymentByEvent(w http.ResponseWriter, r *http.Request) {
	var payments []model.Payment
	params := mux.Vars(r)
	eventid, err := util.ParseID(params["eventid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.GetPayments(&payments, repository.Filter("event_id = ?", eventid))
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &payments)
}

func (paymentcont *PaymentController) GetPayments(w http.ResponseWriter, r *http.Request) {
	var payments []model.Payment
	err := service.GetPayments(&payments)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &payments)
}
