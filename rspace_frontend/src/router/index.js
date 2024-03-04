import { createRouter, createWebHistory } from 'vue-router'
// import Home1 from '../views/Home1.vue'
import Home from '../views/Home.vue'  // 首页

import Login from '../views/Login.vue'

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login',
    name: 'Login',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    component: Login
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
