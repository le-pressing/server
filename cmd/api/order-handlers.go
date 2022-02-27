package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/le-pressing/server/models"
)

func (app *application) getOrderByID(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("Invalid Order ID Parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("Order ID is:", id)

	order := models.Order{
		ID:         id,
		CustomerID: 0000,
		Items:      []models.OrderItem{},
		Status:     "",
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, order, "Order")
}

func (app *application) getAllOrders(w http.ResponseWriter, r *http.Request) {
	//
}
