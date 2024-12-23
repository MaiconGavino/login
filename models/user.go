package models

type User struct {
	ID       int    `json:"id"`       // Identificador único do usuário
	Email    string `json:"email"`    // Email do usuário
	Password string `json:"password"` // Senha criptografada do usuário
}
