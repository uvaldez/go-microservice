package handlers

import (
	"net/http"

	"github.com/uvaldez/go-microservice/data"
)

func (p *Products) PostProduct(rw http.ResponseWriter, r *http.Request) {
	prod := r.Context().Value(KeyProduct{}).(data.Product)
	data.AddProduct(&prod)
}
