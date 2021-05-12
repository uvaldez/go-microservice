package data

import (
	"encoding/json"
	"fmt"
	"io"
	"time"

	"github.com/go-playground/validator"
)

// Product defines the structure for an API product
// swagger:model
type Product struct {
	// the id for the product
	//
	// required: false
	// min: 1
	ID int `json:"id"` // Unique identifier for the product
	// the name for this poduct
	//
	// required: true
	// max length: 255
	Name string `json:"Name" validate:"required"`
	// the description for this poduct
	//
	// required: false
	// max length: 10000
	Description string `json:"Description"`
	// the price for the product
	//
	// required: true
	// min: 0.01
	Price     float32 `json:"Price" validate:"gt=0"`
	createdOn string  `json:"-"`
	updatedOn string  `json:"-"`
}

type Products []*Product

var ErrProdNotFound = fmt.Errorf("Product not found")

func UpdateProduct(id int, p *Product) error {
	_, pos, err := findProduct(id)
	if err != nil {
		return err
	}
	p.ID = id
	ProductList[pos] = p
	return nil
}

func AddProduct(p *Product) {
	p.ID = getNextID()
	ProductList = append(ProductList, p)
}

func GetProducts() Products {
	return ProductList
}

func DeleteProduct(id int) error {
	i := findIndexByProductID(id)

	if i == -1 {
		return ErrProdNotFound
	}

	ProductList = append(ProductList[:i], ProductList[i+1])
	return nil
}

// findIndex finds the index of a product in the database
// returns -1 when no product can be found
func findIndexByProductID(id int) int {
	for i, p := range ProductList {
		if p.ID == id {
			return i
		}
	}
	return -1
}

func (p *Product) FromJson(r io.Reader) error {
	d := json.NewDecoder(r)
	return d.Decode(p)
}

func (p *Products) ToJson(w io.Writer) error {
	e := json.NewEncoder(w)
	return e.Encode(p)
}

func (p *Product) Validate() error {
	validate := validator.New()
	return validate.Struct(p)
}

func getNextID() int {
	lp := ProductList[len(ProductList)-1]
	return lp.ID + 1
}

func findProduct(id int) (*Product, int, error) {
	for i, p := range ProductList {
		if p.ID == id {
			return p, i, nil
		}
	}
	return nil, -1, ErrProdNotFound
}

var ProductList = []*Product{
	{
		ID:          1,
		Name:        "Late",
		Description: "Frothy milky coffe",
		Price:       2.45,
		createdOn:   time.Now().UTC().String(),
		updatedOn:   time.Now().UTC().String(),
	},
	{
		ID:          2,
		Name:        "Espresso",
		Description: "Short and strong coffe without milk",
		Price:       1.45,
		createdOn:   time.Now().UTC().String(),
		updatedOn:   time.Now().UTC().String(),
	},
}
