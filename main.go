package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type Server struct{}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	environment := "production"
	if en := os.Getenv("DEVOPS_ENV"); en != "" {
		environment = en
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(fmt.Sprintf(`{"message": "hello from %s"}`, environment)))
}

func main() {
	s := &Server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
