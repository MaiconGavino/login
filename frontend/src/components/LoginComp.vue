<template>
  <div class="login-page">
    <div class="left-panel">
      <div class="logo-container">
        <img class="logo" src="../assets/logo.png" alt="Logo">
        <h1 class="app-name">AgroScan</h1>
        <p class="tagline">Identificação, classificação e manejo de pragas</p>
        <p class="description">
          Sistema de visão computacional aplicado à identificação e recomendação de manejo de pragas para lavoura de cacau.
        </p>
        <p class="version">Version 1.1</p>
      </div>
    </div>
    <div class="right-panel">
      <div class="login-container">
        <h2 class="login-title">Login AgroScan</h2>
        <p class="login-subtitle">Acesse sua conta.</p>
        <form class="login-form" @submit.prevent="login">
          <label class="login-label" for="email">Email</label>
          <div class="input-container">
            <input
                class="login-input"
                id="email"
                type="email"
                v-model="form.email"
                placeholder="Seu e-mail..."
                aria-label="Email"
                required
            />
          </div>

          <label class="login-label" for="password">Senha</label>
          <div class="input-container">
            <input
                class="login-input"
                id="password"
                type="password"
                v-model="form.password"
                placeholder="Sua senha..."
                aria-label="Senha"
                required
            />
          </div>

          <button class="login-button" type="submit" :disabled="loading">
            <span v-if="loading">Carregando...</span>
            <span v-else>Entrar</span>
          </button>
        </form>
        <div class="signup-section">
          <p>Ainda não tem conta?</p>
          <a href="/register" class="signup-button" @click="register">Registre-se</a>
        </div>
      </div>
    </div>
    <footer class="footer">
      <p>&copy; 2024 All Rights Reserved - Maicon Gavino</p>
    </footer>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      form: {
        email: "",
        password: "",
      },
      loading: false,
    };
  },

  methods: {
    async login() {
      this.loading = true;
      try {
        const response = await axios.post("http://localhost:8080/", {
          email: this.form.email,
          password: this.form.password,
        });

        if (response.status === 200) {
          alert("Login realizado com sucesso!");
          this.$router.push("/");
        }
      } catch (error) {
        console.error("Erro ao realizar login:", error.response || error.message);
        alert("Erro ao realizar login. Verifique suas credenciais.");
      } finally {
        this.loading = false;
      }
    },

    register() {
      this.$router.push("/register");
    },
  },
};
</script>

<style>
/* Configuração global */
html,
body {
  margin: 0;
  padding: 0;
  width: 100%;
  height: 100%;
  overflow: hidden;
  font-family: "Helvetica Neue", Arial, sans-serif;
  background-color: #f6f6f6;
  color: #35495e;
  box-sizing: border-box;
}

*,
*::before,
*::after {
  box-sizing: inherit;
}

/* Layout principal */
.login-page {
  display: flex;
  flex-direction: row;
  width: 100%;
  height: 100vh;
}

/* Divisão dos painéis */
.left-panel,
.right-panel {
  flex: 1;
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
}

.left-panel {
  background-color: #35495e;
  color: #ffffff;
  padding: 40px;
  text-align: center;
}

.right-panel {
  background-color: #ffffff;
  padding: 40px;
}

.logo {
  width: 80px;
  margin-bottom: 20px;
}

.app-name {
  font-size: 2em;
  font-weight: bold;
  margin: 10px 0;
}

.tagline {
  font-size: 1.2em;
  margin-bottom: 10px;
}

.description {
  font-size: 0.9em;
  margin-bottom: 30px;
}

.version {
  font-size: 0.8em;
  color: #b0c4de;
}

/* Login container */
.login-container {
  max-width: 400px;
  width: 100%;
}

.login-title {
  font-size: 1.8em;
  font-weight: bold;
  margin-bottom: 10px;
  color: #35495e;
}

.login-subtitle {
  font-size: 1em;
  margin-bottom: 20px;
  color: #6c757d;
}

.login-form {
  display: flex;
  flex-direction: column;
}

.login-label {
  font-size: 0.9em;
  margin-bottom: 5px;
  color: #495057;
}

.input-container {
  position: relative;
  margin-bottom: 20px;
}

.login-input {
  width: 100%;
  padding: 12px;
  font-size: 1em;
  border: 1px solid #ccc;
  border-radius: 4px;
  outline: none;
}

.login-input:focus {
  border-color: #42b983;
}

.login-button {
  background-color: #42b983;
  color: #ffffff;
  border: none;
  border-radius: 4px;
  padding: 12px;
  font-size: 1em;
  cursor: pointer;
  font-weight: bold;
  transition: background-color 0.3s;
}

.login-button:hover {
  background-color: #369c6d;
}

.signup-section {
  text-align: center;
  margin-top: 20px;
}

.signup-button {
  background-color: transparent;
  color: #42b983;
  font-size: 0.9em;
  border: 1px solid #42b983;
  border-radius: 4px;
  padding: 8px 12px;
  text-decoration: none;
  cursor: pointer;
  transition: all 0.3s;
}

.signup-button:hover {
  background-color: #42b983;
  color: #ffffff;
}

.footer {
  text-align: center;
  padding: 10px 20px;
  font-size: 0.8em;
  color: #6c757d;
  background-color: #ffffff;
  flex-shrink: 0;
}
</style>
