<template>
  <div id="app">
    <h1>Inscription</h1>
    <form @submit.prevent="register">
      <input v-model="form.email" type="email" placeholder="Email" required />
      <input v-model="form.username" type="text" placeholder="Nom d'utilisateur" required />
      <input v-model="form.password" type="password" placeholder="Mot de passe" required />
      <input v-model="form.first_name" type="text" placeholder="Prénom" required />
      <input v-model="form.last_name" type="text" placeholder="Nom" required />
      <input v-model="form.date_of_birth" type="date" placeholder="Date de naissance" required />
      <input v-model="form.avatar" type="text" placeholder="URL de l'avatar (optionnel)" />
      <input v-model="form.nickname" type="text" placeholder="Surnom (optionnel)" />
      <textarea v-model="form.about_me" placeholder="À propos de moi (optionnel)"></textarea>
      <button type="submit">S'inscrire</button>
    </form>

    <h1>Connexion</h1>
    <form @submit.prevent="login">
      <input v-model="loginForm.email_or_username" type="text" placeholder="Email ou nom d'utilisateur" required />
      <input v-model="loginForm.password" type="password" placeholder="Mot de passe" required />
      <button type="submit">Se connecter</button>
    </form>

    <button @click="testProtected">Tester accès protégé</button>

    <div v-if="message">{{ message }}</div>
    <div v-if="protectedMessage">{{ protectedMessage }}</div>
  </div>
</template>

<script>
export default {
  name: 'App',
  data() {
    return {
      form: {
        email: '',
        username: '',
        password: '',
        first_name: '',
        last_name: '',
        date_of_birth: '',
        avatar: '',
        nickname: '',
        about_me: ''
      },
      loginForm: {
        email_or_username: '',
        password: ''
      },
      message: '',
      protectedMessage: ''
    }
  },
  methods: {
    async register() {
      try {
        const response = await fetch('http://localhost:8081/api/register', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify(this.form)
        });
        if (response.ok) {
          this.message = "Inscription réussie !";
        } else {
          const data = await response.json();
          this.message = data.error || "Erreur lors de l'inscription.";
        }
      } catch (err) {
        this.message = "Erreur réseau.";
      }
    },
    async login() {
      const response = await fetch('http://localhost:8081/api/login', {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify({
          email_or_username: this.loginForm.email_or_username,
          password: this.loginForm.password
        })
      });
      if (response.ok) {
        const data = await response.json();
        localStorage.setItem('token', data.token);
        this.message = "Connexion réussie !";
      } else {
        this.message = "Identifiants invalides";
      }
    },
    async testProtected() {
      console.log("Testing protected route...");
      const token = localStorage.getItem('token');
      const response = await fetch('http://localhost:8081/api/protected', {
        headers: { 'Authorization': 'Bearer ' + token }
      });
      if (response.ok) {
        this.protectedMessage = await response.text();
      } else {
        this.protectedMessage = "Accès refusé (token manquant ou invalide)";
      }
    }
  }
}
</script>

<style>
#app {
  max-width: 400px;
  margin: 40px auto;
  padding: 20px;
  border: 1px solid #eee;
  border-radius: 8px;
  background: #fafafa;
}
input, textarea {
  display: block;
  width: 100%;
  margin: 10px 0;
  padding: 8px;
}
button {
  padding: 10px 20px;
  background: #42b983;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>
