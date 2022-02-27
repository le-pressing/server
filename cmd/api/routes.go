package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.HandlerFunc(http.MethodGet, "/status", app.statusHandler)

	router.HandlerFunc(http.MethodGet, "/v1/member/:id", app.getCustomerByID)

	router.HandlerFunc(http.MethodGet, "/v1/orders", app.getAllOrders)
	router.HandlerFunc(http.MethodGet, "/v1/orders/:id", app.getOrderByID)

	return router
}
