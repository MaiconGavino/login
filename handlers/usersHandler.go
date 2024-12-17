package handlers

import (
	"log"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("login")
}
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	log.Println("register")
}
