package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/progmatic-99/microService/data"
)

type Products struct {
	l *log.Logger
}

func NewProduct(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProduct(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.addProducts(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// expect the ID in URI
		p.l.Println("PUT", r.URL.Path)

		reg := regexp.MustCompile("/([0-9]+)")
		// p.l.Println("Regex: ", reg)
		g := reg.FindAllStringSubmatch(r.URL.Path, -1)
		// p.l.Println(g[0][1])

		if len(g) != 1 {
			p.l.Println("Invalid URI, give only 1 id.")
			http.Error(rw, "Bad URI", http.StatusBadRequest)
			return
		}

		if len(g[0]) != 2 {
			p.l.Println("More than 2 capture groups.")
			http.Error(rw, "Bad URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]

		id, err := strconv.Atoi(idString)
		if err != nil {
			p.l.Println("Cannot convert to int.")
			http.Error(rw, "Bad URI", http.StatusBadRequest)
			return
		}

		p.l.Println("Got ID:", id)

		p.updateProduct(id, rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
	p.l.Println("Method not implemented.")
}

func (p *Products) getProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)

	p.l.Println("Handle GET products.")

	if err != nil {
		http.Error(rw, "Cannot get product list.", http.StatusInternalServerError)
		return
	}

	p.l.Println("Handled GET products successfully.")
}

func (p *Products) addProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST prodcuts.")

	pd := &data.Product{}

	err := pd.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Cannot create product.", http.StatusBadRequest)
		return
	}

	data.AddProduct(pd)
}

func (p Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT prodcuts.")

	pd := &data.Product{}

	err := pd.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Cannot create product.", http.StatusBadRequest)
		return
	}

	err = data.UpdateProduct(id, pd)
	if err == data.ErrProdNotFound {
		http.Error(rw, "Product not found.", http.StatusNotFound)
		return
	}

	if err != nil {
		http.Error(rw, "Product not found.", http.StatusInternalServerError)
		return
	}
}
