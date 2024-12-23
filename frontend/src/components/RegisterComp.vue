<script>
import axios from "axios";

export default {
  data() {
    return {
      form: {
        email: "",
        password: "",
      },
    };
  },
  methods: {
    async register() {
      try {
        const response = await axios.post("http://localhost:8080/register", this.form);
        if (response.status === 200) {
          alert("Usuário registrado com sucesso!");
          this.$router.push("/");
        }
      } catch (error) {
        console.error("Erro ao registrar usuário:", error);
        alert("Erro ao registrar. Tente novamente.");
      }
    },
  },
};
</script>

<template>
  <div class="register-page">
    <div class="card">
      <h2 class="register-title">Crie sua Conta</h2>
      <p class="register-subtitle">Preencha os campos abaixo para se registrar.</p>
      <form class="register-form" @submit.prevent="register">
        <div class="input-container">
          <label class="register-label" for="email">Nome</label>
          <input
              type="name"
              name="name"
              id="name"
              v-model="form.name"
              placeholder="Seu Nome..."
              required
              class="register-input"
          />
        </div>
        <div class="input-container">
          <label class="register-label" for="email">Email</label>
          <input
              type="email"
              name="email"
              id="email"
              v-model="form.email"
              placeholder="Seu e-mail..."
              required
              class="register-input"
          />
        </div>
        <div class="input-container">
          <label class="register-label" for="password">Senha</label>
          <input
              type="password"
              name="password"
              id="password"
              v-model="form.password"
              placeholder="Sua senha..."
              required
              class="register-input"
          />
        </div>
        <button class="register-button" type="submit">
          Registrar
        </button>
      </form>
      <div class="login-section">
        <p>Já tem conta?</p>
        <a href="/" class="login-button">Faça Login</a>
      </div>
    </div>
    <footer class="footer">
      <p>&copy; 2024 All Rights Reserved - Maicon Gavino</p>
    </footer>
  </div>
</template>

<style scoped>
/* Configurações gerais */
html,
body {
  height: 100%;
  margin: 0;
  padding: 0;
  width: 100%;
  font-family: "Helvetica Neue", Arial, sans-serif;
  background-color: #f6f6f6; /* Mesma cor de fundo da página de login */
  color: #35495e; /* Texto primário */
  box-sizing: border-box;
}

*,
*::before,
*::after {
  box-sizing: inherit;
}

/* Página de registro */
.register-page {
  display: flex;
  flex-direction: column;
  justify-content: center;
  align-items: center;
  height: 100vh;
  padding: 20px;
  background-color: #f6f6f6;
}

/* Card principal */
.card {
  background: #ffffff; /* Fundo branco */
  padding: 40px;
  border-radius: 8px; /* Bordas arredondadas */
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1); /* Sombra suave */
  text-align: center;
  max-width: 400px;
  width: 100%;
}

/* Título e subtítulo */
.register-title {
  font-size: 1.8rem;
  font-weight: bold;
  color: #35495e; /* Mesma cor do texto principal */
  margin-bottom: 10px;
}

.register-subtitle {
  font-size: 1rem;
  color: #6c757d; /* Cor do subtítulo */
  margin-bottom: 20px;
}

/* Formulário */
.register-form {
  display: flex;
  flex-direction: column;
  gap: 15px;
}

.register-label {
  font-size: 1rem;
  color: #495057; /* Cor secundária */
  margin-bottom: 5px;
  text-align: left;
}

.register-input {
  width: 100%;
  padding: 12px;
  font-size: 1rem;
  border: 1px solid #ccc;
  border-radius: 4px;
  transition: border-color 0.3s, box-shadow 0.3s;
}

.register-input:focus {
  border-color: #42b983; /* Cor do destaque no foco */
  box-shadow: 0 0 5px rgba(66, 185, 131, 0.5); /* Sombra de foco */
  outline: none;
}

/* Botão de registro */
.register-button {
  background-color: #42b983; /* Cor do botão principal */
  color: #fff;
  padding: 12px;
  font-size: 1rem;
  font-weight: bold;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  transition: background-color 0.3s, transform 0.2s;
}

.register-button:hover {
  background-color: #369c6d; /* Cor mais escura no hover */
  transform: scale(1.05); /* Leve crescimento no hover */
}

/* Link de login */
.login-section {
  margin-top: 20px;
}

.login-button {
  display: inline-block;
  font-size: 1rem;
  color: #ffffff; /* Verde característico */
  text-decoration: none;
  font-weight: bold;
  padding: 8px 16px;
  border: 2px solid #42b983; /* Borda verde */
  border-radius: 4px;
  transition: color 0.3s, background-color 0.3s, transform 0.2s;
}

.login-button:hover {
  background-color: #369c6d; /* Fundo verde ao passar o mouse */
  color: #fff; /* Texto branco */
  transform: scale(1.05);
}

/* Rodapé */
.footer {
  position: fixed; /* Fixar no rodapé */
  bottom: 0;
  width: 100%;
  text-align: center;
  padding: 15px 20px;
  font-size: 0.9rem;
  color: #6c757d;
  background-color: #ffffff;
  box-shadow: 0 -2px 5px rgba(0, 0, 0, 0.1); /* Sombra sutil no topo do rodapé */
}
</style>