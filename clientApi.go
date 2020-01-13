package main

import (
	"fmt"
	"log"
	"net/http"
)

type server struct{}

func init() {
	fmt.Println("HttpServer starting up.")
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"message": "hello world"}`))
}

func apiHandle() {
	fmt.Println("apiHandle evoked.")
	s := &server{}
	http.Handle("/", s)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
