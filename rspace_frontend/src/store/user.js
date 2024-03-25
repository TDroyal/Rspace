import $ from 'jquery'

import { jwtDecode } from 'jwt-decode'   //看官网才是最正确的选择
import { ElMessage } from 'element-plus'
import BackendRootURL from '../common_resources/resource'

const ModuleUser = {
    state: {  //存全局数据
        id: "",
        username: "",
        avatar: "",
        // followerCount: 0,
        jwt: "",     // 客户端存储token,并在请求头中携带Token  jwt一般比较短，就几分钟，或者1小时
        refresh_jwt: "",   // 刷新令牌  有效期一般比较长，设置7天
        address: "",
        introduction: "",
        age: "",
        gender: "",
        fanscount:0,
        followercount:0,
        is_login: false,
    },
    mutations: {     //更新state里面的数据
        updateUser(state, user) {
            state.id = user.id;
            state.username = user.username;
            state.avatar = BackendRootURL + "/static/avatar/" + user.avatar;
            // state.followerCount = user.followerCount;
            state.jwt = user.jwt;
            state.refresh_jwt = user.refresh_jwt;
            state.address = user.address
            state.introduction = user.introduction
            state.age = user.age
            state.gender = user.gender
            state.fanscount = user.fanscount
            state.followercount = user.followercount
            state.is_login = user.is_login;
            let uu = {
                user:state
            }
            localStorage.setItem("userStore", JSON.stringify(uu))
            // console.log("updateUser", uu)
        },

        updateJwt(state, jwt) {
            state.jwt = jwt;
            let uu = {
                user:state
            }
            localStorage.setItem("userStore", JSON.stringify(uu))
            // console.log("updateJwt", uu)
        },

        //单独更新头像
        updateAvatar(state, avatar)
        {
            state.avatar = BackendRootURL + "/static/avatar/" + avatar
            let uu = {
                user:state
            }
            localStorage.setItem("userStore", JSON.stringify(uu))
            // console.log("updateAvatar", uu)
        },

        logout(state) {
            state.id = "";
            state.username = "";
            state.avatar = "";
            // state.followerCount = 0;
            state.jwt = "";
            state.refresh_jwt = "";
            state.address = ""
            state.introduction = ""
            state.age = ""
            state.gender = ""
            state.fanscount = 0
            state.followercount = 0
            state.is_login = false;

            window.localStorage.removeItem("userStore")
        },
    },
    actions: {      //对state的各种操作
        //context里面传一些api，data传信息
        login: (context, data) => {
            // console.log(context)
            $.ajax({
                url: BackendRootURL + "/api/token/",
                type: "POST",
                data: {
                    username: data.username,
                    password: data.password,
                },
                success(resp) {   //应该返回一个jwt和一个刷新令牌
                    // console.log(resp);
                    if(resp.status != 0) {
                        data.error();
                        return 
                    }
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
                                context.commit("updateJwt", resp.jwt);
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
                            // console.log(resp)

                            // let gender = ""
                            // if (resp.data.gender === 1) {
                            //     gender = "男"
                            // }
                            // if (resp.data.gender === 2) {
                            //     gender = "女"
                            // }

                            context.commit("updateUser", {  //传入mutations中的方法名称和参数data
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
                            ElMessage({
                                message: '登录成功',
                                type: 'success',
                            })
                            data.success();  //调用login.vue里面的回调函数
                        },
                        
                        error(resp) {
                            console.log(resp)
                            // data.error();
                        }
                    });
                },
                error(resp) {
                    ElMessage({
                        message: '登录失败，请稍后重试',
                        type: 'error',
                    })
                    console.log(resp.responseJSON)  //后端相应的json
                    // data.error();
                }
            });
        },
    },
    modules: {
        
    }
}

export default ModuleUser