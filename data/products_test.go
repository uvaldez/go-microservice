package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "test prod",
		Price: 1.50,
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
