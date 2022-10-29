package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/nilsyadv/baston-eventos/internal/controller"
)

func CreateCategoryRoute(route *mux.Router) {

	// register category
	category := controller.NewCategoryController()
	route.HandleFunc("/category", category.AddCategory).Methods(http.MethodPost)
	route.HandleFunc("/category", category.GetCategories).Methods(http.MethodGet)
	route.HandleFunc("/category/{categoryid}", category.GetCategory).Methods(http.MethodGet)
	route.HandleFunc("/category/{categoryid}", category.UpdateCategory).Methods(http.MethodPut)
	route.HandleFunc("/category/{categoryid}", category.DeleteCategory).Methods(http.MethodDelete)

	// register events
	event := controller.NewEventController()
	route.HandleFunc("/event", event.GetEvents).Methods(http.MethodGet)
	route.HandleFunc("/event", event.AddEvent).Methods(http.MethodPost)
	route.HandleFunc("/event/{eventid}", event.GetEvent).Methods(http.MethodGet)
	route.HandleFunc("/event/{eventid}", event.UpdateEvent).Methods(http.MethodPut)
	route.HandleFunc("/event/{eventid}", event.DeleteEvent).Methods(http.MethodDelete)

	// register payment
	paymntroute := route.PathPrefix("/payment").Subrouter()
	payment := controller.NewPaymentController()
	paymntroute.HandleFunc("/", payment.GetPayments).Methods(http.MethodGet)
	paymntroute.HandleFunc("/", payment.AddPayment).Methods(http.MethodPost)
	paymntroute.HandleFunc("/{paymentid}", payment.GetPayment).Methods(http.MethodGet)
	paymntroute.HandleFunc("/{paymentid}", payment.UpdatePayment).Methods(http.MethodPut)
	paymntroute.HandleFunc("/{paymentid}", payment.DeletePayment).Methods(http.MethodDelete)
	paymntroute.HandleFunc("/event/{eventid}", payment.GetPaymentByEvent).Methods(http.MethodGet)

	payhist := controller.NewPaymentHistoryController()
	paymntroute.HandleFunc("/history", payhist.GetPaymentHistory).Methods(http.MethodGet)
	paymntroute.HandleFunc("/history", payhist.AddPaymentHistory).Methods(http.MethodPost)
	paymntroute.HandleFunc("/history/{histid}", payhist.GetPaymentHistory).Methods(http.MethodGet)
	paymntroute.HandleFunc("/history/{histid}", payhist.UpdatePaymentHistory).Methods(http.MethodPut)
	paymntroute.HandleFunc("/history/{histid}", payhist.DeletePaymentHistory).Methods(http.MethodDelete)
	paymntroute.HandleFunc("/history/payment/{paymentid}", payhist.GetHistoryByPayment).Methods(http.MethodGet)
}
