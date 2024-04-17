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
import Notifications from '../views/Notifications.vue'
import ProfilePostLists from '../views/ProfilePostLists.vue'
import ProfileStar from '../views/ProfileStar.vue'
import ProfileFollowList from '../views/ProfileFollowList.vue'
import ProfileFanList from '../views/ProfileFanList.vue'

const routes = [
  // {
  //   path: '/',
  //   name: 'Home',
  //   component: Home
  // },
  {
    path: '/',
    name: 'Home',
    component: Home,
    beforeEnter: (to, from, next) => {
      const currentPage = to.query.page || 1;
      // console.log(currentPage)
      if (currentPage === 1) {
        // next({path: '/'}); // 在第一页时继续导航到组件
        next();
      } else {
        next({ path: `/?page=${currentPage}` }); // 在其他页时重定向到新路径
      }
    }
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
    redirect: to => {
      return `/profile/${to.params.userid}/posts/`; //进入profile页面后选择显示的默认页面
    },  
    children: [
      {
        path: 'posts/',
        name: 'ProfilePostLists',
        component: ProfilePostLists,
      },
      {
        path: 'star/',
        name: 'ProfileStar',
        component: ProfileStar,
      },
      {
        path: 'following/',
        name: 'ProfileFollowList',
        component: ProfileFollowList,
      },
      {
        path: 'fans/',
        name: 'ProfileFanList',
        component: ProfileFanList,
      },
    ]
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
  {
    path:'/notifications/',
    name:"Notifications",
    component:Notifications,
  },
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
