package main

import (
	"log"
	"login/config"
	"login/routes"
	"net/http"
)

// Middleware para configurar CORS
func enableCORS(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Responde diretamente se for uma requisição OPTIONS (Preflight Request)
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	// Conectar ao banco de dados
	config.ConnectDB()

	// Configurar rotas
	routes.RegisterRoutes()

	// Adicionar middleware CORS
	handler := enableCORS(http.DefaultServeMux)

	// Iniciar o servidor
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatalf("Error starting server: %s", err)
	}
	log.Println("Server started on port 8080")
}
