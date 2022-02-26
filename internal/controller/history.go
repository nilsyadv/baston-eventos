package controller

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Nilesh-Coherent/baston-eventos/internal/model"
	"github.com/Nilesh-Coherent/baston-eventos/internal/service"
	"github.com/Nilesh-Coherent/baston-eventos/internal/util"
	"github.com/Nilesh-Coherent/baston-eventos/internal/web"
	"github.com/Nilesh-Coherent/baston-eventos/log"
	"github.com/Nilesh-Coherent/common-service-evnt/pkg/repository"
)

type PaymentHistoryController struct{}

func NewPaymentHistoryController() *PaymentHistoryController {
	return &PaymentHistoryController{}
}

func (payhistcont *PaymentHistoryController) AddPaymentHistory(w http.ResponseWriter, r *http.Request) {
	var payhist model.PaymentHistory
	log.Log.Println("Add Payment Service Called.....")
	err := web.RequestParse(r, &payhist)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	err = payhist.ValidatePayHist()
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.AddPaymentHistory(&payhist)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, map[string]interface{}{"message": "New Payment Added Successfully"})
}

func (payhistcont *PaymentHistoryController) UpdatePaymentHistory(w http.ResponseWriter, r *http.Request) {
	var payhist model.PaymentHistory
	err := web.RequestParse(r, &payhist)
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

	payhist.ID, err = util.ParseID(params["payhistid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = payhist.ValidatePayHist()
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.UpdatePaymentHistory(&payhist)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK,
		map[string]interface{}{"message": "Payment History Updated Successfully"})
}

func (paymentcont *PaymentHistoryController) DeletePaymentHistory(w http.ResponseWriter, r *http.Request) {
	var payhist model.PaymentHistory
	params := mux.Vars(r)
	err := util.ValidateIDFormat(params["payhistid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	payhist.ID, err = util.ParseID(params["paymentid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	err = service.DeletePaymentHistory(&payhist)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK,
		map[string]interface{}{"message": "Payment Deleted Successfully"})
}

func (paymentcont *PaymentHistoryController) GetPaymentHistory(w http.ResponseWriter, r *http.Request) {
	var payhist model.PaymentHistory
	params := mux.Vars(r)
	paymentid, err := util.ParseID(params["payhistid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.GetPaymentHistory(&payhist, paymentid)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &payhist)
}

func (paymentcont *PaymentHistoryController) GetHistoryByPayment(w http.ResponseWriter, r *http.Request) {
	var payhists []model.PaymentHistory
	params := mux.Vars(r)
	eventid, err := util.ParseID(params["paymentid"])
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}

	err = service.GetPaymentHistories(&payhists, repository.Filter("payment_id = ?", eventid))
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &payhists)
}

func (paymentcont *PaymentHistoryController) GetPaymentHistories(w http.ResponseWriter, r *http.Request) {
	var payhists []model.PaymentHistory
	err := service.GetPaymentHistories(&payhists)
	if err != nil {
		web.RespondErrorMessage(w, err.ResponseCode, err.Message())
		return
	}
	web.RespondJSON(w, http.StatusOK, &payhists)
}
