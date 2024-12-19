<template>
  <div class="container">
    <div class="login">
      <div class="lado-a">
        <div class="login-container">
          <h1 class="login-title">Acessar sua conta</h1>
          <form class="login-form" @submit.prevent="handleLogin">
            <label class="login-label" for="email">Email</label>
            <input
                class="login-input"
                id="email"
                type="email"
                v-model="email"
                placeholder="Seu email"
                required
            >

            <label class="login-label" for="password">Senha</label>
            <input
                class="login-input"
                id="password"
                type="password"
                v-model="password"
                placeholder="Sua senha"
                required
            >

            <button class="login-button" type="submit">
              Entrar
            </button>
          </form>
          <p class="register-text">Ainda não tem conta? <a href="#" class="register-link">Registre-se</a></p>
        </div>
      </div>
      <div class="lado-b">
        <div class="img-container">
          <img class="img" alt="Vue logo" src="../assets/login.png">
        </div>
      </div>
    </div>
    <footer class="footer">
      <p>&copy; 2024 All Rights Reserved - Maicon Gavino</p>
    </footer>
  </div>
</template>

<script>
export default {
  data() {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    handleLogin() {
      console.log('Email:', this.email, 'Senha:', this.password)
    }
  }
}
</script>

<style>
html, body {
  margin: 0;
  padding: 0;
  overflow-x: hidden; /* Impede rolagem horizontal */
  font-family: 'Helvetica Neue', Arial, sans-serif;
  background-color: #f6f6f6;
  color: #35495e;
  box-sizing: border-box;
}

*, *::before, *::after {
  box-sizing: inherit;
}

.container {
  display: flex;
  flex-direction: column;
  min-height: 100vh; /* Garante que o footer fique ao final */
}

.login {
  display: flex;
  flex: 1; /* O login ocupa o espaço restante da tela */
}

.lado-a,
.lado-b {
  flex: 1; /* Divide o espaço igualmente, evitando largura fixa */
  display: flex;
  align-items: center;
  justify-content: center;
  overflow: hidden;
  background-color: #ffffff;
  padding: 20px;
}

.login-container {
  background-color: #ffffff;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 40px;
  width: 100%;
  max-width: 380px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.login-title {
  margin: 0 0 20px;
  text-align: center;
  color: #35495e;
  font-size: 1.8em; /* Aumenta um pouco para melhor aproveitamento */
  font-weight: 600;
}

.login-form {
  display: flex;
  flex-direction: column;
}

.login-label {
  margin-bottom: 5px;
  font-weight: 500;
}

.login-input {
  border: 1px solid #ccc;
  border-radius: 4px;
  padding: 12px; /* Um pouco maior para mais conforto */
  font-size: 1em;
  margin-bottom: 15px;
  outline: none;
  transition: border-color 0.2s;
}

.login-input:focus {
  border-color: #42b983;
}

.login-button {
  background-color: #42b983;
  border: none;
  border-radius: 4px;
  padding: 14px; /* Botão mais alto */
  color: #ffffff;
  font-size: 1em;
  cursor: pointer;
  transition: background-color 0.2s;
  font-weight: 500;
}

.login-button:hover {
  background-color: #369c6d;
}

.register-text {
  margin-top: 20px;
  text-align: center;
  font-size: 1em; /* Levemente maior para melhor leitura */
}

.register-link {
  color: #42b983;
  text-decoration: none;
  font-weight: 500;
}

.register-link:hover {
  text-decoration: underline;
}

.img-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.img {
  background-color: #ffffff;
  border-radius: 8px;
  padding: 20px; /* Um padding menor para a imagem */
  height: auto;
  max-width: 80%; /* Reduz um pouco a imagem para não gerar scroll */
  object-fit: contain;
}

.footer {
  text-align: center; /* Centraliza o texto do rodapé */
  padding: 20px;
  font-size: 0.9em;
  color: #35495e;
}

/* Código destacado */
code {
  color: #42b983;
}
</style>
