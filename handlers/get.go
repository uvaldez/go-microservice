package handlers

import (
	"net/http"

	"github.com/uvaldez/go-microservice/data"
)

// swagger:route GET /products products listProducts
// Return a list of products from the database
//
// responses:
//	200: productsResponse

// GetProducts returns the products from the data store
func (p *Products) GetProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJson(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusInternalServerError)
	}
}
