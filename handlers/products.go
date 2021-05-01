// Package Classification Product API
//
// Documentation for Product API
//
//  Schemes:   http
//  Base Path: /
//	Version: 1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
//
// swagger:meta
package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"

	currency "github.com/progmatic-99/gRPC/proto/currency"
	"github.com/progmatic-99/microService/data"
)

type Products struct {
	l  *log.Logger
	cc currency.CurrencyClient
}

// A list of products
// swagger:response productsResponse
type productsResponse struct {
	// in: body
	Body []data.Product
}

// swagger:response noContent
type noContent struct {
}

// swagger:parameters deleteProduct
type productIDParam struct {
	// ID of the product to be deleted.
	// in: path
	// required: true
	ID int `json:"id"`
}

func NewProduct(l *log.Logger, cc currency.CurrencyClient) *Products {
	return &Products{l, cc}
}

type KeyProduct struct{}

func (p Products) MiddlewareValidateProduct(next http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		pd := data.Product{}

		err := pd.FromJSON(r.Body)
		if err != nil {
			p.l.Println("Error deserializing product.")
			http.Error(rw, "Error deserializing product.", http.StatusBadRequest)
			return
		}

		// Validating JSON
		err = pd.Validate()
		if err != nil {
			p.l.Println("Error validating product.")
			http.Error(
				rw,
				fmt.Sprintf("Error validating product %s", err),
				http.StatusBadRequest,
			)
			return
		}

		// Add product to ctx
		ctx := context.WithValue(r.Context(), KeyProduct{}, pd)
		r = r.WithContext(ctx)

		next.ServeHTTP(rw, r)
	})
}
