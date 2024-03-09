import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'

import 'jquery/dist/jquery.min'
import "bootstrap/dist/css/bootstrap.css"
import "bootstrap/dist/js/bootstrap"


import ElementPlus from 'element-plus'
import 'element-plus/dist/index.css'

// // 添加全局编译时特性标志定义
// window.__VUE_PROD_DEVTOOLS__ = false; // 禁用Vue开发工具
// window.__VUE_PROD_HYDRATION__ = true; // 启用Vue的生产环境渲染


createApp(App).use(store).use(router).use(ElementPlus).mount('#app')
