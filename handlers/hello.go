package handlers

import (
	"log"
	"net/http"
)

type HelloHandler struct{}

func (h *HelloHandler) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	log.Printf("Hello http base handler")
}
func NewHelloHandler() *HelloHandler {
	return &HelloHandler{}
}
