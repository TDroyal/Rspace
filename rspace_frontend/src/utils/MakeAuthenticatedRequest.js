//先封装一个判断用户是否处于登录状态的函数：一是判断is_login是否为true，二是判断refresh_token是否过期


// 封装一个判断当前用户的token是否过期，
// 如果过期应该怎么办，如果过期，则利用refresh_token重新去获取新的token(无感刷新token)              再携带该token重新去请求一次后端即可

// import { useStore } from "vuex"
import $ from 'jquery'
import { BackendRootURL } from "../common_resources/resource"
import router from "../router"
import { ElMessage } from "element-plus";

export const CheckIsLogin = async (store) => {
    return new Promise((resolve) => {
        // const store = useStore();
        // console.log(store.state.user.refresh_jwt === '')
        // if (store.state.user.is_login === false && (store.state.user.refresh_jwt === '' || store.state.user.refresh_jwt === null)) {  //说明压根没登录过
        //     // ElMessage.warning("登录状态过期，请重新登录")
        //     store.commit("logout");
        //     router.push({
        //         name: "Login",
        //     });
        //     resolve(false);
        // } else 
        if(store.state.user.refresh_jwt !== ''){
            // console.log(111)
            // if(store.state.user.refresh_jwt === '') {
            //     console.log(222)
            // }
            $.ajax({
                url: BackendRootURL + "/api/checktoken/",
                type: "GET",
                headers: {
                    'Authorization': "Bearer " + store.state.user.refresh_jwt,
                },
                success(resp) {
                    // console.log(resp) // 啥都不返回空白的，说明没过期
                    if (resp.status === -1 || resp.status === 401) {
                    ElMessage.warning("登录状态过期，请重新登录")
                    store.commit("logout");
                    router.push({
                        name: "Login",
                    });
                        resolve(false);  //将结果传递出去
                    } else {
                        resolve(true);
                    }
                },
            
            });
        }
    });
};


// 用refresh_token去刷新token
export const RefreshToken = async (store)=> {  //每隔9分钟去刷新一次jwt
    return new Promise((resolve) => {
        $.ajax({
            url: BackendRootURL + "/api/token/refresh/",
            type: "POST",
            data: {
                refresh_jwt: store.state.user.refresh_jwt,
            },
            success(resp) {  //返回新的jwt
                store.commit("updateJwt", resp.jwt);
                resolve(resp.jwt)
            },
            error() {
                resolve(false)
            }
        })
    })
}


// export default {CheckIsLogin, RefreshToken}