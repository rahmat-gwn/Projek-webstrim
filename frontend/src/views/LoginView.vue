<template>
  <div class="login-container">
    <div class="login-card">
      <h2>Login</h2>
      <form @submit.prevent="login">
        <div class="form-group">
          <label for="username">Username</label>
          <input
            type="text"
            id="username"
            v-model="username"
            placeholder="Enter your username"
            required
          />
        </div>
        <div class="form-group">
          <label for="password">Password</label>
          <input
            type="password"
            id="password"
            v-model="password"
            placeholder="Enter your password"
            required
          />
        </div>
        <button type="submit">Login</button>
      </form>
      <div class="footer">
        <span>Don't have an account?</span>
        <router-link to="/register">Register</router-link>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue';

const username = ref('');
const password = ref('');

// Function untuk menangani login
const login = () => {
  if (username.value === 'admin' && password.value === 'password') {
    alert('Login Successful');
    // Redirect ke halaman lain setelah login berhasil
    window.location.href = '/home';
  } else {
    alert('Invalid username or password');
  }
};
</script>

<style scoped>
/* Background Fullscreen */
body {
  background: linear-gradient(to right, #ff7e5f, #feb47b);
  font-family: 'Arial', sans-serif;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 0;
}

/* Card Styling */
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}

.login-card {
  background: #fff;
  padding: 40px 40px;
  border-radius: 12px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
  text-align: center;
}

h2 {
  margin-bottom: 20px;
  font-size: 24px;
  color: #333;
}

/* Form Group Styling */
.form-group {
  margin-bottom: 20px;
}

label {
  display: block;
  font-weight: bold;
  margin-bottom: 8px;
  color: #555;
}

input {
  width: 100%;
  padding: 12px;
  border: 1px solid #ddd;
  border-radius: 8px;
  box-sizing: border-box;
  font-size: 16px;
  margin-bottom: 10px;
}

input:focus {
  border-color: #ff7e5f;
  outline: none;
}

/* Button Styling */
button {
  width: 100%;
  padding: 12px;
  background-color: #ff7e5f;
  color: white;
  font-size: 16px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #feb47b;
}

/* Footer Section */
.footer {
  margin-top: 20px;
  font-size: 14px;
  color: #777;
}

footer a {
  text-decoration: none;
  color: #ff7e5f;
  font-weight: bold;
}

footer a:hover {
  text-decoration: underline;
}
</style>
