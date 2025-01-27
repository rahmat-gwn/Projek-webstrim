import { createRouter, createWebHistory } from "vue-router";
import LoginView from "@/views/LoginView.vue";
import RegisterView from "@/views/RegisterView.vue";

// Mengimpor halaman yang membutuhkan proteksi
import ProtectedView from "@/views/ProtectedView.vue";

const routes = [
  // Rute untuk login
  { path: "/login", component: LoginView },

  // Rute untuk register
  { path: "/register", component: RegisterView },

  // Rute untuk halaman yang dilindungi
  {
    path: "/protected",
    component: ProtectedView,
    meta: { requiresAuth: true }  // Menandakan bahwa halaman ini memerlukan autentikasi
  }
];

const router = createRouter({
  history: createWebHistory(),
  routes
});

// Menambahkan guard untuk memeriksa apakah pengguna sudah login sebelum mengakses halaman yang memerlukan autentikasi
router.beforeEach((to, from, next) => {
  const isAuthenticated = localStorage.getItem("token");  // Mengecek apakah ada token di localStorage

  // Jika rute yang dituju membutuhkan autentikasi dan pengguna tidak terautentikasi
  if (to.meta.requiresAuth && !isAuthenticated) {
    next("/login");  // Redirect ke halaman login
  } else {
    next();  // Jika tidak membutuhkan autentikasi, lanjutkan ke halaman berikutnya
  }
});

export default router;
