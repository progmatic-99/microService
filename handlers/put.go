package handlers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/progmatic-99/microService/data"
)

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
