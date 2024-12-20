package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"golang.org/x/crypto/bcrypt"
	"login/config"
	"login/models"
)

type LoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Message string `json:"message"`
}

type RegisterRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

// LoginHandler realiza a autenticação do usuário
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodifica os dados do corpo da requisição
	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	// Busca o usuário no banco de dados
	var user models.User
	err := config.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", loginReq.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		http.Error(w, "Email ou senha incorretos", http.StatusUnauthorized)
		return
	} else if err != nil {
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		log.Printf("Erro ao buscar usuário: %v", err)
		return
	}

	// Verifica a senha
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		http.Error(w, "Email ou senha incorretos", http.StatusUnauthorized)
		return
	}

	// Responde com sucesso
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(LoginResponse{Message: "Login realizado com sucesso"})
}

// RegisterHandler registra um novo usuário
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodifica os dados do corpo da requisição
	var registerReq RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&registerReq); err != nil {
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}

	// Gera o hash da senha
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(registerReq.Password), bcrypt.DefaultCost)
	if err != nil {
		http.Error(w, "Erro ao processar senha", http.StatusInternalServerError)
		log.Printf("Erro ao gerar hash da senha: %v", err)
		return
	}

	// Insere o novo usuário no banco de dados
	query := "INSERT INTO users (email, password) VALUES ($1, $2)"
	_, err = config.DB.Exec(query, registerReq.Email, string(hashedPassword))
	if err != nil {
		http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
		log.Printf("Erro ao inserir usuário: %v", err)
		return
	}

	// Responde com sucesso
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RegisterResponse{Message: "Usuário registrado com sucesso"})
}
