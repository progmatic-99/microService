package data

import "testing"

func TestValidateProduct(t *testing.T) {
	p := &Product{
		Name:  "Tea",
		Price: 2.00,
		SKU:   "abc",
	}

	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
