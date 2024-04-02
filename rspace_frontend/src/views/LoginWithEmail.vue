<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-12 col-md-5">
                <div class="card">
                    <div class="card-body">
                        
                        <img src="../assets/logo.png"  class=" rounded mx-auto d-block" alt="">
                        

                        <form @submit.prevent="login">
                            <div class="input-group mb-3">
                                <div style="width: 80%; margin: auto;">
                                    <div class="row">
                                        <div class="col-7" style="padding-right: 0px">
                                            <input v-model="email_str" type="text" class="form-control" placeholder="邮箱" aria-label="email" aria-describedby="basic-addon2" required>
                                        </div>
                                        <div class="col-5" style="padding-left: 0px;">
                                            <select class="form-select" aria-label="Default select example" id="email_suffix">
                                                <option value="1" selected>@qq.com</option>
                                                <option value="2">@gmail.com</option>
                                                <option value="3">@163.com</option>
                                            </select>
                                        </div>
                                    </div>
                                </div>
                            </div>
                       
                            <div class="input-group mb-3">
                                <div style="width: 80%; margin: auto;">
                                    <div class="row">
                                        <div class="col-7" style="padding-right: 0px">
                                            <input v-model="email_code" type="text" class="form-control" placeholder="邮箱验证码" aria-label="email_code" aria-describedby="button-addon2" required>
                                        </div>
                                        <div class="col-5" style="padding-left: 0px; border-radius:0px;">
                                            <div style="width: 100%;" class="btn btn-outline-secondary" type="button" id="get_code" @click="get_email_code" :disabled="countdown > 0">
                                                <!-- 发送 -->
                                                <template v-if="countdown === 0">发送</template>
                                                <template v-else>重新发送({{ countdown }}s)</template>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                
                            </div>
                            
                            

                            <!-- d-flex 类将创建一个弹性容器，而 justify-content-center 类将使其内容在水平方向上居中对齐。 -->
                            <div class="d-flex justify-content-center">
                                <button type="submit" class="btn btn-primary position-relative">登录</button>
                            </div>
                        </form>
                        <!-- 下部选项导航栏 -->
                        <div class="row nav">
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'Login'}">账号登录</router-link>|
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'Register'}">注册账号</router-link>|
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'ConnectWithUs'}">联系我们</router-link>
                            
                        </div>
                        <BottomTips></BottomTips>
                    </div>
                </div>
            </div>
        </div>
    </Content>
</template>

<script>
import { ref, reactive } from 'vue'
import Content from '../components/Content.vue'
import BottomTips from '../components/BottomTips.vue'
import { BackendRootURL } from '../common_resources/resource'
import $ from 'jquery'
import { ElMessage } from 'element-plus';
import router from '@/router/index';   //@定位src目录
import { useStore } from 'vuex';
export default{
    name:"LoginWithEmail",
    components:{Content, BottomTips},
    setup(){
        const store = useStore()
        const email_str = ref('')
        const email_suffix = ref('@qq.com')
        const email_code = ref('')
        const countdown = ref(0);
        const email_suffix_map = reactive({
            '0':'',
            '1':'@qq.com',
            '2':'@gmail.com',
            '3':'@163.com',
        })

        //获取邮箱验证码
        const get_email_code = ()=>{
            if(email_str.value === '') {
                ElMessage.warning('请输入邮箱')
                return 
            }
            //然后就根据邮箱发送验证码
            let email_suffix_idx = $('#email_suffix option:selected').val()
            email_suffix.value = email_suffix_map[email_suffix_idx]
            // console.log(email_str.value + email_suffix.value)

            let email_address = email_str.value + email_suffix.value
            $.ajax({
                url: BackendRootURL + '/api/sendemailcode/',
                type:"POST",
                data:{
                    email:email_address,
                    type:"login",
                },
                success(resp) {
                    if(resp.status !== 0) {
                        if(resp.status === -1) {
                            if (resp.message === 'send email code error550 The recipient may contain a non-existent account, please check the recipient address.') {
                                ElMessage.error("此邮箱不存在")
                            }else {
                                ElMessage.error("邮箱获取验证码失败，请稍后重试")
                            }
                        }else if(resp.status === -2) {
                            ElMessage.warning(resp.message)
                            // console.log(resp)
                        }
                        return 
                    }
                    ElMessage.success("成功发送验证码，请注意查收")
                    //成功后都需要等待60秒，发送按钮变成倒计时60秒
                    countdown.value = 60
                    const sendButton = document.getElementById('get_code');
                    sendButton.setAttribute('disabled', 'disabled');
                    sendButton.classList.add('disabled'); // 添加禁用样式

                    const countdownTimer = setInterval(() => {
                        countdown.value--;

                        if (countdown.value === 0) {
                        // Reset the countdown and enable the send button
                            clearInterval(countdownTimer);
                            sendButton.removeAttribute('disabled');
                            sendButton.classList.remove('disabled');
                        }
                    }, 1000);

                    // console.log(resp)
                    // ElMessage.success("你邮箱得到的验证码为："+ resp.data)
                },
                error() {
                    ElMessage.error("邮箱获取验证码失败，请稍后重试")
                }
            })
        }

        //登录
        const login = ()=>{
            //先去后台查找邮箱是否存在，存在再找出对应的账号，密码，生成token
            let email_suffix_idx = $('#email_suffix option:selected').val()
            email_suffix.value = email_suffix_map[email_suffix_idx]
            let email_address = email_str.value + email_suffix.value

            store.dispatch("login", {
                login_method:"login_with_email",
                email:email_address,
                email_code: email_code.value,
                success() {
                    // 登录成功跳转到首页
                    // console.log("success");
                    router.push({name:"Home"});  //跳转路由
                },
                error: (resp) => {
                    // 登录失败
                    if(resp.status !== 0) {
                        if(resp.status === -2) {
                            ElMessage.error(resp.message)
                        }else if(resp.status === -1) {
                            ElMessage.error("邮箱或验证码错误")
                        }
                    }
                }
            })
        }

        return {
            email_str,
            email_suffix,
            email_code,
            countdown,
            email_suffix_map,
            get_email_code,
            login,
        }

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

button{
    width: 80%;
    margin: auto;
}

.input-control {
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
    font-size: 25px;
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