<template>
    <Content v-if="$store.state.user.is_login === true">
        <div class="card card-out">
            <div class="card-body">
                <div class="row justify-content-center">
                    <div class="col-md-2 col-6">
                        <el-upload
                            action="#"
                            :http-request="httpRequest"   
                            :show-file-list="false"
                            :before-upload="beforeAvatarUpload"
                        >   
                            <img class="img-fluid my-avatar" :src="$store.state.user.avatar" alt="" >
                            <span class="avatar-text">更换头像</span>
                        </el-upload>
                        <!-- <img class="img-fluid my-avatar" :src="$store.state.user.avatar" alt="" >
                        <span class="avatar-text">更换头像</span> -->
                        <!-- <div class="text-center username">{{$store.state.user.username}}</div> -->
                    </div>
                </div>
                <form @submit.prevent="modify_userinfo" style="margin-top: 20px;">
                    <div class="row justify-content-center" >
                        <div class="col-md-5 col-12">
                            <div class="mb-3">
                                <label for="username" class="form-label">昵称</label>
                                <input type="text" class="form-control" id="username" placeholder="你的昵称" v-model="userinfo.username">
                            </div>
                        </div>
                        <div class="col-md-5 col-12">
                            <div class="mb-3">
                                <label for="gender" class="form-label">性别</label>
                                <div class="row" id="gender">
                                    <div class="col-6 text-center">
                                        <div class="gender-select form-control" @click="change_gender(1)" :style="{ backgroundColor: gender_select === 1 ? '#EBF5FF' : '#F7F7F7', color: gender_select === 1 ? '#007AFF' : '#262626BF'}">
                                            <svg t="1709730743916" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1525" width="200" height="200"><path d="M511.843434 512m-446.708971 0a446.708971 446.708971 0 1 0 893.417942 0 446.708971 446.708971 0 1 0-893.417942 0Z" fill="#FFFFFF" p-id="1526"></path><path d="M424.111301 818.825061c-59.328223 0-115.140367-23.107277-157.101038-65.081251-86.620823-86.620823-86.620823-227.581252 0-314.215378 41.960671-41.973974 97.771791-65.081251 157.101038-65.081251 59.355853 0 115.140367 23.12058 157.101037 65.081251 41.960671 41.973974 65.081251 97.771791 65.081251 157.11434s-23.12058 115.140367-65.081251 157.101038-97.745185 65.081251-157.101037 65.081251z m0-360.620268c-36.97103 0-71.733765 14.409175-97.881285 40.543392-53.957913 53.984518-53.957913 141.804656 0 195.775872 26.14752 26.14752 60.910255 40.543392 97.881285 40.543391s71.733765-14.395872 97.881284-40.543391c26.14752-26.14752 40.543392-60.910255 40.543392-97.881285s-14.395872-71.733765-40.543392-97.894587c-26.146497-26.14752-60.909232-40.543392-97.881284-40.543392z" fill="#75B9EB" p-id="1527"></path><path d="M551.602973 511.016603c-10.715039 0-21.430078-4.090155-29.609365-12.269442-16.358573-16.358573-16.358573-42.874483 0-59.219753L672.577209 288.943808h-42.833551c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878806h143.958716c16.931624 0 32.200376 10.210549 38.689161 25.847691 6.488785 15.650445 2.889817 33.67189-9.078773 45.641503L581.212338 498.747161c-8.179286 8.179286-18.894326 12.269441-29.609365 12.269442z" fill="#75B9EB" p-id="1528"></path><path d="M773.703397 288.943808h-143.958716c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878806h143.958716c23.12058 0 41.878806 18.744923 41.878806 41.878806s-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1529"></path><path d="M779.864724 439.050548c-23.12058 0-41.878806-18.744923-41.878806-41.878806V253.226329c0-23.133883 18.758226-41.878806 41.878806-41.878807s41.878806 18.744923 41.878806 41.878807v143.945413c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1530"></path><path d="M779.864724 439.050548c-23.12058 0-41.878806-18.744923-41.878806-41.878806V253.226329c0-23.133883 18.758226-41.878806 41.878806-41.878807s41.878806 18.744923 41.878806 41.878807v143.945413c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1531"></path></svg>
                                            男
                                        </div>
                                    </div>
                                    <div class="col-6 text-center">
                                        <div class="gender-select form-control" @click="change_gender(2)" :style="{ backgroundColor: gender_select === 2 ? '#FFEEF9' : '#F7F7F7' , color: gender_select === 2 ? '#FF74D0' : '#262626BF'}">
                                            <svg  t="1709731126558" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1669" width="200" height="200"><path d="M510.887666 512m-446.708971 0a446.708971 446.708971 0 1 0 893.417942 0 446.708971 446.708971 0 1 0-893.417942 0Z" fill="#FFFFFF" p-id="1670"></path><path d="M510.87948 578.902736c-123.673717 0-224.282113-100.607372-224.282113-224.282113s100.607372-224.282113 224.282113-224.282113 224.282113 100.607372 224.282112 224.282113-100.608396 224.282113-224.282112 224.282113z m0-364.80559c-77.486792 0-140.523477 63.036685-140.523477 140.523477s63.036685 140.523477 140.523477 140.523477 140.523477-63.036685 140.523477-140.523477-63.036685-140.523477-140.523477-140.523477z" fill="#FF3EC9" p-id="1671"></path><path d="M510.87948 896.635217c-23.12058 0-41.878806-18.744923-41.878806-41.878806V537.02393c0-23.133883 18.758226-41.878806 41.878806-41.878806s41.878806 18.744923 41.878806 41.878806v317.732481c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#FF3EC9" p-id="1672"></path><path d="M669.752884 737.762837H352.033705c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878807h317.719179c23.12058 0 41.878806 18.744923 41.878806 41.878807s-18.758226 41.878806-41.878806 41.878806z" fill="#FF3EC9" p-id="1673"></path></svg>
                                            女
                                        </div>
                                    </div>
                                </div>
                                <!-- <input type="text" class="form-control" id="exampleFormControlInput1" placeholder="请选择你的性别" v-model="userinfo.gender" > -->
                            </div>
                        </div>
                        <div class="col-md-5 col-6">
                            <div class="mb-3">
                                <label for="age" class="form-label">年龄</label>
                                <input type="number" class="form-control" id="age" placeholder="你的年龄" v-model="userinfo.age">
                            </div>
                        </div>
                        <div class="col-md-5 col-6">
                            <!-- <div class="mb-3">
                                <label for="exampleFormControlInput1" class="form-label">地址</label>
                                <input type="text" class="form-control" id="exampleFormControlInput1" placeholder="你所在地" v-model="userinfo.address">
                            </div> -->
                            <div class="mb-3">
                                <label for="address" class="form-label">地址</label>
                                <select class="form-select form-select-md" aria-label=".form-select-md example" id="address">
                                    <option value="0" :selected="address_id === 0">你的地址</option>
                                    <option value="1" :selected="address_id === 1">成都</option>
                                    <option value="2" :selected="address_id === 2">重庆</option>
                                    <option value="3" :selected="address_id === 3">巴中</option>
                                </select>
                            </div>
                        </div>
                        <div class="col-md-10 col-12">
                            <div class="mb-3">
                                <label for="introduction" class="form-label">个人介绍</label>
                                <textarea class="form-control" id="introduction" rows="3" placeholder="关于你的个性、兴趣..." v-model="userinfo.introduction"></textarea>
                            </div>
                        </div>
                    </div>
                    <div class="row justify-content-center" style="margin-top: 20px;">
                        <div class="col-md-3 col-12 d-flex">
                            <button type="submit" class="btn position-relative">保存</button>
                        </div>    
                    </div>
                </form>
                
            </div>
        </div>
    </Content>
</template>

<script>
import { reactive, ref } from 'vue';
import Content from '../components/Content.vue'
import {useStore} from 'vuex'
import $ from 'jquery'
import { ElMessage } from 'element-plus'
import {BackendRootURL} from '../common_resources/resource';
import router from '@/router/index';   //@定位src目录
export default {
    name:"EditUserInfo",
    components:{Content},
    setup(){
        const store = useStore()
        const check_is_login = ()=>{
            if(store.state.user.is_login === false) {
                router.push({
                    name:"Login",
                })
                return false
            }
            return true
        }
        if(check_is_login() === false) {
            return
        }

        

        const userinfo = reactive({  //reactive不需要使用.value
            ...store.state.user
        })
        // console.log('---------', userinfo)
        let gender_select = ref(userinfo.gender)  // 1是男  2是女
        // console.log(userinfo.gender, gender_select.value)
        

        // 从userinfo中得到address_id
        const address_map_reverse = reactive({
            '':0,
            '成都':1,
            '重庆':2,
            '巴中':3,
        })


        let address_id = address_map_reverse[userinfo.address]

        // 下拉框的value映射的address
        const address_map = reactive({
            '0':'',
            '1':'成都',
            '2':'重庆',
            '3':'巴中',
        })

        // const gender_map = reactive({
        //     0:'',
        //     1:'男',
        //     2:'女',
        // })

        // if(userinfo.gender === '男') {
        //     gender_select.value = 1
        // }else if (userinfo.gender === '女') {
        //     gender_select.value = 2
        // }

        const modify_userinfo = ()=>{
            if(userinfo.username === '') {
                // console.log("昵称不能为空")
                ElMessage({
                    message: '昵称不能为空',
                    type: 'warning',
                })
                return 
            }

            let address_id_ = $('#address option:selected').val()
            userinfo.address = address_map[address_id_]
            // console.log(userinfo)

            // 通过post方法将修改后的userinfo传回去，不将头像传回去
            $.ajax({
                url: BackendRootURL + '/myspace/updateuserinfo/',
                type:'POST',
                data: {
                    // id:userinfo.id,
                    name:userinfo.username,
                    gender:gender_select.value,
                    age:userinfo.age,
                    address:userinfo.address,
                    introduction:userinfo.introduction,
                },
                headers: {  //jwt验证
                    'Authorization': "Bearer " + userinfo.jwt,
                },
                success(resp) {
                    // 弹窗更新成功
                    // console.log(resp)
                    if(resp.status != 0) {
                        // console.log("更新失败")
                        ElMessage({
                            message: '更新失败，请稍后重试',
                            type: 'error',
                        })
                        return 
                    }
                    store.commit("updateUser", {  //传入mutations中的方法名称和参数data
                        id: resp.data.id,
                        username: resp.data.name,
                        avatar: resp.data.avatar,
                        jwt: userinfo.jwt,
                        refresh_jwt: userinfo.refresh_jwt,
                        address: resp.data.address,
                        introduction: resp.data.introduction,
                        age: resp.data.age,
                        gender: resp.data.gender,
                        is_login: true,
                    });
                    ElMessage({
                        message: '更新成功',
                        type: 'success',
                    })
                },
                error(err) {
                    //更新失败
                    console.log(err)
                    ElMessage({
                        message: '更新失败，请稍后重试',
                        type: 'error',
                    })
                    return 
                    
                }
            })

        }

        const change_gender = (gender)=> {
            gender_select.value = gender
            // userinfo.gender = gender_map[gender_select.value]
        }

        const beforeAvatarUpload = (rawFile) => {
            // console.log(rawFile)
            if (rawFile.type !== 'image/jpg' && rawFile.type !== 'image/png' && rawFile.type !== 'image/jpeg') {
                // 不是这个类型的就应该删掉该文件
                ElMessage.warning('图片类型只能为JPG、PNG、JPEG格式')
                
                return false
            }
            if (rawFile.size / 1024 / 1024 > 2) {
                // console.log('caonm')
                ElMessage.warning('图片大小不能超过2MB')
                return false
            }
            return true  //返回true才能上传图片
        }

        const httpRequest = (data)=>{   //覆盖默认的 Xhr 行为，允许自行实现上传文件的请求
            const formData = new FormData();
            formData.append('avatar', data.file);
            $.ajax({
                url: BackendRootURL + "/myspace/updateavatar/",
                type:"POST",
                data: formData,
                // 下面两行必须要上，不然报错
                processData: false,
                contentType: false,
                headers:{
                    'Authorization': "Bearer " + userinfo.jwt,
                },
                success(resp) {
                    if(resp.status !== 0) {
                        ElMessage.error("更换头像失败，请稍后重试!")
                        return
                    }
                    ElMessage.success("更换头像成功")
                    store.commit("updateAvatar", resp.data)
                },
                error(resp) {
                    ElMessage.error("更换头像失败，请稍后重试!")
                    console.log(resp)
                }
            })
        }

        return {
            // store,
            userinfo,
            gender_select,
            address_map,
            address_map_reverse,
            address_id,
            // gender_map,
            modify_userinfo,
            change_gender,
            beforeAvatarUpload,
            httpRequest,
        }
    },
}

</script>

<style scoped>

.card-out {
    margin-top: 20px;
    /* box-shadow: none; */
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
    border: none;
}

.my-avatar{
    border-radius: 50%;
}

.my-avatar:hover{
    cursor: pointer;
}

.username{
    margin-top:5px;
    font-weight: bold;
}

button{
    width: 100%;
    margin: auto;
    background-color: #EFF9F2;
    color: #2DB55D;
    border: none;
}

button:hover {
    background-color: #92e4aa;
    color: white;
}

input {
    background-color: #F7F7F7;
}

.form-select{
    background-color: #F7F7F7;
}

textarea {
    background-color: #F7F7F7;
}

.gender-select {
    background-color: #F7F7F7;

}

.gender-select:hover {
    cursor: pointer;

}

svg {
    max-height: 20px;
    max-width: 20px;
}

/* option {
    min-height: 50px;
} */

/* .myoption {
    margin-top: 20px;
    
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
    border: none;
} */

/* 在下面的代码中，我们添加了一个<span>元素作为my-avatar的兄弟元素，
    并设置其内容为"更换头像"。然后，我们使用CSS样式来控制avatar-text的显示和位置。
    初始状态下，我们将avatar-text的display属性设置为none，以隐藏它。
    当鼠标悬停在my-avatar上时，我们使用兄弟选择器(+)选择到avatar-text，
    并将其display属性设置为block，以显示它。 */
.avatar-text {
    /* display不显示 */
    display: none;
    position: absolute;
    top: 20%;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: rgba(0, 0, 0, 0.8);
    color: white;
    padding: 6px 12px;
    border-radius: 4px;
}

.my-avatar:hover + .avatar-text {
    display: block;
}

</style>