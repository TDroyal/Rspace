<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-12 col-md-5">
                <div class="card">
                    <div class="card-body">
                        
                        <img src="../assets/logo.png"  class=" rounded mx-auto d-block" alt="">
                        

                        <form @submit.prevent="login">
                            <div class="mb-3">
                                <!-- <label for="username" class="form-label">用户名</label> -->
                                <input v-model="username" type="text" class="form-control" id="username" placeholder="用户名" autocomplete="username" required>
                            </div>
                       
                            <div class="mb-3">
                                <!-- <label for="password" class="form-label">密码</label> -->
                                <input v-model="password" type="password" class="form-control" id="password" placeholder="密码" autocomplete="current-password" required>
                            </div>
                            
                            <div class="mb-3">
                                <!-- <label for="password" class="form-label">密码</label> -->
                                <div class="row" style="display: flex; justify-content: center;">
                                    <div class="col-6" style="padding-right: 0px;">
                                        <input v-model="codestr" type="text" class="form-control" id="codestr" placeholder="验证码" autocomplete="codestr" maxlength="4" required>
                                    </div>
                                    <div class="col-5 randcode" style="padding-left: 20px;">
                                        <div style="width: 70%; text-align: center;" @click="refreshRandomCode" class="randcodestr">{{randcodestr}}</div>
                                    </div>
                                    <div class="error-message">{{ error_message }}</div>
                                </div>
                            </div>

                            <!-- d-flex 类将创建一个弹性容器，而 justify-content-center 类将使其内容在水平方向上居中对齐。 -->
                            <div class="d-flex justify-content-center">
                                <button type="submit" class="btn btn-primary position-relative">登录</button>
                            </div>
                        </form>
                        <!-- 下部选项导航栏 -->
                        <div class="row nav">
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'LoginWithEmail'}">邮箱登录</router-link>|
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'Register'}">注册账号</router-link>|
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'ConnectWithUs'}">联系我们</router-link>
                            <!-- <div class="col-3 nav-bottom nav-button">
                                邮箱登录
                            </div>|
                            <div class="col-3 nav-bottom nav-button">
                                注册
                            </div>|
                            <div class="col-3 nav-bottom nav-button">
                                联系我们
                            </div> -->
                        </div>
                        <BottomTips></BottomTips>
                    </div>
                </div>
            </div>
        </div>
    </Content>
</template>


<script>
import { useStore } from 'vuex';
import Content from '../components/Content.vue'
import BottomTips from '../components/BottomTips.vue'
import {ref} from 'vue'
// import { ElMessage } from 'element-plus';
import router from '@/router/index';   //@定位src目录
import { ElMessage } from 'element-plus';
export default {
    name: "Login",
    components: { Content, BottomTips},

    setup() {
        const store = useStore()
        let username = ref('')
        let password = ref('')
        let error_message = ref('')
        const codestr = ref('')
        const randcodestr = ref('')

        const generateRandomCode = ()=>{
            let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789'
            let n = characters.length
            let code = ''
            for (var i = 0; i < 4; i++) {
                var randomIndex = Math.floor(Math.random() * n);
                code += characters.charAt(randomIndex);
            }
            
            return code;
        }
        randcodestr.value = generateRandomCode()
        const refreshRandomCode = ()=>{
            randcodestr.value = generateRandomCode()
        }


        const login = ()=>{
            //随机生成的验证码为数组形式，此处将其转为字符串并去掉中间相隔的逗号
            if(codestr.value !== randcodestr.value) {
                error_message.value = "请输入正确的验证码"
                // ElMessage.error("请输入正确的验证码")
                refreshRandomCode()
                return
            }

            error_message.value = ""
            store.dispatch("login", {
                login_method:"login_with_username",
                username: username.value,
                password:password.value,
                success() {
                    // 登录成功跳转到首页
                    // console.log("success");
                    router.push({name:"Home"});  //跳转路由
                },
                error: () => {
                    // 登录失败
                    ElMessage.error("用户名或密码错误")
                    error_message.value = "用户名或密码错误";
                    refreshRandomCode()
                    // console.log("failed");
                }
            })
        }


        return {
            username,
            password,
            error_message,
            randcodestr,
            codestr,
            login,
            generateRandomCode,
            refreshRandomCode,
        }
    }

}

</script>


<style scoped>
.card{
    box-shadow: 3px 3px 3px lightgray,  -3px 0 3px lightgray;
}

img {
    width: 20%;
    margin: auto;
    margin-bottom: 10px;
}

button{
    width: 80%;
    margin: auto;
}

input {
    width: 80%;
    margin: auto;
}

.error-message {
    color: red;
    /* margin-top: 0px; */
    width: 80%;
    margin: auto;
    margin-top: 5px;
}

.randcodestr{
    background: lightgray;
}

.randcodestr:hover {
    cursor: pointer;
}

.randcode {
    display: flex; 
    align-items: center; 
    justify-content: center; 
    font-weight: bold;
    font-size: 20px;
    user-select: none;
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


</style>