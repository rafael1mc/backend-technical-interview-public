package server

import (
	"entity-service/controller"
	"errors"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

type Server struct {
	port string

	c *controller.Controller
}

func New(c *controller.Controller) *Server {
	p := os.Getenv("PORT")

	return &Server{p, c}
}

func (s *Server) Run() error {
	r := mux.NewRouter()

	rest := r.NewRoute().Subrouter()

	// Entities
	rest.HandleFunc("/{entityType}/{entityID}", s.c.GetEntityByIDHandler).Methods("GET")
	rest.HandleFunc("/{entityType}", s.c.ListEntityByIDHandler).Methods("POST")

	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return errors.New("failed to create listener" + err.Error())
	}

	log.Printf("TCP Server starting %d", listener.Addr().(*net.TCPAddr).Port)

	return http.Serve(listener, r)
}
