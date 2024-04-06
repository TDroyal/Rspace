<template>
    <div class="card card-out" style="background-color:transparent; z-index: 3;">
        <div class="card-body" >
            <div class="card">
                <div class="card-body">
                    <div class="row justify-content-center">
                        <div class="col-12">
                            <div class="follow-title" >关注列表</div>
                            <div class="horizontal-line"></div>
                        </div>
                    </div>
                    <br/>
                    <div v-for="user in followerList.users" :key="user.id">
                        <div class="card  user-list" @click="enterUserProfile(user.id)">
                            <div class="card-body" style="padding: 8px">
                                <div class="row">
                                    <div class="col-md-1 col-3">
                                        <img class="img-fluid avatar"  :src="user.avatar" alt="">
                                    </div>
                                    <div class="col-md-9 col-9">
                                        <div class="username">
                                            {{user.name}}
                                        </div>
                                        <div v-if="user.gender === 1">
                                            <svg t="1709730743916" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1525" width="200" height="200"><path d="M511.843434 512m-446.708971 0a446.708971 446.708971 0 1 0 893.417942 0 446.708971 446.708971 0 1 0-893.417942 0Z" fill="#FFFFFF" p-id="1526"></path><path d="M424.111301 818.825061c-59.328223 0-115.140367-23.107277-157.101038-65.081251-86.620823-86.620823-86.620823-227.581252 0-314.215378 41.960671-41.973974 97.771791-65.081251 157.101038-65.081251 59.355853 0 115.140367 23.12058 157.101037 65.081251 41.960671 41.973974 65.081251 97.771791 65.081251 157.11434s-23.12058 115.140367-65.081251 157.101038-97.745185 65.081251-157.101037 65.081251z m0-360.620268c-36.97103 0-71.733765 14.409175-97.881285 40.543392-53.957913 53.984518-53.957913 141.804656 0 195.775872 26.14752 26.14752 60.910255 40.543392 97.881285 40.543391s71.733765-14.395872 97.881284-40.543391c26.14752-26.14752 40.543392-60.910255 40.543392-97.881285s-14.395872-71.733765-40.543392-97.894587c-26.146497-26.14752-60.909232-40.543392-97.881284-40.543392z" fill="#75B9EB" p-id="1527"></path><path d="M551.602973 511.016603c-10.715039 0-21.430078-4.090155-29.609365-12.269442-16.358573-16.358573-16.358573-42.874483 0-59.219753L672.577209 288.943808h-42.833551c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878806h143.958716c16.931624 0 32.200376 10.210549 38.689161 25.847691 6.488785 15.650445 2.889817 33.67189-9.078773 45.641503L581.212338 498.747161c-8.179286 8.179286-18.894326 12.269441-29.609365 12.269442z" fill="#75B9EB" p-id="1528"></path><path d="M773.703397 288.943808h-143.958716c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878806h143.958716c23.12058 0 41.878806 18.744923 41.878806 41.878806s-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1529"></path><path d="M779.864724 439.050548c-23.12058 0-41.878806-18.744923-41.878806-41.878806V253.226329c0-23.133883 18.758226-41.878806 41.878806-41.878807s41.878806 18.744923 41.878806 41.878807v143.945413c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1530"></path><path d="M779.864724 439.050548c-23.12058 0-41.878806-18.744923-41.878806-41.878806V253.226329c0-23.133883 18.758226-41.878806 41.878806-41.878807s41.878806 18.744923 41.878806 41.878807v143.945413c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1531"></path></svg>
                                        </div>
                                        <div v-else-if="user.gender === 2">
                                            <svg  t="1709731126558" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1669" width="200" height="200"><path d="M510.887666 512m-446.708971 0a446.708971 446.708971 0 1 0 893.417942 0 446.708971 446.708971 0 1 0-893.417942 0Z" fill="#FFFFFF" p-id="1670"></path><path d="M510.87948 578.902736c-123.673717 0-224.282113-100.607372-224.282113-224.282113s100.607372-224.282113 224.282113-224.282113 224.282113 100.607372 224.282112 224.282113-100.608396 224.282113-224.282112 224.282113z m0-364.80559c-77.486792 0-140.523477 63.036685-140.523477 140.523477s63.036685 140.523477 140.523477 140.523477 140.523477-63.036685 140.523477-140.523477-63.036685-140.523477-140.523477-140.523477z" fill="#FF3EC9" p-id="1671"></path><path d="M510.87948 896.635217c-23.12058 0-41.878806-18.744923-41.878806-41.878806V537.02393c0-23.133883 18.758226-41.878806 41.878806-41.878806s41.878806 18.744923 41.878806 41.878806v317.732481c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#FF3EC9" p-id="1672"></path><path d="M669.752884 737.762837H352.033705c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878807h317.719179c23.12058 0 41.878806 18.744923 41.878806 41.878807s-18.758226 41.878806-41.878806 41.878806z" fill="#FF3EC9" p-id="1673"></path></svg>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <!-- <a :href="'/profile/' + user.id + '/'" target="_black">
                            <div class="card  user-list">
                                <div class="card-body" style="padding: 8px">
                                    <div class="row">
                                        <div class="col-md-1 col-3">
                                            <img class="img-fluid avatar"  :src="user.avatar" alt="">
                                        </div>
                                        <div class="col-md-9 col-9">
                                            <div class="username">
                                                {{user.name}}
                                            </div>
                                            <div v-if="user.gender === 1">
                                                <svg t="1709730743916" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1525" width="200" height="200"><path d="M511.843434 512m-446.708971 0a446.708971 446.708971 0 1 0 893.417942 0 446.708971 446.708971 0 1 0-893.417942 0Z" fill="#FFFFFF" p-id="1526"></path><path d="M424.111301 818.825061c-59.328223 0-115.140367-23.107277-157.101038-65.081251-86.620823-86.620823-86.620823-227.581252 0-314.215378 41.960671-41.973974 97.771791-65.081251 157.101038-65.081251 59.355853 0 115.140367 23.12058 157.101037 65.081251 41.960671 41.973974 65.081251 97.771791 65.081251 157.11434s-23.12058 115.140367-65.081251 157.101038-97.745185 65.081251-157.101037 65.081251z m0-360.620268c-36.97103 0-71.733765 14.409175-97.881285 40.543392-53.957913 53.984518-53.957913 141.804656 0 195.775872 26.14752 26.14752 60.910255 40.543392 97.881285 40.543391s71.733765-14.395872 97.881284-40.543391c26.14752-26.14752 40.543392-60.910255 40.543392-97.881285s-14.395872-71.733765-40.543392-97.894587c-26.146497-26.14752-60.909232-40.543392-97.881284-40.543392z" fill="#75B9EB" p-id="1527"></path><path d="M551.602973 511.016603c-10.715039 0-21.430078-4.090155-29.609365-12.269442-16.358573-16.358573-16.358573-42.874483 0-59.219753L672.577209 288.943808h-42.833551c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878806h143.958716c16.931624 0 32.200376 10.210549 38.689161 25.847691 6.488785 15.650445 2.889817 33.67189-9.078773 45.641503L581.212338 498.747161c-8.179286 8.179286-18.894326 12.269441-29.609365 12.269442z" fill="#75B9EB" p-id="1528"></path><path d="M773.703397 288.943808h-143.958716c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878806h143.958716c23.12058 0 41.878806 18.744923 41.878806 41.878806s-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1529"></path><path d="M779.864724 439.050548c-23.12058 0-41.878806-18.744923-41.878806-41.878806V253.226329c0-23.133883 18.758226-41.878806 41.878806-41.878807s41.878806 18.744923 41.878806 41.878807v143.945413c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1530"></path><path d="M779.864724 439.050548c-23.12058 0-41.878806-18.744923-41.878806-41.878806V253.226329c0-23.133883 18.758226-41.878806 41.878806-41.878807s41.878806 18.744923 41.878806 41.878807v143.945413c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#75B9EB" p-id="1531"></path></svg>
                                            </div>
                                            <div v-else-if="user.gender === 2">
                                                <svg  t="1709731126558" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="1669" width="200" height="200"><path d="M510.887666 512m-446.708971 0a446.708971 446.708971 0 1 0 893.417942 0 446.708971 446.708971 0 1 0-893.417942 0Z" fill="#FFFFFF" p-id="1670"></path><path d="M510.87948 578.902736c-123.673717 0-224.282113-100.607372-224.282113-224.282113s100.607372-224.282113 224.282113-224.282113 224.282113 100.607372 224.282112 224.282113-100.608396 224.282113-224.282112 224.282113z m0-364.80559c-77.486792 0-140.523477 63.036685-140.523477 140.523477s63.036685 140.523477 140.523477 140.523477 140.523477-63.036685 140.523477-140.523477-63.036685-140.523477-140.523477-140.523477z" fill="#FF3EC9" p-id="1671"></path><path d="M510.87948 896.635217c-23.12058 0-41.878806-18.744923-41.878806-41.878806V537.02393c0-23.133883 18.758226-41.878806 41.878806-41.878806s41.878806 18.744923 41.878806 41.878806v317.732481c0 23.133883-18.758226 41.878806-41.878806 41.878806z" fill="#FF3EC9" p-id="1672"></path><path d="M669.752884 737.762837H352.033705c-23.12058 0-41.878806-18.744923-41.878806-41.878806s18.758226-41.878806 41.878806-41.878807h317.719179c23.12058 0 41.878806 18.744923 41.878806 41.878807s-18.758226 41.878806-41.878806 41.878806z" fill="#FF3EC9" p-id="1673"></path></svg>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </a> -->
                        <div class="horizontal-line"></div>
                    </div>
                    <el-pagination hide-on-single-page  style="justify-content: right; " :page-size="page_size" v-model:current-page="current_page" large background layout="prev, pager, next" :total="followerList.total_count" class="mt-4" @change="changePage"/>
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
import { ref, reactive } from 'vue';
export default {
    name: "UserProfileFollowList",
    components:{},
    setup() {        
        const store = useStore()
        const followerList = reactive({
            total_count:0,
            count:0,
            users:[],
        })
        const page_size = ref(10)
        const current_page = ref(1)
        //在这里请求关注用户的id, 头像avatar，名称name，性别gender     /myspace/getfollowersinfo/
        const get_follower_info = ()=>{
            $.ajax({
                url: BackendRootURL + "/myspace/getfollowersinfo/",
                type:"GET",
                data: {
                    user_id: router.currentRoute.value.params.userid,
                    page_size: page_size.value,
                    current_page: current_page.value,
                    search_type: 1,
                },
                headers:{
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp) {
                    // console.log(resp)
                    if(resp.status !== 0) {
                        ElMessage.error("获取关注列表失败")
                        return 
                    }
                    if (resp.data === null || resp.data.data === null) {
                        return
                    }
                    followerList.count = resp.data.data.length
                    followerList.users = resp.data.data
                    followerList.total_count = resp.data.count
                    for(let i = 0; i < followerList.count; i ++ ) {
                        followerList.users[i].avatar = BackendRootURL + "/static/avatar/" + followerList.users[i].avatar
                    }
                    
                    // console.log(followerList)
                },
                error() {
                    // console.log(resp)
                    ElMessage.error("获取关注列表失败")
                }
            })
        }
        get_follower_info()
        // 进入用户的个人空间
        const enterUserProfile = (user_id)=>{
            if(store.state.user.is_login) {
                // context.emit('changeNavbar', '0')
                window.open(FrontendRootURL + `/profile/${user_id}/`, '_blank');
                // router.push({
                //     name:"Profile",
                //     params:{
                //         userid:user_id,
                //     }
                // })
            }else {
                router.push({
                    name:"Login",
                })
            }
            
        }

        const changePage = ()=>{
            get_follower_info()
        }        

        return {
            followerList,
            enterUserProfile,
            get_follower_info,

            page_size,
            current_page,
            changePage,
        }
    }
}

</script>


<style scoped>
/* .card{
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
} */

.card-out {
    box-shadow: none;
    border: none;
}

.user-list {
    box-shadow: none;
    border: none;
}

.follow-title{
    font-weight: bold;
}

.horizontal-line {
  border-bottom: 1px solid lightgray;
}

.avatar {
    border-radius: 50%;
    
}

.username {
    font-weight: bold;
}

.user-list:hover{
    scale: calc(1.01);
    cursor: pointer;
    transition: 100ms;
}

svg {
    max-height: 20px;
    max-width: 20px;
}

a {
    text-decoration: none;
}

</style>