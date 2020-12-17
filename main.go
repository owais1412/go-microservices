package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/owais1412/simpleServer/handlers"
)

const (
	port = "9090"
)

func main() {

	loggerProductAPI := log.New(os.Stdout, "product-api", log.LstdFlags)
	ph := handlers.NewProducts(loggerProductAPI)

	sm := http.NewServeMux()
	sm.Handle("/", ph)
	// Custom server
	s := &http.Server{
		Addr:         fmt.Sprintf(":%s", port),
		Handler:      sm,
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
