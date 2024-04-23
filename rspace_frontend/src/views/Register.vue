<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-12 col-md-5">
                <div class="card">
                    <div class="card-body">
                        
                        <img src="../assets/logo.png"  class=" rounded mx-auto d-block" alt="">
                        
                        <!-- <select aria-label="Default select example" style="background-color: #F8F9FA;">
                                            
                                            <option selected value="1">@qq.com</option>
                                            <option value="2">@gmail.com</option>
                                            <option value="3">@163.com</option>
                                        </select> -->

                        <form @submit.prevent="register">
                            <!-- 邮箱 + 获取验证码 -->
                            <div class="input-group mb-3">
                                <div style="width: 80%; margin: auto;">
                                    <div class="row">
                                        <div class="col-7" style="padding-right: 0px">
                                            <input v-model="email_str" type="text" class="form-control" placeholder="邮箱" aria-label="email" aria-describedby="basic-addon2" required>
                                        </div>
                                        <div class="col-5" style="padding-left: 0px;">
                                            <!-- <span class="input-group-text" id="basic-addon2">
                                                @qq.com
                                            </span> -->
                                            <!-- <div class="input-group mb-3">
                                                <input type="text" class="form-control" aria-label="Text input with dropdown button">
                                                <button class="btn btn-outline-secondary dropdown-toggle" type="button" data-bs-toggle="dropdown" aria-expanded="false">Dropdown</button>
                                                <ul class="dropdown-menu dropdown-menu-end">
                                                    <li><a class="dropdown-item" href="#">Action</a></li>
                                                    <li><a class="dropdown-item" href="#">Another action</a></li>
                                                    <li><a class="dropdown-item" href="#">Something else here</a></li>
                                                    <li><hr class="dropdown-divider"></li>
                                                    <li><a class="dropdown-item" href="#">Separated link</a></li>
                                                </ul>
                                            </div> -->
                                            <!-- <select class="form-select" aria-label="Default select example" id="email_suffix">
                                                <option value="1" selected>@qq.com</option>
                                                <option value="2">@gmail.com</option>
                                                <option value="3">@163.com</option>
                                            </select> -->
                                            <el-select
                                                v-model="email_suffix_idx"
                                                size="large"
                                                class="custom-select"
                                                style="font-weight: bold;"
                                            >
                                                <el-option 
                                                    v-for="item in email_options"
                                                    :key="item.value"
                                                    :label="item.label"
                                                    :value="item.value"
                                                />
                                            </el-select>
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

                            <div class="mb-3">
                                <!-- <label for="username" class="form-label">用户名</label> -->
                                <input v-model="username" type="text" class="form-control input-control" id="username" placeholder="用户名" autocomplete="username" required>
                            </div>
                       
                            <div class="mb-3">
                                <!-- <label for="password" class="form-label">密码</label> -->
                                <input v-model="password" type="password" class="form-control input-control" id="password" placeholder="密码" autocomplete="current-password" required>
                            </div>
                            <div class="mb-3">
                                <!-- <label for="password" class="form-label">密码</label> -->
                                <input v-model="repeat_password" type="password" class="form-control input-control" id="repeat_password" placeholder="再次输入密码" autocomplete="current-password" required>
                            </div>

                            <!-- d-flex 类将创建一个弹性容器，而 justify-content-center 类将使其内容在水平方向上居中对齐。 -->
                            <div class="d-flex justify-content-center">
                                <button type="submit" class="btn btn-primary position-relative">注册</button>
                            </div>
                        </form>
                        <!-- 下部选项导航栏 -->
                        <div class="row nav">
                            <router-link class=" col-3 nav-bottom nav-button" :to="{name:'Login'}">返回登录</router-link>
                        </div>
                        <BottomTips></BottomTips>
                    </div>
                </div>
            </div>
        </div>
        <div class="bottom-container">
            <div class="row justify-content-center">
                <div class="col-md-7 col-12">
                    <CopyRight></CopyRight>
                </div>
            </div>
        </div>
    </Content>
</template>

<script>

import Content from '../components/Content.vue'
import { ref, reactive } from 'vue'

import { ElMessage } from 'element-plus';
import $ from 'jquery'
import { BackendRootURL } from '../common_resources/resource';
import BottomTips from '../components/BottomTips.vue'
import {useStore} from 'vuex'
import router from '@/router/index';   //@定位src目录
import { jwtDecode } from 'jwt-decode'   //看官网才是最正确的选择
import CopyRight from '../components/CopyRight.vue';
export default{
    name:"Register",
    components:{Content, BottomTips, CopyRight},
    setup(){
        const store = useStore()
        const email_str = ref('')
        const email_suffix = ref('@qq.com')
        const email_code = ref('')
        const username = ref('')
        const password = ref('')
        const repeat_password = ref('')
        //记时
        const countdown = ref(0);
        const email_suffix_map = reactive({
            '0':'',
            '1':'@qq.com',
            '2':'@gmail.com',
            '3':'@163.com',
        })

        const email_options = [
            {
                value: '1',
                label: '@qq.com',
            },
            {
                value: '2',
                label: '@gmail.com',
            },
            {
                value: '3',
                label: '@163.com',
            },
        ]
        const email_suffix_idx = ref('1')
        //获取邮箱验证码
        const get_email_code = ()=>{
            if(email_str.value === '') {
                ElMessage.warning('请输入邮箱')
                return 
            }
            //然后就根据邮箱发送验证码
            // let email_suffix_idx = $('#email_suffix option:selected').val()
            email_suffix.value = email_suffix_map[email_suffix_idx.value]
            // console.log(email_str.value + email_suffix.value)

            let email_address = email_str.value + email_suffix.value
            $.ajax({
                url: BackendRootURL + '/api/sendemailcode/',
                type:"POST",
                data:{
                    email:email_address,
                    type:"register",
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

        // 注册
        const register = ()=>{
            //前端：
            //首先判断两次密码是否一致
            if(password.value !== repeat_password.value) {
                ElMessage.warning("两次密码不一致")
                return 
            }
            //将相关的值传过去，
            $.ajax({
                url: BackendRootURL + "/api/register/",
                type:"POST",
                data:{
                    email: email_str.value + email_suffix.value,
                    email_code: email_code.value,
                    username: username.value,
                    password: password.value,
                },
                success(resp){
                    //后端：
                    //再判断用户名是否已经被注册过
                    //然后再从redis读取此邮箱对应的code
                    //如果读取不存在，则说明验证码已经过期，如果存在，再判断redis里面存的和用户传过去的code是否一致
                    //如果一致则注册，创建新的账号，密码（MD5加密），邮箱，用户名默认为账号，头像默认为灰色头像
                    if(resp.status !== 0) {
                        if(resp.status === -2) {
                            ElMessage.error(resp.message)
                        }else {
                            ElMessage.error("注册失败，请稍后重试")
                        }
                        // console.log(resp)
                        return
                    }
                    //注册成功就直接自动登录，并进入系统
                    ElMessage.success("注册成功")
                    
                    // 获取token 和 refresh_token
                    const jwt = resp.data.token
                    const refresh_jwt = resp.data.refresh_token
                    
                    //对jwt进行解密获取里面的用户信息，包括id, username等
                    const jwt_obj = jwtDecode(jwt)  //对token进行解密
                    // console.log(jwt, refresh_jwt, jwt_obj)

                    setInterval(() => {  //每隔9分钟去刷新一次jwt
                        $.ajax({
                            url: BackendRootURL + "/api/token/refresh/",
                            type: "POST",
                            data: {
                                refresh_jwt:refresh_jwt,
                            },
                            success(resp) {  //返回新的jwt
                                // console.log('9分钟刷新的一次token令牌：', resp);
                                store.commit("updateJwt", resp.jwt);
                            },
                            error(resp) {
                                console.log(resp)
                            },
                        });
                    }, 1000 * 60 * 9);

                    $.ajax({
                        url: BackendRootURL + "/myspace/getuserinfo/",  //就登录成功的时候获取一次
                        type: "GET",
                        data: {
                            user_id: jwt_obj.user_id,
                        },
                        headers: {  //jwt验证
                            'Authorization': "Bearer " + jwt,
                        },
                        success(resp) {
                            if (resp.status != 0) {
                                return 
                            }
                            store.commit("updateUser", {  //传入mutations中的方法名称和参数data
                                id: resp.data.normal_userinfo.id,
                                username: resp.data.normal_userinfo.name,
                                avatar: resp.data.normal_userinfo.avatar,
                                jwt: jwt,
                                refresh_jwt: refresh_jwt,
                                address: resp.data.normal_userinfo.address,
                                introduction: resp.data.normal_userinfo.introduction,
                                age: resp.data.normal_userinfo.age,
                                gender: resp.data.normal_userinfo.gender,
                                is_login: true,
                                fanscount:resp.data.fanscount,
                                followercount:resp.data.fanscount,
                            });
                            router.push({name:"Home"});  //跳转路由
                        },
                        
                        error(resp) {
                            console.log(resp)
                            // data.error();
                        }
                    });

                },
                error() {
                    ElMessage.error("注册失败，请稍后重试")
                    return
                }
            })
            
        }

        return {
            email_str,
            email_suffix,
            email_suffix_idx,
            email_options,
            email_code,
            username,
            password,
            repeat_password,
            countdown,
            get_email_code,
            register,
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

.bottom-container {
  position: fixed;
  left: 0;
  bottom: 0;
  width: 100%;
  padding: 20px;
}

</style>