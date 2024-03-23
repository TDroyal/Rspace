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
                            <div v-if="userinfo.gender === 1">
                                <svg t="1709730743916" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1525" width="200" height="200"><path d="M511.843434 512m-446.708971 0a446.708971 446.708971 0 1 0 893.417942 0 446.708971 446.708971 0 1 0-893.417942 0Z" fill="#FFFFFF" p-id="1526"></path><path d="M424.111301 818.825061c-59.328223 0-115.140367-23.107277-157.101038-65.081251-86.620823-86.620823-86.620823-227.581252 0-314.215378 41.960671-41.973974 97.771791-65.081251 157.101038-65.081251 59.355853 0 115.140367 23.12058 157.101037 65.081251 41.960671 41.973974 65.081251 97.771791 65.081251 157.11434s-23.12058 115.140367-65.081251 157.101038-97.745185 65.081251-157.101037 65.081251z m0-360.620268c-36.97103 0-71.733765 14.409175-97.881285 40.543392-53.957913 53.984518-53.957913 141.804656 0 195.775872 26.14752 26.14752 60.910255 40.543392 97.881285 40.543391s71.733765-14.395872 97.881284-40.543391c26.14752-26.14752 40.543392-60.910255 40.543392-97.881285s-14.395872-71.733765-40.543392-97.894587c-26.146497-26.14752-60.909232-40.543392-97.881284-40.543392z" fill="#75B9EB" p-id="1527"></path><path d="M551.602973 511.016603c-10.715039 0-21.430078-4.090155-29.609365-12.269442-16.358573-16.358573-16.358573-42.874483 0-59.219753L672.577209 288.943808h-42.833551c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878806h143.958716c16.931624 0 32.200376 10.210549 38.689161 25.847691 6.488785 15.650445 2.889817 33.67189-9.078773 45.641503L581.212338 498.747161c-8.179286 8.179286-18.894326 12.269441-29.609365 12.269442z" fill="#75B9EB" p-id="1528"></path><path d="M773.703397 288.943808h-143.958716c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878806h143.958716c23.12058 0 41.878806 18.744923 41.878806 41.878806s-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1529"></path><path d="M779.864724 439.050548c-23.12058 0-41.878806-18.744923-41.878806-41.878806V253.226329c0-23.133883 18.758226-41.878806 41.878806-41.878807s41.878806 18.744923 41.878806 41.878807v143.945413c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1530"></path><path d="M779.864724 439.050548c-23.12058 0-41.878806-18.744923-41.878806-41.878806V253.226329c0-23.133883 18.758226-41.878806 41.878806-41.878807s41.878806 18.744923 41.878806 41.878807v143.945413c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1531"></path></svg>
                            </div>
                            <div v-else-if="userinfo.gender === 2">
                                <svg  t="1709731126558" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1669" width="200" height="200"><path d="M510.887666 512m-446.708971 0a446.708971 446.708971 0 1 0 893.417942 0 446.708971 446.708971 0 1 0-893.417942 0Z" fill="#FFFFFF" p-id="1670"></path><path d="M510.87948 578.902736c-123.673717 0-224.282113-100.607372-224.282113-224.282113s100.607372-224.282113 224.282113-224.282113 224.282113 100.607372 224.282112 224.282113-100.608396 224.282113-224.282112 224.282113z m0-364.80559c-77.486792 0-140.523477 63.036685-140.523477 140.523477s63.036685 140.523477 140.523477 140.523477 140.523477-63.036685 140.523477-140.523477-63.036685-140.523477-140.523477-140.523477z" fill="#FF3EC9" p-id="1671"></path><path d="M510.87948 896.635217c-23.12058 0-41.878806-18.744923-41.878806-41.878806V537.02393c0-23.133883 18.758226-41.878806 41.878806-41.878806s41.878806 18.744923 41.878806 41.878806v317.732481c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#FF3EC9" p-id="1672"></path><path d="M669.752884 737.762837H352.033705c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878807h317.719179c23.12058 0 41.878806 18.744923 41.878806 41.878807s-18.758226 41.878806-41.878806 41.878806z" fill="#FF3EC9" p-id="1673"></path></svg>
                            </div>
                            <!-- <div class="gender">{{gender_map[userinfo.gender]}}</div> -->
                        </div>
                    </div>
                    <div class="row justify-content-center" style="margin-top: 10px;">
                        <div class="col-5 follow-list">
                            <div class="text-center follow-title">关注了</div>
                            <div class="text-center">{{userinfo.followercount}}</div>
                        </div><span style="width: 10px; padding: 0px; color: gray; text-align: center;">|</span>
                        <div class="col-5 fans-list">
                            <div class="text-center follow-title">关注者</div>
                            <div class="text-center">{{userinfo.fanscount}}</div>
                        </div>
                    </div>
                    <!-- 点击头像进入编辑个人信息 -->
                    <div class="row justify-content-center" style="margin-top: 10px;">
                        <!-- d-flex  position-relative margin:auto; 就可以居中-->
                        <div class="col-12  d-flex" v-if="is_me === true">
                            <button type="button" class="btn  position-relative edit-userinfo" @click="editPersionalInfomation">编辑个人信息</button>
                        </div>
                        <div class="col-12  d-flex" v-if="is_me === false && userinfo.is_followed === false">
                            <button type="button" class="btn  position-relative follow" @click="follow">+ 关注</button>
                        </div>
                        <div class="col-12  d-flex" v-if="is_me === false && userinfo.is_followed === true">
                            <button type="button" class="btn  position-relative unfollow" @click="unfollow">取消关注</button>
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
                            <div style="margin-top: 5px;">
                                爱好：{{userinfo.introduction}}
                            </div>
                            <div style="margin-top: 5px;">
                                地址：{{userinfo.address}}
                            </div>
                            
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
import $ from 'jquery'
import BackendRootURL from '../common_resources/resource';

export default {
    name:"UserProfileInfo",
    components:{},
    props:{
        userinfo: {
            type:Object,
            required: true,
        },
        is_me: {
            type:Boolean,
            required:true,
        }
    },
    setup(props, context) {
        // console.log(props.is_me)
        // console.log(context)
        // console.log(props)
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

        const follow = () => {  //对该用户进行关注
            $.ajax({
                url: BackendRootURL + "/myspace/follow/",
                type: "POST",
                data: {
                    user_id: props.userinfo.id,
                },
                headers: {
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp) {
                    // console.log(resp)
                    if (resp.status === 0) {
                        context.emit('follow');  //调用父组件传过来的函数
                    }
                }
            });
        };

        const unfollow = () => {   //对该用户取消关注
            $.ajax({
                url: BackendRootURL + "/myspace/follow/",
                type: "POST",
                data: {
                    user_id: props.userinfo.id,
                },
                headers: {
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp) {
                    if (resp.status === 0) {
                        context.emit('unfollow');
                    }
                }
            });
        }


        return {
            store,
            gender_map,
            editPersionalInfomation,
            follow,
            unfollow,
        }
    }
}


</script>

<style scoped>
.edit-userinfo{
    width: 80%;
    margin: auto;
    background-color: #EFF9F2;
    color: #2DB55D;
    border: none;
}

.edit-userinfo:hover {
    background-color: #92e4aa;
    color: white;
}

.follow {
    width: 80%;
    margin: auto;
    background-color: #2DB55D;
    color: white;
    border: none;
}

.follow:hover {
    background-color: #269A4F;
    color: white;
}

.unfollow {
    width: 80%;
    margin: auto;
    background-color: #F2F3F4;
    color: #262626BF;
    border: none;
}

.unfollow:hover {
    background-color: #E5E6E8;
    color: #262626BF;
}


svg {
    max-height: 20px;
    max-width: 20px;
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
    margin-top: 5px;
    font-weight: bold;
}

.introduction{
    font-weight: bold;   
}

.follow-title{
    color: gray;
}

.follow-list {
    cursor: pointer;
}

.fans-list {
    cursor: pointer;
}

</style>