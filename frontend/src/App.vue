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
    <div v-if="message">{{ message }}</div>
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
      message: ''
    }
  },
  methods: {
    async register() {
      try {
        const response = await fetch('/api/register', {
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
