<template>
    <!-- 头像基本信息 -->
    <div class="card card-out">
        <div class="card-body" style="background-color: #F7F8FA;">
            <div class="card">
                <div class="card-body">
                    <div class="row justify-content-center">
                        <div class="col-4">
                            <!-- 点击头像进入编辑个人资料页面 -->
                            <img class="img-fluid" :src="userinfo.avatar" alt="">
                        </div>
                        <div class="col-8">
                            <div class="username">{{userinfo.username}}</div>
                            <div class="gender">{{gender_map[userinfo.gender]}}</div>
                        </div>
                    </div>
                    <div class="row justify-content-center" style="margin-top: 10px;">
                        <div class="col-6">
                            <div class="text-center">关注者</div>
                            <div class="text-center">2</div>
                        </div>|
                        <div class="col-5">
                            <div class="text-center">关注了</div>
                            <div class="text-center">10</div>
                        </div>
                    </div>
                    <!-- 点击头像进入编辑个人信息 -->
                    <div class="row justify-content-center" style="margin-top: 10px;">
                        <!-- d-flex  position-relative margin:auto; 就可以居中-->
                        <div class="col-12  d-flex" >
                            <button type="button" class="btn  position-relative" @click="editPersionalInfomation">编辑个人信息</button>
                        </div>
                    </div>
                </div>
            </div>
            <div class="card" style="margin-top: 10px;">
                <div class="card-body">
                    <div class="row justify-content-center">
                        <div class="col-12">
                            <div class="introduction">个人简介</div>
                        </div>
                        <div class="col-12" style="margin-top: 5px;">
                            爱好：玩
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                            <br/>
                        </div>
                    </div>
                </div>
            </div>
        </div>

    </div>
</template>

<script>
import {useStore} from 'vuex'
import router from '@/router/index';   //@定位src目录
import { reactive } from 'vue';

export default {
    name:"UserProfileInfo",
    components:{},
    props:{
        userinfo: {
            type:Object,
            required: true,
        }
    },
    setup() {
        const store = useStore()
        const gender_map = reactive({
            0:'',
            1:'男',
            2:'女',
        })
        // 进入编辑个人信息页面
        const editPersionalInfomation = ()=>{
            router.push({name:"EditUserInfo"})
        }

        return {
            store,
            gender_map,
            editPersionalInfomation,
        }
    }
}


</script>

<style scoped>
button{
    width: 80%;
    margin: auto;
    background-color: #EFF9F2;
    color: #2DB55D;
    border: none;
}

button:hover {
    background-color: #92e4aa;
    color: white;
}

img {
    border-radius: 50%;
    /* width: 20px;
    height: 20px; */
}

.card{
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
}

.card-out {
    box-shadow: none;
    border: none;
}

.username {
    font-weight: bold;
}

.introduction{
    font-weight: bold;   
}

</style>