package handlers

import (
	"net/http"

	"github.com/progmatic-99/microService/data"
)

func (p *Products) AddProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST prodcuts.")

	pd := r.Context().Value(KeyProduct{}).(data.Product)

	data.AddProduct(&pd)
}
