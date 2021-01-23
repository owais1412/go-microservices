package handlers

import (
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/owais1412/go-microservices/data"
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

	if r.Method == http.MethodPost {
		p.addProduct(rw, r)
		return
	}

	if r.Method == http.MethodPut {
		// get id from URL
		re := regexp.MustCompile("/([0-9]+)")
		g := re.FindAllStringSubmatch(r.URL.Path, -1)
		p.l.Printf("Group: %#v", g)
		if (len(g) != 1) || (len(g[0]) != 2) {
			http.Error(rw, "Invalid URI", http.StatusBadRequest)
			return
		}

		idString := g[0][1]
		id, err := strconv.Atoi(idString)
		if err != nil {
			http.Error(rw, "Failed converting ID to int", http.StatusBadRequest)
			return
		}
		p.updateProduct(id, rw, r)
		return
	}

	// catch all
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle GET Products")

	// fetch the products from the datastore
	lp := data.GetProducts()
	err := lp.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to marshal json", http.StatusBadRequest)
	}
}

func (p *Products) addProduct(rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle POST Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
	}

	data.AddProduct(prod)
}

func (p *Products) updateProduct(id int, rw http.ResponseWriter, r *http.Request) {
	p.l.Println("Handle PUT Product")

	prod := &data.Product{}
	err := prod.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "Unable to unmarshal json", http.StatusBadRequest)
		return
	}

	er := data.UpdateProduct(id, prod)
	if er == data.ErrProductNotFound {
		http.Error(rw, data.ErrProductNotFound.Error(), http.StatusNotFound)
		return
	}
	if er != nil {
		http.Error(rw, data.ErrProductNotFound.Error(), http.StatusInternalServerError)
		return
	}
	p.l.Printf("New Prod: %#v", prod)
}
