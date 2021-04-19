package handlers

import (
	"net/http"

	"github.com/progmatic-99/microService/data"
)

// swagger:route GET / products listProducts
// Returns a list of products
// responses:
//  200: productsResponse
func (p *Products) GetProduct(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()

	// Add header to the response so that test passes
	rw.Header().Add("Content-Type", "application/json")

	err := lp.ToJson(rw)

	p.l.Println("Handle GET products.")

	if err != nil {
		http.Error(rw, "Cannot get product list.", http.StatusInternalServerError)
		return
	}
}
