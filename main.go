package main

import (
	"log"
	"net/http"

	"github.com/ninjadiego/golang-jwt/server"
)

func main() {
	s := server.New()
	log.Fatal(http.ListenAndServe(":8080", s.Router()))
}
