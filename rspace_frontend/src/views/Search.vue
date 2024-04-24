<template>
    <Content v-infinite-scroll="load"  :infinite-scroll-disabled="disabled">
        <div class="row justify-content-center" >
            <div class="col-md-9 col-12">
                <div class="card card-out">
                    <div class="card-body">
                        <el-menu
                            :default-active="activeIndex"
                            class="el-menu-demo"
                            mode="horizontal"
                            text-color="#3C3C4399"
                            active-text-color="#262626"
                            @select="handleSelect"
                        >
                            <el-menu-item index="2">新鲜事</el-menu-item>
                            <el-menu-item index="1">日常</el-menu-item>
                            <el-menu-item index="3">笔记</el-menu-item>
                            <el-menu-item index="5">用户</el-menu-item>
                            <el-menu-item index="4">其它</el-menu-item>
                        </el-menu>

                        <!-- 下面渲染的是 -->
                        <div v-if="route.query.query_type !== 'user'">
                            <div v-for="post in posts.posts" :key="post.post_id">
                                <div class="post-list" style="padding-top: 10px;" @click="enterPostDetail(post.post_id)">
                                    <div class="row" >
                                        <div class="col-md-1 col-3">
                                            <img class="img-fluid avatar" :src="post.avatar" alt="">
                                        </div>
                                        <div class="col-md-5 col-8">
                                            <div class="row">
                                                <div class="col-12 name">
                                                    {{post.name}}
                                                </div>
                                                <div class="col-12 publish-time">
                                                    {{post.created_at}}
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="row" style="margin-top: 5px;">
                                        <div class="col-12">
                                            {{post.content}}
                                        </div>
                                    </div> 
                                    
                                    <div class="row justify-content-center" style="padding-top: 10px; ">
                                        <div class="d-flex flex-wrap">
                                            <div class="col-md-2 col-4" v-for="(image) in post.image" :key="image.id">
                                                <img class="img-thumbnail" :src="image.image" alt="">
                                            </div>
                                        </div>
                                    </div>

                                    <div class="row">
                                        <div class="col-12">
                                            <svg t="1710466679468" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5759" width="200" height="200"><path d="M512 901.746939c-13.583673 0-26.122449-4.179592-37.093878-13.061225-8.881633-7.314286-225.697959-175.020408-312.424489-311.379592C133.746939 532.37551 94.040816 471.24898 94.040816 384.522449c0-144.718367 108.146939-262.269388 240.326531-262.269388 67.395918 0 131.657143 30.82449 177.632653 84.636735 45.453061-54.334694 109.191837-84.636735 177.110204-84.636735 132.702041 0 240.326531 117.55102 240.326531 262.269388 0 85.159184-37.093878 143.673469-67.395919 191.216327l-1.044898 1.567346c-86.726531 136.359184-303.542857 304.587755-312.424489 311.379592-10.44898 8.359184-22.987755 13.061224-36.571429 13.061225z" fill="#8a8a8a" p-id="5760" data-spm-anchor-id="a313x.search_index.0.i15.73333a81m6RRCQ" class="selected"></path></svg>
                                            <span style="padding-right: 7px;">{{post.collection_count}}</span>
                                            <svg t="1710466380019" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4748" width="200" height="200"><path d="M853.333333 768c35.413333 0 64-20.650667 64-55.978667V170.581333A63.978667 63.978667 0 0 0 853.333333 106.666667H170.666667C135.253333 106.666667 106.666667 135.253333 106.666667 170.581333v541.44C106.666667 747.285333 135.338667 768 170.666667 768h201.173333l110.016 117.44a42.752 42.752 0 0 0 60.586667 0.042667L651.904 768H853.333333z m-219.029333-42.666667h-6.250667l-115.797333 129.962667c-0.106667 0.106667-116.010667-129.962667-116.010667-129.962667H170.666667c-11.776 0-21.333333-1.621333-21.333334-13.312V170.581333A21.205333 21.205333 0 0 1 170.666667 149.333333h682.666666c11.776 0 21.333333 9.536 21.333334 21.248v541.44c0 11.754667-9.472 13.312-21.333334 13.312H634.304zM341.333333 490.666667a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z m170.666667 0a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z m170.666667 0a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z" fill="#3D3D3D" p-id="4749"></path></svg>
                                            <span >{{post.comment_count}}</span>
                                        </div>
                                    </div>

                                    <div class="horizontal-line" style="margin-top: 10px;"></div>
                                </div>
                            </div>
                        </div>
                        <div v-else>
                            <div v-for="user in users.users" :key="user.user_id">
                                <div class="post-list" style="padding-top: 10px;" @click="enterProfile(user.user_id)">
                                    <div class="row" >
                                        <div class="col-md-1 col-3">
                                            <img class="img-fluid avatar" :src="user.avatar" alt="">
                                        </div>
                                        <div class="col-md-5 col-8">
                                            <div class="row">
                                                <div class="col-12 name">
                                                    {{user.name}}
                                                </div>
                                                <div class="col-12 publish-time">
                                                    粉丝数：{{user.fanscount}}
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                    <div class="horizontal-line" style="margin-top: 10px;"></div>
                                </div>
                            </div>
                        </div>
                        <!-- <p class="load-status" v-if="loading">加载中...</p> -->
                        <div class="load-status" v-if="loading">
                            <div class="d-flex justify-content-center">
                                <div class="spinner-border" role="status">
                                    <span class="visually-hidden">Loading...</span>
                                </div>
                            </div>
                        </div>
                        <p class="load-status" v-if="noMore">没有更多搜索结果</p>
                    </div>
                </div>
                <CopyRight></CopyRight>
            </div>
        </div>
    </Content>
</template>

<script>
import Content from '../components/Content.vue'
import {useRoute} from 'vue-router'
import {ref, reactive, watch, computed} from 'vue'
import router from '@/router/index'
import $ from 'jquery'
import { BackendRootURL } from '../common_resources/resource'
import { ElMessage,} from 'element-plus'
import {FormatDateTime, GetTimePeriod} from '../utils/DateTime'
import ParseImageUrl from '../utils/ParseImageUrl'
import CopyRight from '../components/CopyRight.vue'
export default {
    name:"Search",
    components:{Content, CopyRight},
    setup()
    {
        // 假设 URL 为 /search/?query_content=royal_111      ? 后面的内容被解析为查询参数。
        const route = useRoute()
        // console.log(route.query.query_type, route.query.query_content)   //royal_111
        //就是根据搜索内容去后端模糊匹配

        //     '1':'日常',  daily
            //     '2':'新鲜事', news
            //     '3':'笔记',   note
            //     '4':'其它'    other
            //          用户     user        

        const search_type_map = reactive({
            'news':'2',
            'daily':'1',
            'note':'3',
            'user':'5',
            'other':'4',
        })
        
        //（除了用户）其它均返回用户名称，用户头像，帖子发布时间，图片和内容，点赞量和评论量（并且点击每一个搜到的帖子或者用户都新窗口打开）
        // 用户的话，只返回用户头像，名称，粉丝数量
        //封装两个请求数据的函数，一个是请求帖子，一个是请求用户

        const page_size = ref(5)  //每次请求page_size个
        const current_page = ref(1)

        // 总的帖子列表
        const posts = reactive({
            count: 0,
            posts:[],
        })

        //单次的帖子列表
        const get_posts = reactive({
            count: 0,
            posts:[],
        })

        // 总的用户列表
        const users = reactive({
            count: 0,
            users:[],
        })

        //单次的用户列表
        const get_users = reactive({
            count: 0,
            users:[],
        })

        const loading = ref(false)
        let noMore = computed(() => get_posts.count < page_size.value)
        const disabled = computed(() => loading.value || noMore.value)

        const getPostList = (post_type)=>{
            // post_type = parseInt(post_type) 
            // console.log("获取帖子数据", post_type)


            // console.log("请求页数：", current_page.value)
            $.ajax({
                url: BackendRootURL +  "/search/getpost/",
                type:"GET",
                data:{
                    page_size:page_size.value,
                    current_page:current_page.value,
                    post_type:parseInt(post_type),
                    query_content:route.query.query_content,
                },
                success(resp){
                    if(resp.status !== 0) {
                        ElMessage.error("搜索失败，请稍后重试！")
                        return 
                    }
                    if(resp.data !== null) {
                        get_posts.count = resp.data.length
                        get_posts.posts = resp.data
                        posts.count += get_posts.count
                        for(let i = 0; i < get_posts.count; i ++ ) {
                            get_posts.posts[i].image = ParseImageUrl(get_posts.posts[i].image)
                            get_posts.posts[i].created_at = GetTimePeriod(FormatDateTime(get_posts.posts[i].created_at))
                            get_posts.posts[i].avatar = BackendRootURL + "/static/avatar/" + get_posts.posts[i].avatar
                            posts.posts.push(get_posts.posts[i])
                        }
                        loading.value = false
                        current_page.value += 1
                    }else {
                        get_posts.count = 0
                        get_posts.posts =[]
                    }
                    noMore = computed(() => get_posts.count < page_size.value)
                    if(noMore.value === true) {
                        loading.value = false
                    }
                },
                error(){
                    ElMessage.error("搜索失败，请稍后重试！")
                    return 
                }
            })
        }

        const getUserList = ()=>{
            // console.log("获取用户数据")
            // console.log("请求页数：", current_page.value)
            $.ajax({
                url: BackendRootURL +  "/search/getuser/",
                type:"GET",
                data:{
                    page_size:page_size.value,
                    current_page:current_page.value,
                    query_content:route.query.query_content,
                },
                success(resp){
                    if(resp.status !== 0) {
                        ElMessage.error("搜索失败，请稍后重试！")
                        return 
                    }
                    // console.log(resp)
                    if(resp.data !== null) {
                        get_users.count = resp.data.length
                        get_users.users = resp.data
                        users.count += get_users.count

                        for(let i = 0; i < get_users.count; i ++ ) {
                            get_users.users[i].avatar = BackendRootURL + "/static/avatar/" + get_users.users[i].avatar
                            users.users.push(get_users.users[i])
                        }
                        // console.log(users)

                        loading.value = false
                        current_page.value += 1
                    }else {
                        get_users.count = 0
                        get_users.users =[]
                    }
                    noMore = computed(() => get_users.count < page_size.value)
                    if(noMore.value === true) {
                        loading.value = false
                    }
                },
                error(){
                    ElMessage.error("搜索失败，请稍后重试！")
                    return 
                }
            })
        }

        //使用watch监听路由参数是否变化，变化了就重新加载数据
        watch(() => route.query.query_type, () => {
            // loadData();
            // console.log(route.query.query_type, route.query.query_content)
            // 切换搜索类型后，将posts和users清空
            
            posts.count = 0
            posts.posts = []
            get_posts.count = 0
            get_posts.posts = []

            users.count = 0
            users.users = []
            get_users.count = 0
            get_users.users = []

            loading.value = false
            current_page.value = 1
            if(route.query.query_type === 'user') {
                getUserList()
            }else {
                // console.log("watch")
                getPostList(search_type_map[route.query.query_type])
            }

        }, { immediate: true }); // immediate: true 表示在组件初始化时立即执行

        const activeIndex = ref(search_type_map[route.query.query_type])  //应该根据url的搜索类型来确定，避免刷新跳回默认页面
        const handleSelect = (key)=>{
            if(key === '1') {
                router.push({
                    name:"Search",
                    query:{
                        query_type:"daily",
                        query_content:route.query.query_content,
                    },
                })
            }else if(key == '2') {
                router.push({
                    name:"Search",
                    query:{
                        query_type:"news",
                        query_content:route.query.query_content,
                    },
                })
            }else if(key == '3') {
                router.push({
                    name:"Search",
                    query:{
                        query_type:"note",
                        query_content:route.query.query_content,
                    },
                })
            }else if(key == '4') {
                router.push({
                    name:"Search",
                    query:{
                        query_type:"other",
                        query_content:route.query.query_content,
                    },
                })
                
            }else if(key == '5') {
                router.push({
                    name:"Search",
                    query:{
                        query_type:"user",
                        query_content:route.query.query_content,
                    },
                })
            }
        }

        // 进入帖子详情页面
        const enterPostDetail = (post_id)=> {
            // router.push({
            //     name:"PostDetail",
            //     params:{
            //         postid:post_id
            //     },
            // })
            window.open(`/post/${post_id}/`, '_blank');  //新窗口打开
        }

        const enterProfile = (user_id)=> {
            // router.push({
            //     name:"Profile",
            //     params:{
            //         userid:user_id,
            //     }
            // })
            window.open(`/Profile/${user_id}/`, '_blank');  //新窗口打开
        }

        const load = ()=>{
            loading.value = true
            if(route.query.query_type === 'user') {
                getUserList()
            }else {
                // console.log("load", current_page.value)
                getPostList(search_type_map[route.query.query_type])
            }
        }

        return {
            activeIndex,
            handleSelect,
            posts,
            enterPostDetail,
            loading,
            noMore,
            disabled, 
            load,
            route,
            users,
            enterProfile,
        }

    }
}
</script>

<style scoped>
.card-out {
    margin-top: 20px;
    border: none;
}

.horizontal-line {
  border-bottom: 1px solid lightgray;
}

.avatar{
    /* width: 30px; 
    height: 30px;  */
    border-radius: 50%;
}


.name {
    font-weight: bold;
}

.publish-time{
    color: #808080;
}

.img-thumbnail {
    aspect-ratio: 1/1;
    max-width: 100%;
    object-fit: cover;
    border: none;
}

.post-list:hover{
    cursor: pointer;
    background-color: #F6F9FA;
}

svg {
    max-height: 20px;
    max-width: 20px;
}

.load-status {
    color: #99A2AA;
    justify-content: center;
    text-align: center;
    margin-top: 16px;
    margin-bottom: 0px;
}

</style>