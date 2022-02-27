package main

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"
	"github.com/le-pressing/server/models"
)

func (app *application) getCustomerByID(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))
	if err != nil {
		app.logger.Print(errors.New("Invalid Customer ID Parameter"))
		app.errorJSON(w, err)
		return
	}

	app.logger.Println("Customer ID is:", id)

	customer := models.Customer{
		ID:             id,
		FirstName:      "",
		LastName:       "",
		BuildingNumber: 000,
		StreetName:     "",
		City:           "",
		State:          "",
		ZipCode:        10017,
		Orders:         []models.Order{},
		CreatedAt:      time.Now(),
		UpdatedAt:      time.Now(),
	}

	err = app.writeJSON(w, http.StatusOK, customer, "Customer")
}
