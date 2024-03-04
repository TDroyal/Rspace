<template>
    <Content>
        <div class="row justify-content-md-center">
            <div class="col-4">
                <div class="card">
                    <div class="card-body">
                        <form @submit.prevent="login">
                            <div class="mb-3">
                                <!-- <label for="username" class="form-label">用户名</label> -->
                                <input v-model="username" type="text" class="form-control" id="username" placeholder="用户名" autocomplete="username" required>
                            </div>
                       
                            <div class="mb-3">
                                <!-- <label for="password" class="form-label">密码</label> -->
                                <input v-model="password" type="password" class="form-control" id="password" placeholder="密码" autocomplete="current-password" required>
                            </div>
                            <!-- <div class="error-message">{{ error_message }}</div> -->
                            <!-- d-flex 类将创建一个弹性容器，而 justify-content-center 类将使其内容在水平方向上居中对齐。 -->
                            <div class="d-flex justify-content-center">
                                <button type="submit" class="btn btn-primary position-relative">登录</button>
                            </div>
                        </form>
                    </div>
                </div>
            </div>
        </div>
    </Content>
</template>


<script>
import { useStore } from 'vuex';
import Content from '../components/Content.vue'
import {ref} from 'vue'
import router from '@/router/index';   //@定位src目录
export default {
    name: "Login",
    components: { Content },

    setup() {
        const store = useStore()
        let username = ref('')
        let password = ref('')
        let error_message = ref('')

        const login = ()=>{
            error_message.value = ""
            store.dispatch("login", {
                username: username.value,
                password:password.value,
                success() {
                    // 登录成功跳转到首页
                    console.log("success");
                    router.push({name:"Home"});  //跳转路由
                },
                error: () => {
                    // 登录失败
                    error_message.value = "用户名或密码错误";
                    // console.log("failed");
                }
            })
        }


        return {
            username,
            password,
            error_message,
            login,
        }
    }

}

</script>


<style scoped>
.card{
    box-shadow: 3px 3px 3px lightgray,  -3px 0 3px lightgray;
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
    
    margin: auto;
}

</style>