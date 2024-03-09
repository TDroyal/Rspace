<template>
    <nav class="navbar navbar-expand-lg mynav">
        <div class="container">
            <router-link class="navbar-brand" :to="{name:'Home'}">R空间</router-link>
            <button 
                class="navbar-toggler" 
                type="button" 
                data-bs-toggle="collapse" 
                data-bs-target="#navbarText" 
                aria-controls="navbarText" 
                aria-expanded="false" 
                aria-label="Toggle navigation">
                <span class="navbar-toggler-icon"></span>
            </button>
            <div class="collapse navbar-collapse" id="navbarText">
            <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                <li class="nav-item">
                    <router-link class="nav-link" :to="{name:'Home'}">首页</router-link>
                </li>
                <li class="nav-item">
                    <router-link class="nav-link" :to="{name:'Share'}">分享</router-link>
                </li>
                <!-- <li class="nav-item">
                <router-link class="nav-link" href="#">Pricing</router-link>
                </li> -->
            </ul>
            <ul class="navbar-nav" v-if="!$store.state.user.is_login">
                <li class="nav-item">
                    <router-link class="nav-link" :to="{name:'Login'}">登录/注册</router-link>
                </li>
            </ul>

            <!-- 登录成功后，显示用户头像，以及用户名称 -->
            <ul class="navbar-nav" v-else>
                <!-- <li class="nav-item">
                    <router-link class="nav-link" :to="{name:'Login'}">退出</router-link>
                </li> -->
                <div class="nav-item dropdown">
                    <a class="nav-item dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                        <img class="img-fluid avatar" :src="$store.state.user.avatar" alt="">
                        {{$store.state.user.username}}
                    </a>

                    <ul class="dropdown-menu">
                        <li><router-link class="dropdown-item" :to="{name:'Profile', params:{userid:$store.state.user.id}}">个人中心</router-link></li>
                        <!-- <li><div class="dropdown-item" @click="enterProfile($store.state.user.id)">个人中心</div></li> -->
                        <li><router-link class="dropdown-item" :to="{name:'Star'}">我的收藏</router-link></li>
                        <li><hr class="dropdown-divider"></li>
                        <li><router-link class="dropdown-item" :to="{name:'Login'}" @click="logout">退出</router-link></li>
                    </ul>
                </div>
            </ul>

            </div>
        </div>
    </nav>
</template>


<script>

import { useStore } from 'vuex'
// import router from '@/router/index';   //@定位src目录
export default{
    name:"Navbar",
    setup(){
        const store = useStore()
        const logout = ()=>{
            store.commit('logout');//调用user.js全局里面mutations里面的函数用commit
        }

        // const enterProfile = (user_id)=> {
        //     router.push({
        //         name:"Profile",
        //         params:{
        //             userid:user_id,
        //         }
        //     })
        // }


        return {
            logout, 
            // enterProfile,
        }
    }
}


</script>


<style scoped>


.avatar{
    width: 30px; 
    height: 30px; 
    border-radius: 50%;
}

a{
    text-decoration: none; /* 取消下划线 */
    color: black; /* 修改字体颜色，这里以红色为例 */
}

.mynav{
    background-color: white;
    position: fixed;
    top: 0;
    width: 100vw;
    z-index: 1000;
}

</style>