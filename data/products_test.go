package data

import "testing"

func TestChecksValidation(t *testing.T) {
	p := &Product{
		Name:  "Prod",
		Price: 1,
		SKU:   "ab-cd-ef",
	}

	v := NewValidation()
	err := v.Validate(p)

	if err != nil {
		t.Fatal(err)
	}
}
