package server

import (
	"net/http"

	"github.com/gorilla/mux"

	"github.com/Nilesh-Coherent/baston-eventos/internal/controller"
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
	payment := controller.NewPaymentController()
	route.HandleFunc("/payment", payment.GetPayments).Methods(http.MethodGet)
	route.HandleFunc("/payment", payment.AddPayment).Methods(http.MethodPost)
	route.HandleFunc("/payment/{paymentid}", payment.GetPayment).Methods(http.MethodGet)
	route.HandleFunc("/payment/{paymentid}", payment.UpdatePayment).Methods(http.MethodPut)
	route.HandleFunc("/payment/{paymentid}", payment.DeletePayment).Methods(http.MethodDelete)
}
