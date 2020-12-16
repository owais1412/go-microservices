package handlers

import (
	"log"
	"net/http"
)

// Goodbye struct
type Goodbye struct {
	l *log.Logger
}

// NewGoodbye does dependency injection on Goodbye object
// injects log
func NewGoodbye(l *log.Logger) *Goodbye {
	return &Goodbye{l}
}

func (g *Goodbye) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	g.l.Println("Goodbye")
	rw.Write([]byte("Bye"))
}
