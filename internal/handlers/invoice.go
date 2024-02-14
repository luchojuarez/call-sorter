package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/luchojuarez/call-sorter/internal/domain/invoice"
)

type InvoiceProcessor interface {
	Generate(ctx context.Context, phoneNumber string, month time.Month, year int) (invoice.Model, error)
}
type InvoiceHandler struct {
	processor InvoiceProcessor
}

func NewHandler(p InvoiceProcessor) InvoiceHandler {
	return InvoiceHandler{
		processor: p,
	}
}

func (h InvoiceHandler) Generate(w http.ResponseWriter, req *http.Request) {
	phoneNumber, year, month, err := getParams(req)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(getErrorResponseBody(err))

		return
	}

	invoiceModel, err := h.processor.Generate(context.Background(), phoneNumber, month, year)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(w).Encode(getErrorResponseBody(err))

		return
	}

	w.WriteHeader(http.StatusOK)

	json.NewEncoder(w).Encode(invoiceModel)

}

func getParams(req *http.Request) (phoneNumber string, year int, month time.Month, err error) {
	vars := mux.Vars(req)
	yearStr := vars["year"]
	if yearStr != "" {
		y, e := strconv.Atoi(yearStr)
		if e != nil {
			err = fmt.Errorf("invalid year: (%w) '%+v'", e, yearStr)
			return
		}
		year = y
	} else {
		err = ErrInvalidYear
		return
	}

	monthStr := vars["month"]
	if monthStr != "" {
		m, e := strconv.Atoi(monthStr)
		if e != nil {
			err = fmt.Errorf("A Month specifies a month of the year (January = 1, ...) %w", e)
			return
		}
		month = time.Month(m)
	} else {
		err = ErrInvalidMonth
		return
	}

	phoneNumber = req.Header.Get("phone_number")
	if phoneNumber == "" {
		err = ErrInvalidPhoneNumber
		return
	}

	return
}
