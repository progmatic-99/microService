package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/progmatic-99/microService/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	d, err := json.Marshal(lp)

	if err != nil {
		http.Error(rw, "Cannot get product list.", http.StatusInternalServerError)
	}

	p.l.Println("Product list returned.")
	rw.Write(d)
}
