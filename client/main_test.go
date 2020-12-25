package main

import (
	"fmt"
	"testing"

	"github.com/owais1412/simpleServer/client/client"
	"github.com/owais1412/simpleServer/client/client/products"
)

func TestClient(t *testing.T) {
	cfg := client.DefaultTransportConfig().WithHost("localhost:9090")
	c := client.NewHTTPClientWithConfig(nil, cfg)
	params := products.NewListProductsParams()
	prod, err := c.Products.ListProducts(params)

	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n", prod.GetPayload()[0])
	// t.Fail()
}
