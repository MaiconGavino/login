package routes

import (
	"login/handlers"
	"net/http"
)

func RegisterRoutes() {
	http.HandleFunc("/", handlers.LoginHandler)
	http.HandleFunc("/register", handlers.RegisterHandler)
}
