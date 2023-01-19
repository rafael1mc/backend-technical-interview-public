package server

import (
	"errors"
	"log"
	"net"
	"net/http"
	"os"

	"play-service/controller"

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

	// Athletes
	rest.HandleFunc("/picks", s.c.HandleCreatePick).Methods("POST")
	rest.HandleFunc("/roster/{userID}", s.c.HandleGetRoster).Methods("GET")

	listener, err := net.Listen("tcp", ":"+s.port)
	if err != nil {
		return errors.New("failed to create listener" + err.Error())
	}

	log.Printf("TCP Server starting %d", listener.Addr().(*net.TCPAddr).Port)

	return http.Serve(listener, r)
}
