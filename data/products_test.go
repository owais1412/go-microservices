package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Prod",
		Price: 1,
		SKU:   "ab-cd-ef",
	}
	err := p.Validate()

	if err != nil {
		t.Fatal(err)
	}
}
