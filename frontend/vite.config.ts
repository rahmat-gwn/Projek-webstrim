import { fileURLToPath, URL } from 'node:url';
import { defineConfig } from 'vite';
import vue from '@vitejs/plugin-vue';
import dotenv from 'dotenv';

// Memuat file .env
dotenv.config();

// https://vite.dev/config/
export default defineConfig({
  plugins: [vue()],
  define: {
    'process.env': process.env,  // Menyediakan process.env untuk aplikasi
  },
  resolve: {
    alias: {
      '@': fileURLToPath(new URL('./src', import.meta.url)),  // Menambahkan alias @ untuk folder src
    },
  },
});
