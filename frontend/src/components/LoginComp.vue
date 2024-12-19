<template>
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
        <img class="img" alt="Vue logo" src="../assets/sign-in.png">
      </div>
    </div>
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

<style scoped>
body {
  margin: 0;
  font-family: 'Helvetica Neue', Arial, sans-serif;
  background-color: #f6f6f6;
  color: #35495e;
}

.login {
  height: 100vh;
  display: flex;
}

.lado-a,
.lado-b {
  overflow: auto;
  width: 50%;
  height: 100%;
  box-sizing: border-box;
  padding: 0 20px;
}

.lado-a {
  border-right: 1px solid #ccc;
  background-color: #f6f6f6;
  display: flex;
  align-items: center;
  justify-content: center;
}

.lado-b {
  background-color: #ffffff;

  /* Centraliza o conteúdo horizontal e verticalmente */
  display: flex;
  justify-content: center;
  align-items: center;
}

/* Contêiner para a imagem e o texto */
.img-container {
  display: flex;
  flex-direction: column;
  align-items: center;
}

.login-container {
  background-color: #ffffff;
  border: 1px solid #ccc;
  border-radius: 8px;
  padding: 40px;
  width: 100%;
  max-width: 320px;
  box-shadow: 0 2px 8px rgba(0,0,0,0.1);
}

.login-title {
  margin: 0 0 20px;
  text-align: center;
  color: #35495e;
  font-size: 1.5em;
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
  padding: 10px;
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
  padding: 12px;
  color: #ffffff;
  font-size: 1em;
  cursor: pointer;
  transition: background-color 0.2s;
}

.login-button:hover {
  background-color: #369c6d;
}

.register-text {
  margin-top: 20px;
  text-align: center;
  font-size: 0.9em;
}

.register-link {
  color: #42b983;
  text-decoration: none;
  font-weight: 500;
}

.register-link:hover {
  text-decoration: underline;
}

.img {
  background-color: #ffffff;
  border-radius: 8px;
  padding: 40px;
  width: 90%;
  max-width: 100%;
}


/* Código destacado */
code {
  color: #42b983;
}
</style>
