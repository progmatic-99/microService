package handlers

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/progmatic-99/microService/data"
)

// swagger:route DELETE /{id} products deleteProduct
// Returns nothing.
// responses:
//  201: noContent
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
