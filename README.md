# Tela de Login e de Cadastro usando Vue3 e Golang.

### Descritivo:
Este projeto é uma aplicação web desenvolvida utilizando **Vue.js** no frontend e **Golang** no backend, ele implementa funcionalidades básicas de registro e login com integração com o banco de dados **PostgreSQL**.
## 1. Estrutura do projeto:
```bash
|-- config
|--- db.go            # Configuração de conexão com o banco de dados
|-- handlers
|--- userHandler.go   # Handlers para registro e login
|-- models
|--- user.go          # Modelo de dados do usuário
|-- routes
|--- routes.go        # Rotas definidas para as APIs
|-- main.go           # Arquivo principal do backend
|-- frontend
|--- src
|---- components
|----- RegisterComp.vue  # Componente de registro
|---- views
|----- LoginView.vue     # Tela de login
|----- RegisterView.vue  # Tela de registro
|---- router
|----- index.js          # Rotas do frontend
|---- App.vue            # Componente principal do Vue
|---- main.js            # Ponto de entrada do frontend
```
## 2. Tecnologias Utilizadas
### Frontend
- **Framework:** Vue.js 3
- **HTTP Client:** Axios
- **Estilização:** CSS

### Backend
- **Linguagem:** Golang
- **Criptografia:** Bcrypt
- **Banco de Dados:** PostgreSQL

## 3. Configuração do Ambiente
### Condiguração do Banco de Dados
Execute os seguintes comandos SQL para criar o banco de dados e a tabela:
```bash
CREATE DATABASE login_project;

\c login_project

CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    phone VARCHAR(15) NOT NULL,
    email VARCHAR(255) UNIQUE NOT NULL,
    password TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT NOW()
);
```
### 4. Variáveis de Ambiente
Crie uma arquivo **.env** na raiz do backend e adicione as seguintes configurações:
```bash
DB_USER=seu_usuario
DB_PASSWORD=sua_senha
DB_NAME=login_project
DB_HOST=localhost
DB_PORT=5432
DB_SSLMODE=disable
```
### 5. Execução do Projeto
### Backend
**1. Instale as dependências:**
```bash
go mod tidy
```
###
**2. Inicie o servidor:**
```bash
go run main.go
```
###
**3. O servidor estará disponível em:**
```bash
http://localhost:8080
```
###

### Frontend
**1. Acesse a pasta:**
```bash
cd frontend
```
###
**2. Instalar as dependências:**
```bash
npm install
```
###
**3. Inicie o servidor:**
```bash
npm run serve
```
###
**4. O servidor estará disponível em:**
```bash
http://localhost:8081
```
###