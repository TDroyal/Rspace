<!-- 消息中心 -->
<template>
    <Content>
        <div class="row justify-content-center" >
            <div class="col-md-9 col-12">
                <div class="card card-out">
                    <div class="card-body">
                        <div class="row">
                            <div class="col-12">
                                <div style="font-weight: bold; font-size: 20px;">
                                    消息中心
                                    
                                </div>
                            </div>
                            <div class="horizontal-line"></div>
                            <div v-for="(notice, ) in notifications.messages" :key="notice.id">
                                <div class="card  card-sigle">
                                    <div class="card-body">
                                        <div class="row">
                                            <div class="col-md-1 col-3" >
                                                <img class="img-fluid avatar" :src="notice.avatar" alt="" @click="enterUserProfile(notice.user_id, notice.id, notice.notification_status)">
                                            </div>
                                            
                                            <div class="col-md-10 col-6 notification-message" v-if="notice.notification_type !== 3">
                                                <p class="username" @click="enterUserProfile(notice.user_id, notice.id, notice.notification_status)"> {{notice.name}}</p>
                                                <p >在</p>
                                                <p class="post_content" @click="enterPostDetail(notice.post_id, notice.id, notice.notification_status)">{{notice.post_content}}</p>
                                                <p v-if="notice.notification_type === 1" >中评论了你</p>
                                                <p v-if="notice.notification_type === 2" >中赞了你</p>
                                                
                                            </div>
                                            <!-- 表示是关注通知(与帖子无关的通知) -->
                                            <div class="col-md-10 col-6 notification-message"  v-if="notice.notification_type === 3">
                                                <p class="username" @click="enterUserProfile(notice.user_id, notice.id, notice.notification_status)"> {{notice.name}}</p>
                                                <p>关注了你</p>
                                                
                                            </div>

                                            <div class="col-md-1 col-3" style="display: flex; align-items: center; padding: 0px;">
                                                <div class="notification-time" style="color: gray;">
                                                    {{notice.created_at}}
                                                </div>
                                                <div v-if="notice.notification_status === 1" style="padding-left: 5px;">
                                                    <el-badge is-dot class="item"></el-badge>    
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="horizontal-line"></div>
                                </div>
                            </div>
                            <el-pagination hide-on-single-page  style="justify-content: right; " :page-size="page_size" v-model:current-page="current_page" large background layout="prev, pager, next" :total="notifications.total_count" class="mt-4" @change="changePage" :pager-count="5"/>
                            <div v-if="notifications.total_count === 0">
                                <el-empty description="没有数据" />
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </Content>
</template>

<script>
import Content from '../components/Content.vue'
import router from '@/router/index';   //@定位src目录
import { useStore} from 'vuex'
import { ref, reactive } from 'vue';
import {  BackendRootURL } from '../common_resources/resource';
import $ from 'jquery'
import { ElMessage } from 'element-plus';
import {FormatDateTime, GetTimePeriod} from '../utils/DateTime'

export default {
    name:"Notifications",
    components:{Content},
    setup() {
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

        const notifications = reactive({
            total_count:0,
            count: 0,
            messages:[],
            // total_count:1,
            // count: 1,
            // messages: [
            //     {
            //         id:1,
            //         user_id:2,
            //         username: '天生',  //是用户名，不是账号名
            //         avatar: BackendRootURL + '/static/avatar/' + 'QQ图片20220704142508.jpg',
            //         notification_status:1, //默认是1， 1是此消息未读
            //         notification_type: 2,   //2是点赞
            //         CreatedAt: "2天前",
            //         post_id:48,
            //         post_content:'测试',
            //     },
            // ],
        })

        const page_size = ref(10)
        const current_page = ref(1)

        const get_notifications = ()=>{
            $.ajax({
                url: BackendRootURL +  "/notification/get/",
                type:"GET",
                data:{   //传入page_size和current_page
                    page_size: page_size.value,
                    current_page: current_page.value,
                },
                headers:{
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp) {
                    // console.log(resp)
                    if(resp.status !== 0) {
                        ElMessage.error(resp.message)
                        return 
                    }
                    // console.log(resp)
                    notifications.messages = resp.data.data
                    notifications.count = resp.data.data.length
                    notifications.total_count = resp.data.count

                    for(let i = 0; i < notifications.count; i ++ )
                    {
                        notifications.messages[i].avatar = BackendRootURL + '/static/avatar/' +  notifications.messages[i].avatar
                        notifications.messages[i].created_at = GetTimePeriod(FormatDateTime(notifications.messages[i].created_at))
                    }

                },error(){
                    
                }
            })
        }   
        get_notifications()
        // console.log(notifications)

        // 将通知设为已读
        const changeNotificationStatus = (notice_id, notification_status_)=>{
            if(notification_status_ === 0) {
                return 
            }
            $.ajax({
                url: BackendRootURL + "/notification/changenotificationstatus/",
                type:"POST",
                data: {
                    notice_id: notice_id,
                    notification_status: 0,  //设为已读
                },
                headers:{
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp){
                    //表示此消息已读，然后将总的未读消息减一即可(存入全局store里面)
                    if(resp.status === 0) {
                        store.commit("updateUnreadNotificationCount", store.state.user.unread_notification_count - 1)
                    }
                }
                
                
                
            })
        }

        // 进入用户的个人空间
        const enterUserProfile = (user_id, notice_id, notification_status)=>{
            if(store.state.user.is_login) {
                // window.open(FrontendRootURL + `/profile/${user_id}/`, '_blank');
                router.push({
                    name:"Profile",
                    params:{
                        userid:user_id,
                    }
                })
                //点击后把此通知的状态设为已读
                changeNotificationStatus(notice_id, notification_status)

            }else {
                router.push({
                    name:"Login",
                })
            }
        }

        // 进入帖子详情页面 (后续改为用新窗口打开)
        const enterPostDetail = (post_id, notice_id, notification_status)=> {
            // window.open(`/post/${post_id}/`, '_blank');  //新窗口打开
            router.push({
                name:"PostDetail",
                params:{
                    postid:post_id
                },
            })

            //点击后把此通知的状态设为已读
            changeNotificationStatus(notice_id, notification_status)

        }

        const changePage = ()=>{
            get_notifications()
        }

        //通知消息数量(刷新一次获取一次，然后点击一次之前未查看消息，就减一)
        // 获取未读取的信息数
        const getUnreadNotificationsCount = ()=>{
            $.ajax({
                url: BackendRootURL + "/notification/getunreadnotificationcount/",
                type:"GET",
                headers: {
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp){
                    // notifications_count.value = resp.data
                    store.commit("updateUnreadNotificationCount", resp.data)
                }
            })
        }

        // 未登录不能获取未读通知数量
        if(store.state.user.is_login === true) {
            getUnreadNotificationsCount()
        }

        return {
            page_size,
            current_page,
            changePage,
            notifications,
            get_notifications,
            enterUserProfile,
            enterPostDetail,
        }
    }
}

</script>


<style scoped>

.avatar {
    border-radius: 50%;
    
}

.card-out {
    margin-top: 20px;
    border: none;
}

.horizontal-line {
  border-bottom: 1px solid lightgray;
}

.card-sigle{
    border: none;
}

.card-sigle:hover {
    background-color: #F7F7F7;
    cursor: pointer;   
}

.username {
    color: #007AFF;
    /* font-weight: bold; */
}

.post_content {
    color: #007AFF;
}

.notification-time{
    text-align: center;
}

.notification-message {
    padding-left: 0px;
}

.notification-message p {
    /* 水平排列 连续显示，而不会换行。 */
    display: inline;
}

.item {

  margin-top: 10px;
  margin-left: 4px;
}

</style>