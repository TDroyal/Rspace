<template>
    <nav class="navbar navbar-expand-lg mynav">
        <div class="container">
            <router-link class="navbar-brand" :to="{name:'Home', query:{page:1}}">
                <svg t="1711091592185" class="icon" viewBox="0 0 1162 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="11316" width="200" height="200"><path d="M1162 453.2C1162 238.2 901.8 64 581 64S0 238.2 0 453.2C0 644.8 206.6 804 478.8 836.2V960h198.2v-123c48.6-5.4 95.2-14.8 138.8-27.8L896 960h224l-134.8-227.4c109-70.8 176.8-169.8 176.8-279.4z m-933.6 29c0-147 197.8-266 441.6-266s423.8 81.4 423.8 266c0 100.2-53 170-140.6 212.8-4.8-3.2-9.4-5.8-12.8-7.4-20.4-10.4-55.6-21-55.6-21s173.2-12.8 173.2-185.4-181.2-175.8-181.2-175.8h-398V722c-148.2-43-250.4-134.2-250.4-239.8z m450.2 76.6v-111.2c115.6 0 175.6-13.6 175.6 54.6 0 73-76.4 56.6-175.6 56.6z m-1.8 145H730c21.6 0 37.8 23.4 48 38.4-32.2 3.8-66 5.6-101.2 5.8v-44.2z" fill="#2c2c2c" p-id="11317"></path></svg>
                <span style="font-weight: bold;">空间</span>
            </router-link>
            <!-- <a href="/?page="  class="navbar-brand">
                <svg t="1711091592185" class="icon" viewBox="0 0 1162 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="11316" width="200" height="200"><path d="M1162 453.2C1162 238.2 901.8 64 581 64S0 238.2 0 453.2C0 644.8 206.6 804 478.8 836.2V960h198.2v-123c48.6-5.4 95.2-14.8 138.8-27.8L896 960h224l-134.8-227.4c109-70.8 176.8-169.8 176.8-279.4z m-933.6 29c0-147 197.8-266 441.6-266s423.8 81.4 423.8 266c0 100.2-53 170-140.6 212.8-4.8-3.2-9.4-5.8-12.8-7.4-20.4-10.4-55.6-21-55.6-21s173.2-12.8 173.2-185.4-181.2-175.8-181.2-175.8h-398V722c-148.2-43-250.4-134.2-250.4-239.8z m450.2 76.6v-111.2c115.6 0 175.6-13.6 175.6 54.6 0 73-76.4 56.6-175.6 56.6z m-1.8 145H730c21.6 0 37.8 23.4 48 38.4-32.2 3.8-66 5.6-101.2 5.8v-44.2z" fill="#2c2c2c" p-id="11317"></path></svg>
                <span style="font-weight: bold; text-align: center;">空间</span>
            </a> -->
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
                    <router-link class="nav-link mynav-link" :to="{name:'Home', query:{page:1}}">首页</router-link>
                    <!-- <a href="/?page=1"  class="nav-link mynav-link">首页</a> -->
                </li>
                <li class="nav-item" v-if="$store.state.user.is_login">
                    <router-link class="nav-link mynav-link" :to="{name:'Share'}">分享</router-link>
                </li>
                <!-- <li class="nav-item">
                <router-link class="nav-link" href="#">Pricing</router-link>
                </li> -->
            </ul>
            <ul class="navbar-nav" v-if="!$store.state.user.is_login">
                <li class="nav-item">
                    <router-link class="nav-link mynav-link" :to="{name:'Login'}">登录/注册</router-link>
                </li>
            </ul>

            <!-- 登录成功后，显示用户头像，以及用户名称 -->
            <ul class="navbar-nav" v-else>
                <!-- <li class="nav-item">
                    <router-link class="nav-link" :to="{name:'Login'}">退出</router-link>
                </li> -->
                <div class="nav-item dropdown">
                    <a class="nav-item dropdown-toggle mynav-link" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                        <img class="img-fluid avatar" :src="$store.state.user.avatar" alt="">
                        {{$store.state.user.username}}
                        <!-- <div class="mynav-link">{{$store.state.user.username}}</div> -->
                    </a>

                    <ul class="dropdown-menu">
                        <li><router-link class="dropdown-item" :to="{name:'Profile', params:{userid:$store.state.user.id}}">个人中心</router-link></li>
                        <!-- <li><div class="dropdown-item" @click="enterProfile($store.state.user.id)">个人中心</div></li> -->

                        <!-- <li><router-link class="dropdown-item" :to="{name:'Star'}">我的收藏</router-link></li> -->
                        
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
            store.commit('logout_page')
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
    /* background-color: #F7F7F7; */
    position: fixed;
    top: 0;
    width: 100vw;
    z-index: 1000;
}

svg {
    max-width: 30px;
    max-height: 30px;
}

.mynav-link{
    color: rgb(107, 107, 107);
}

.mynav-link:hover{
    color: black;
}

</style>