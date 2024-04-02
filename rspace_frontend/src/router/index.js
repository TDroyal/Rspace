import { createRouter, createWebHistory } from 'vue-router'
// import Home1 from '../views/Home1.vue'
import Home from '../views/Home.vue'  // 首页

import Login from '../views/Login.vue'
import LoginWithEmail from '../views/LoginWithEmail.vue'
import Register from '../views/Register.vue'
import ConnectWithUs from '../views/ConnectWithUs.vue'

import Profile from '../views/Profile.vue'
import Post from '../views/Post.vue'
import Star from '../views/Star.vue'
import EditUserInfo from '../views/EditUserInfo.vue'
import PostDetail from '../views/PostDetail.vue'



const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/login/',
    name: 'Login',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    // component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
    component: Login
  },
  {
    path: '/profile/:userid/',
    name: 'Profile',
    component: Profile,
  },
  {
    path: '/share/',
    name: 'Share',
    component: Post,
  },
  {
    path:'/mystar/',
    name:'Star',
    component:Star,
  },
  {
    path:'/edituserinfo/',
    name:'EditUserInfo',
    component:EditUserInfo,
  },
  {
    path:'/post/:postid/',
    name:"PostDetail",
    component:PostDetail,
  },
  {
    path:'/loginwithemail/',
    name:"LoginWithEmail",
    component:LoginWithEmail,
  },
  {
    path:'/register/',
    name:"Register",
    component:Register,
  },
  {
    path:'/connection/',
    name:"ConnectWithUs",
    component:ConnectWithUs,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
