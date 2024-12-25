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
	Name     string `json:"name"`
	Phone    string `json:"phone"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RegisterResponse struct {
	Message string `json:"message"`
}

// LoginHandler realiza a autenticação do usuário
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var loginReq LoginRequest
	if err := json.NewDecoder(r.Body).Decode(&loginReq); err != nil {
		log.Printf("Erro ao decodificar JSON: %v", err)
		http.Error(w, "Erro ao processar dados", http.StatusBadRequest)
		return
	}
	log.Printf("Tentativa de login com o email: %s", loginReq.Email)

	var user models.User
	err := config.DB.QueryRow("SELECT id, email, password FROM users WHERE email = $1", loginReq.Email).Scan(&user.ID, &user.Email, &user.Password)
	if err == sql.ErrNoRows {
		log.Println("Usuário não encontrado")
		http.Error(w, "Email ou senha incorretos", http.StatusUnauthorized)
		return
	} else if err != nil {
		log.Printf("Erro ao buscar usuário no banco: %v", err)
		http.Error(w, "Erro interno", http.StatusInternalServerError)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginReq.Password))
	if err != nil {
		log.Println("Senha incorreta")
		http.Error(w, "Email ou senha incorretos", http.StatusUnauthorized)
		return
	}

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
	query := "INSERT INTO users (name, phone, email, password) VALUES ($1, $2, $3, $4)"
	_, err = config.DB.Exec(query, registerReq.Name, registerReq.Phone, registerReq.Email, string(hashedPassword))
	if err != nil {
		http.Error(w, "Erro ao salvar usuário", http.StatusInternalServerError)
		log.Printf("Erro ao inserir usuário: %v", err)
		return
	}
	//log.Printf("email e senha", registerReq.Email, registerReq.Password)
	// Responde com sucesso
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(RegisterResponse{Message: "Usuário registrado com sucesso"})
}
