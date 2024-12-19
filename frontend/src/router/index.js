import { createRouter, createWebHashHistory } from 'vue-router'
import HomeView from '../views/LoginView.vue'

const routes = [
  {
    path: '/',
    name: 'Login',
    component: HomeView
  }
]

const router = createRouter({
  history: createWebHashHistory(),
  routes
})

export default router
