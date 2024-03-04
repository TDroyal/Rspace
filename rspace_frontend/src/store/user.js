import $ from 'jquery'

import { jwtDecode } from 'jwt-decode'   //看官网才是最正确的选择

const ModuleUser = {
    state: {  //存全局数据
        id: "",
        username: "",
        photo: "",
        followerCount: 0,
        jwt: "",     // 客户端存储token,并在请求头中携带Token
        refresh_jwt: "",
        is_login: false,
    },
    mutations: {     //更新state里面的数据
        updataUser(state, user) {
            state.id = user.id;
            state.username = user.username;
            state.photo = user.photo;
            state.followerCount = user.followerCount;
            state.jwt = user.jwt;
            state.refresh_jwt = user.refresh_jwt;
            state.is_login = user.is_login;
        },

        updataJwt(state, jwt) {
            state.jwt = jwt;
        },

        logout(state) {
            state.id = "";
            state.username = "";
            state.photo = "";
            state.followerCount = 0;
            state.jwt = "";
            state.refresh_jwt = "";
            state.is_login = false;
        },
    },
    actions: {      //对state的各种操作
        //context里面传一些api，data传信息
        login: (context, data) => {
            console.log(context)
            $.ajax({
                url: "http://127.0.0.1:9090/api/token/",
                type: "POST",
                data: {
                    username: data.username,
                    password: data.password,
                },
                success(resp) {
                    console.log(resp);
                    const jwt = resp.data.token
                    console.log(jwt)
                    const jwt_obj = jwtDecode(jwt)  //对token进行解密
                    console.log(jwt, jwt_obj)
                    // const { access, refresh } = resp;
                    // const access_obj = jwt_decode(access);  //对jwt解密获取id
                    // // console.log(access_obj, refresh);

                    // setInterval(() => {  //每隔4.5分钟去刷新一次jwt
                    //     $.ajax({
                    //         url: "http://127.0.0.1:9090/api/token/refresh/",
                    //         type: "POST",
                    //         data: {
                    //             refresh,
                    //         },
                    //         success(resp) {
                    //             // console.log(resp);
                    //             context.commit("updataAccess", resp.access);
                    //         }
                    //     });
                    // }, 1000 * 60 * 4.5);

                    // $.ajax({
                    //     url: "http://127.0.0.1:9090/myspace/getinfo/",
                    //     type: "GET",
                    //     data: {
                    //         user_id: access_obj.user_id,
                    //     },
                    //     headers: {  //jwt验证
                    //         'Authorization': "Bearer " + access,
                    //     },
                    //     success(resp) {
                    //         console.log(resp);
                    //         context.commit("updataUser", {  //传入mutations中的方法名称和参数data
                    //             ...resp,  //...把数组或对象展开成一系列用逗号隔开的值
                    //             access: access,
                    //             refresh: refresh,
                    //             is_login: true,
                    //         });
                    //         data.success();  //调用login.vue里面的回调函数
                    //     },
                    // });
                },
                error(resp) {
                    console.log(resp.responseJSON)  //后端相应的json
                    data.error();
                }
            });
        },
    },
    modules: {
    }
}

export default ModuleUser