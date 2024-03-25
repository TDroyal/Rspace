<template>
    <div class="card card-out">
        <div class="card-body" style="background-color: #F7F8FA;">
            <div class="card">
                <div class="card-body">
                    <!-- <div class="row justify-content-center">
                        <div class="col-12">
                            <div class="star-title">
                                收藏
                            </div>
                            <div class="horizontal-line"></div>
                        </div>
                    </div>
                    <br/> -->

                    <!-- 只有当前查看的用户是自己的才显示取消以及取消符号 -->
                    <div class="row justify-content-center" v-if="is_me === true">
                        <div class="col-md-1 col-2 star-title-subtitle">
                            类型
                        </div>
                        <div class="col-md-4 col-2 star-title-subtitle">
                            内容
                        </div>
                        <div class="col-md-3 col-2 star-title-subtitle">
                            作者
                        </div>
                        <div class="col-md-3 col-4 star-title-subtitle">
                            发布日期
                        </div>
                        <div class="col-md-1 col-2 star-title-subtitle">
                            取消
                        </div>
                    </div>
                    <div class="row justify-content-center" v-if="is_me === false"> 
                        <div class="col-md-1 col-2 star-title-subtitle">
                            类型
                        </div>
                        <div class="col-md-4 col-2 star-title-subtitle">
                            内容
                        </div>
                        <div class="col-md-3 col-2 star-title-subtitle">
                            作者
                        </div>
                        <div class="col-md-4 col-6 star-title-subtitle">
                            发布日期
                        </div>
                    </div>
                    <div class="horizontal-line"></div>
                    <!-- 获取用户收藏的帖子，包含用户头像，id，name, 帖子id, 帖子content，类型，帖子发布日期createAt, -->
                    <div v-for="post in starposts_info.posts" :key="post.post_id">
                        <div class="card  post-list">
                            <div class="card-body" style="padding: 8px">
                                <div class="row" >
                                    <div class="col-md-1 col-1 post_type center-position">
                                        {{post_type_map[post.post_type]}}
                                    </div>
                                    <div class="col-md-4 col-3 content center-position" @click="enterPostDetail(post.post_id)">
                                        {{post.content}}
                                    </div>
                                    <div class="col-md-1 col-2 center-position" >
                                        <img class="img-fluid avatar"  :src="post.avatar" alt="" @click="enterUserProfile(post.user_id)">
                                    </div>
                                    <div class="col-md-2 col-1 username center-position" @click="enterUserProfile(post.user_id)">
                                        {{post.name}}
                                    </div>
                                    <div class="col-md-3 col-3 push_time center-position" v-if="is_me === true">
                                        {{post.push_time}}
                                    </div>
                                    <div class="col-md-1 col-2 unstar center-position" v-if="is_me === true" @click="changeCollectStatusForAPost(post.post_id)">
                                        <!-- <svg t="1711355455892" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2626" id="mx_n_1711355455893" width="200" height="200"><path d="M510.65 963C264.6 963 63.3 761.26 63.3 514.7S264.6 66.41 510.65 66.41 958 268.14 958 514.7 756.69 963 510.65 963zM655 858.61a377.28 377.28 0 0 0 198.65-199.12 371.82 371.82 0 0 0 0-289.58A377.18 377.18 0 0 0 655 170.79a369.15 369.15 0 0 0-288.64 0 377.09 377.09 0 0 0-198.71 199.12 371.82 371.82 0 0 0 0 289.58A377.12 377.12 0 0 0 366.33 858.6a369.21 369.21 0 0 0 288.64 0z" fill="#2c2c2c" p-id="2627"></path><path d="M287.806909 685.824011m26.516505-26.516504l349.31075-349.31075q26.516504-26.516504 53.033008 0l0 0q26.516504 26.516504 0 53.033008l-349.31075 349.31075q-26.516504 26.516504-53.033008 0l0 0q-26.516504-26.516504 0-53.033008Z" fill="#2c2c2c" p-id="2628"></path><path d="M686.881618 742.132969m-26.516505-26.516504l-349.31075-349.31075q-26.516504-26.516504 0-53.033008l0 0q26.516504-26.516504 53.033009 0l349.31075 349.310749q26.516504 26.516504 0 53.033009l0 0q-26.516504 26.516504-53.033009 0Z" fill="#2c2c2c" p-id="2629"></path></svg> -->
                                        <svg t="1711355596548" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2759" width="200" height="200"><path d="M510.65 963C264.6 963 63.3 761.26 63.3 514.7S264.6 66.41 510.65 66.41 958 268.14 958 514.7 756.69 963 510.65 963zM655 858.61a377.28 377.28 0 0 0 198.65-199.12 371.82 371.82 0 0 0 0-289.58A377.18 377.18 0 0 0 655 170.79a369.15 369.15 0 0 0-288.64 0 377.09 377.09 0 0 0-198.71 199.12 371.82 371.82 0 0 0 0 289.58A377.12 377.12 0 0 0 366.33 858.6a369.21 369.21 0 0 0 288.64 0z" fill="#515151" p-id="2760"></path><path d="M287.806909 685.824011m26.516505-26.516504l349.31075-349.31075q26.516504-26.516504 53.033008 0l0 0q26.516504 26.516504 0 53.033008l-349.31075 349.31075q-26.516504 26.516504-53.033008 0l0 0q-26.516504-26.516504 0-53.033008Z" fill="#515151" p-id="2761"></path><path d="M686.881618 742.132969m-26.516505-26.516504l-349.31075-349.31075q-26.516504-26.516504 0-53.033008l0 0q26.516504-26.516504 53.033009 0l349.31075 349.310749q26.516504 26.516504 0 53.033009l0 0q-26.516504 26.516504-53.033009 0Z" fill="#515151" p-id="2762"></path></svg>
                                    </div>
                                    <div class="col-md-4 col-4 push_time center-position" v-if="is_me === false">
                                        {{post.push_time}}
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="horizontal-line"></div>
                    </div>

                </div>
            </div>
        </div>
    </div>
</template>

<script>
import $ from 'jquery'
import {BackendRootURL, FrontendRootURL} from '../common_resources/resource';
import router from '@/router/index';   //@定位src目录 
import { ElMessage } from 'element-plus';
import { useStore } from 'vuex';
import { reactive } from 'vue';
import FormatDateTime from '../utils/DateTime'
export default {
    name: "UserProfileStar",
    components:{},
    props:{
        is_me:{
            type:Boolean,
            requried:true,
        },
    }, 
    setup() {
        const store = useStore()
        const post_type_map = reactive({
            '1':'日常',
            '2':'新鲜事',
            '3':'笔记',
            '4':'其它'
        })
        const starposts_info = reactive({
            count:0,
            posts:[],
        })
        $.ajax({
            url: BackendRootURL + "/myspace/getstarpostinfo/",
            type:"GET",
            data:{
                user_id: router.currentRoute.value.params.userid,
            },
            headers:{
                'Authorization': "Bearer " + store.state.user.jwt,
            },
            success(resp) {
                // console.log(resp)
                if(resp.status !== 0) {
                    ElMessage.error("获取收藏列表失败，请稍后重试")
                    return 
                }
                if(resp.data === null) {
                    return 
                }
                // 
                starposts_info.count = resp.data.length
                starposts_info.posts = resp.data
                for(let i = 0; i < starposts_info.count; i ++ ) {
                    starposts_info.posts[i].avatar = BackendRootURL + "/static/avatar/" + starposts_info.posts[i].avatar
                    starposts_info.posts[i].push_time = FormatDateTime(starposts_info.posts[i].push_time)
                }
            },
            error() {
                ElMessage.error("获取收藏列表失败，请稍后重试")
            }
        })

        // 进入用户的个人空间 (后续改为用新窗口打开)
        const enterUserProfile = (user_id)=>{
            if(store.state.user.is_login) {
                // context.emit('changeNavbar', '0')
                
                // router.push({
                //     name:"Profile",
                //     params:{
                //         userid:user_id,
                //     }
                // })
                window.open(FrontendRootURL + `/profile/${user_id}`, '_blank');  //新窗口打开
            }else {
                router.push({
                    name:"Login",
                })
            }
            
        }

        // 进入帖子详情页面 (后续改为用新窗口打开)
        const enterPostDetail = (post_id)=> {
            // router.push({
            //     name:"PostDetail",
            //     params:{
            //         postid:post_id
            //     },
            // })
            window.open(`/post/${post_id}`, '_blank');  //新窗口打开
        }

        // 取消收藏，先从starposts_info列表数组中删除根据post_id，然后再后端改status值=0
        const changeCollectStatusForAPost = (post_id) =>{
            starposts_info.posts = starposts_info.posts.filter(post => post.post_id !== post_id);  //前端先删除收藏的帖子
            starposts_info.count = starposts_info.posts.length;
            
            $.ajax({
                url: BackendRootURL + "/post/changecollectStatus/",
                type:"POST",
                data:{
                    post_id: post_id,
                    user_id: store.state.user.id,
                    status: 0, //0表示未收藏
                },
                headers:{
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp) {
                    if(resp.status !== 0) {
                        ElMessage.error("取消收藏失败")
                        return 
                    }
                },
                error(resp) {
                    ElMessage({
                        message: '取消收藏失败，请稍后重试',
                        type: 'error',
                    })
                    console.log(resp)
                }
            })
        }

        return {
            post_type_map,
            starposts_info,
            enterUserProfile,
            enterPostDetail,
            changeCollectStatusForAPost,
        }
    }
}

</script>


<style scoped>
.card{
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
}

.card-out {
    box-shadow: none;
    border: none;
}

.post-list {
    box-shadow: none;
    border: none;
}
.star-title{
    font-weight: bold;
}

.horizontal-line {
  border-bottom: 1px solid lightgray;
}

.star-title-subtitle {
    font-weight: bold;
}

.avatar {
    border-radius: 50%;
    
}

.avatar:hover {
    cursor: pointer;
}

.username {
    padding-left: 0px;
    color: #337AB7;
}

.username:hover {
    cursor: pointer;
}

.center-position {
    display: flex;
    text-align: left;
    align-items: center;
}

.content {
    padding-left: 0px;
    color: #337AB7;
}

.content:hover {
    cursor: pointer;
}

.post_type {
    padding-left: 0px;
}

svg {
    max-height: 30px;
    max-width: 30px;
}

.unstar {
    cursor: pointer;
}

</style>