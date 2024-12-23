import { createRouter, createWebHashHistory } from 'vue-router'
import RegisterComp from "@/components/RegisterComp.vue";
import LoginView from "@/views/LoginView.vue";

const routes = [
  {
    path: '/',
    name: 'Login',
    component: LoginView
  },
  {
    path: '/register',
    name: 'Register',
    component: RegisterComp
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
