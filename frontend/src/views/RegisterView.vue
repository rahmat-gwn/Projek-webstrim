<template>
  <div class="register-container">
    <div class="register-card">
      <h1>Create an Account</h1>
      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label for="username">Username</label>
          <input
            id="username"
            v-model="username"
            type="text"
            placeholder="Enter your username"
            required
          />
        </div>

        <div class="form-group">
          <label for="email">Email</label>
          <input
            id="email"
            v-model="email"
            type="email"
            placeholder="Enter your email"
            required
          />
        </div>

        <div class="form-group">
          <label for="phone">Phone Number</label>
          <input
            id="phone"
            v-model="phone"
            type="tel"
            placeholder="Enter your phone number"
            required
          />
        </div>

        <div class="form-group">
          <label for="password">Password</label>
          <input
            id="password"
            v-model="password"
            type="password"
            placeholder="Enter your password"
            required
          />
        </div>

        <button type="submit">Register</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, reactive } from "vue";
import axios from "axios";

export default defineComponent({
  name: "RegisterView",
  setup() {
    const form = reactive({
      username: "",
      email: "",
      phone: "",
      password: ""
    });

    const handleRegister = async () => {
      try {
        await axios.post(`${import.meta.env.VITE_API_BASE_URL}/register`, {
          username: form.username,
          email: form.email,
          phone: form.phone,
          password: form.password
        });
        alert("Registration successful!");
        // Redirect to login or home page
        window.location.href = "/login";
      } catch (err) {
        alert("Registration failed!");
      }
    };

    return {
      ...form,
      handleRegister
    };
  }
});
</script>

<style scoped>
/* Fullscreen Background */
body {
  background: linear-gradient(to right, #4facfe, #00f2fe);
  font-family: 'Arial', sans-serif;
  margin: 0;
  height: 100vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

/* Centered Registration Card */
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  width: 100%;
  height: 100%;
}

.register-card {
  background: white;
  padding: 40px;
  border-radius: 12px;
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
  width: 100%;
  max-width: 400px;
  text-align: center;
}

h1 {
  margin-bottom: 20px;
  font-size: 24px;
  color: #333;
}

/* Styling for Input Fields */
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
  font-size: 16px;
  margin-bottom: 10px;
}

input:focus {
  border-color: #4facfe;
  outline: none;
}

/* Button Styling */
button {
  width: 100%;
  padding: 12px;
  background-color: #4facfe;
  color: white;
  font-size: 16px;
  border: none;
  border-radius: 8px;
  cursor: pointer;
  transition: background-color 0.3s ease;
}

button:hover {
  background-color: #00f2fe;
}

/* Responsive Design */
@media (max-width: 600px) {
  .register-card {
    width: 90%;
    padding: 20px;
  }
}
</style>
