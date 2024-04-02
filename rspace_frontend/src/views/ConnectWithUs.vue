<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-12 col-md-5">
                <div class="card">
                    <div class="card-body">
                        
                        <img src="../assets/logo.png"  class=" rounded mx-auto d-block" alt="">
                        
                        <div class="row">
                            <div class="col-12 warning-title">
                                如果遇到任何问题和BUG，请联系royalwangchao@gmail.com
                            </div>
                        </div>
                        
                        <!-- 下部选项导航栏 -->
                        <div class="row nav">
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'Login'}">账号登录</router-link>|
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'LoginWithEmail'}">邮箱登录</router-link>|
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'Register'}">注册账号</router-link>
                        </div>
                        <BottomTips></BottomTips>
                    </div>
                </div>
            </div>
        </div>
    </Content>
</template>

<script>
import Content from '../components/Content.vue'
export default{
    name:"ConnectWithUs",
    components:{Content},
    setup(){

    }
}

</script>


<style scoped>
/* .card{
    box-shadow: 3px 3px 3px lightgray,  -3px 0 3px lightgray;
} */

img {
    width: 20%;
    margin: auto;
    margin-bottom: 10px;
}

.nav {
    margin-top: 20px;
    display: flex;
    justify-content: center;
    text-align: center;
    font-size: 14px;
    color: #6c757D;
}

.nav-bottom {
    text-decoration: none;
    text-align: center;
    font-size: 14px;
    color: #6c757D;
    padding: 0px;
}

.nav-button {
    cursor: pointer;
}

.nav-button:hover {
    color: rgb(19, 18, 18);
}

.warning-title {
    text-align: center;
    color: rgb(83, 78, 78);
}

</style>