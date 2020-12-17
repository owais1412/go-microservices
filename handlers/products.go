package handlers

import (
	"log"
	"net/http"

	"github.com/owais1412/simpleServer/data"
)

// Products struct
type Products struct {
	l *log.Logger
}

// NewProducts does dependency injection on Products object
// injects log
func NewProducts(l *log.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}
}
