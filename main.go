package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/go-openapi/runtime/middleware"

	gohandlers "github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/owais1412/simpleServer/data"
	"github.com/owais1412/simpleServer/handlers"
)

const (
	port = "9090"
)

func main() {

	loggerProductAPI := log.New(os.Stdout, "product-api", log.LstdFlags)
	validation := data.NewValidation()

	// create the handlers
	ph := handlers.NewProducts(loggerProductAPI, validation)

	// create a new serve mux and register the handlers
	sm := mux.NewRouter()

	// handlers for API
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/products", ph.ListAll)
	getRouter.HandleFunc("/products/{id:[0-9]+}", ph.ListSingle)

	putRouter := sm.Methods(http.MethodPut).Subrouter()
	putRouter.HandleFunc("/products", ph.Update)
	putRouter.Use(ph.MiddlewareValidateProduct)

	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/products", ph.Create)
	postRouter.Use(ph.MiddlewareValidateProduct)

	deleteRouter := sm.Methods(http.MethodDelete).Subrouter()
	deleteRouter.HandleFunc("/products/{id:[0-9]+}/", ph.Delete)

	// handler for documnetation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)
	getRouter.Handle("/docs", sh)
	getRouter.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

	// CORS handler
	ch := gohandlers.CORS(gohandlers.AllowedOrigins([]string{"*"}))

	// Custom server
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      ch(sm),
		IdleTimeout:  120 * time.Second,
		ReadTimeout:  1 * time.Second,
		WriteTimeout: 1 * time.Second,
	}

	go func() {
		loggerProductAPI.Printf("Server running at localhost:%v\n", port)
		err := s.ListenAndServe()
		if err != nil {
			loggerProductAPI.Fatal(err)
		}
	}()

	sigChan := make(chan os.Signal)
	signal.Notify(sigChan, os.Interrupt)
	signal.Notify(sigChan, os.Kill)

	sig := <-sigChan
	loggerProductAPI.Println("Recieved terminate, graceful shutdown", sig)

	tc, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(tc)
}
