package router

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Config struct {
	Host string
	Port string
}

type Router struct {
	config  Config
	Router  *mux.Router
	Handler Handler
}

func NewRouter(config Config, h *Handler) *Router {
	r := mux.NewRouter()
	r.HandleFunc("/books", h.Get).Methods(http.MethodGet)
	r.HandleFunc("/books", h.Add).Methods(http.MethodPost)
	// r.HandleFunc("/books/{id}", deleteBook).Methods(http.MethodDelete)
	// r.HandleFunc("/books/{id}/borrow", borrowBook).Methods(http.MethodPut)
	// r.HandleFunc("/books/{id}/return", returnBook).Methods(http.MethodPut)
	// r.HandleFunc("/random-quote", getRandomQuote).Methods(http.MethodGet)
	return &Router{
		config: config,
		Router: r,
	}
}

func (r *Router) Run() error {
	return http.ListenAndServe(fmt.Sprintf("%s:%s", r.config.Host, r.config.Port), r.Router)
}
