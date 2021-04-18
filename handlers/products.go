// Package Classification of Product API
//
// Documentation for Product API
//
//  Schemes:   http
//  Base Path: /
//  Version:   1.0.0
//
//  Consumes:
//  - application/json
//
//  Produces:
//  - application/json
// swagger:meta

package handlers

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/progmatic-99/microService/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)

	p.l.Println("Handle GET products.")

	if err != nil {
		http.Error(rw, "Cannot get product list.", http.StatusInternalServerError)
		return
	}
}

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST prodcuts.")

	pd := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&pd)
}

func (p Products) UpdateProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Cannot convert id.", http.StatusBadRequest)
		return
	}

	p.l.Println("Handle PUT prodcuts.")

	pd := r.Context().Value(KeyProduct{}).(data.Product)

	err = data.UpdateProduct(id, &pd)
	if err == data.ErrProdNotFound {
		http.Error(rw, "Product not found.", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found.", http.StatusInternalServerError)
		return
	}
}

func (p Products) DeleteProduct(rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(rw, "Cannot convert id.", http.StatusBadRequest)
		return
	}

	err = data.DeleteProduct(id)
	if err != nil {
		p.l.Println(err)
		http.Error(rw, "Product not found.", http.StatusBadRequest)
		return
	}

	p.l.Println("Deleted Product with id ", id)
	fmt.Fprintf(rw, "Deleted Product with id %d.\n", id)
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
