<template>
    <div class="card card-out" style="background-color:transparent; z-index: 3; ">
        <div class="card-body" >
            <div class="card" >
                <div class="card-body">
                    <div class="row justify-content-center">
                        <div class="col-12">
                            <div class="mypost-title" v-if="is_me === true">我的帖子</div>
                            <div class="mypost-title" v-else>{{user.username}}的帖子</div>
                        </div>
                    </div>
                    <br/>
                    <div v-if="posts && user" >
                        <div v-for="(post, index) in posts.posts" :key="post.ID">
                            <div class="horizontal-line"></div>
                            <!-- :class="{ expanded: index === expandedIndex }" -->
                            <!-- 动态绑定style -->
                            <div class="card single-post"   :style="{ maxHeight: isCollapsed[index] ? '300px' : 'none'}">
                                <div class="card-body">
                                    <!-- 展示头像和姓名 -->
                                    <div class="row">
                                        <div class="col-md-1 col-3">
                                            <img class="img-fluid avatar" :src="user.avatar" alt="">
                                        </div>
                                        <!-- align-items: center; -->
                                        <div class="col-md-11 col-9" style="display: flex;  padding-left: 0px; ">
                                            <div class="row" style="width: 100%;">
                                                <div class="col-8 col-md-9" style="font-weight: bold; padding-right: 0px;">
                                                    {{user.username}}
                                                </div>
                                                <div class="col-4 col-md-3 post-details" @click="enterPostDetail(post.ID)">
                                                    详情
                                                </div>
                                                <div class="col-12 post-time" style="color: gray; padding-right: 0px;">
                                                    {{post.CreatedAt}}
                                                </div>
                                            </div>
                                        </div>
                                    </div>

                                    <!-- 展示文字信息 -->
                                    <div class="row" style="padding-top: 10px; ">
                                        <!-- <div class="col-12"> -->
                                            <div>{{post.content}}</div>
                                        <!-- </div> -->
                                    </div>
                                    <!-- 下面展示9宫格图片 -->
                                    <div class="row justify-content-center" style="padding-top: 10px; ">
                                        <!-- 我们使用了d-flex类和flex-wrap类来创建一个弹性容器，并将图片容器包装在其中。d-flex类将容器设置为弹性布局，而flex-wrap类将图片容器设置为自动换行。 -->
                                        <div class="d-flex flex-wrap">
                                            <div class="col-4" v-for="(image) in post.image" :key="image.id">
                                                <img class="img-thumbnail" :src="image.image" alt="" @click="showFullImage(image.image)">
                                            </div>
                                        </div>
                                    </div>
                                </div>
                                
                            </div>
                              
                            <div class="expand">
                                <button v-if="isCollapsed[index] === true" @click="expandCard(index)">展开</button>
                                <button v-if="isCollapsed[index] === false" @click="collapseCard(index)">收起</button>
                            </div>
                            <el-pagination hide-on-single-page  v-if="index === posts.count - 1" style="justify-content: center; " :page-size="page_size" v-model:current-page="$store.state.pagination.user_postlist_page" large background layout="prev, pager, next" :total="posts.total_count" class="mt-4" @change="changePage" :pager-count="5"/>
                                  
                        </div>
                    </div>
                    
                    <div v-if="posts.total_count === 0">
                        <el-empty description="没有数据" />
                    </div>
                </div>
            </div>
        </div>

        <!-- 模态框 -->
        <div class="modal" :class="{ 'show': showModal }" v-if="modalImage" @click="closeModal">
            <!-- <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content"> -->
                <button class="close-button" @click="closeModal">&times;</button>
                <img :src="modalImage" class="img-fluid show-image"  alt="Full Image">
                <!-- </div>
            </div> -->
        </div>
    </div>
</template>

<script>
import { useStore } from 'vuex'
import { useRoute } from 'vue-router';
import { ref, reactive, watch } from 'vue'
import router from '@/router/index';   //@定位src目录
import $ from 'jquery'
import {BackendRootURL} from '../common_resources/resource';
import { ElMessage } from 'element-plus'
import {FormatDateTime} from '../utils/DateTime'
import ParseImageUrl from '../utils/ParseImageUrl'

export default{
    name:"ProfilePostLists",
    components:{},
    setup(){
        //需要获取换页，posts，user, is_me 
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

        const route = useRoute();
        let user_id = parseInt(route.params.userid);  //当前打开这个用户的id，从url上获取
        const is_me = ref(user_id === store.state.user.id)  //判断当前查看的空间是否是我自己
        const user = reactive({
            // ...store.state.user
        })
        const page_size = ref(3)
        const current_page = ref(1)
        const posts = reactive({
            total_count:0,
            count:0,
            posts:[],
        })

        //先获取用户个人基本信息
        const getUserInfo = async () => {
            // if (is_me.value === false) {
                // console.log(is_me.value)
                try {
                    const resp = await $.ajax({
                        url: BackendRootURL + "/myspace/getuserinfo/",
                        type: "GET",
                        data: {
                            user_id: user_id,  //这个是当前查看的用户id
                        },
                        headers: {
                            'Authorization': "Bearer " + store.state.user.jwt,  //这个里面存入的是当前登录的用户id
                        },
                    });

                    if (resp.status != 0) {
                        return;
                    }
                    // console.log(resp.data)
                    user.id = resp.data.normal_userinfo.id,
                    user.username = resp.data.normal_userinfo.name,
                    user.age = resp.data.normal_userinfo.age,
                    user.avatar = BackendRootURL + '/static/avatar/' + resp.data.normal_userinfo.avatar,
                    user.gender = resp.data.normal_userinfo.gender,
                    user.address = resp.data.normal_userinfo.address,
                    user.introduction = resp.data.normal_userinfo.introduction
                    // 再得到后端传过来的粉丝数量，关注数量，以及当前登录用户是否关注此用户
                    user.fanscount = resp.data.fanscount
                    user.followercount = resp.data.followercount
                    user.is_followed = resp.data.is_followed

                    // console.log("111111111111111", user);
                    // console.log(user)
                } catch (error) {
                    console.log(error);
                }
            // }
        };

        const getUserPosts = async () => {
            try {
                const resp = await $.ajax({
                    url:  BackendRootURL + "/myspace/getuserposts/",
                    type: "GET",
                data: {
                    user_id: user_id,
                    page_size: page_size.value,
                    current_page: store.state.pagination.user_postlist_page,
                },
                headers: {
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                });

                if (resp.status != 0) {
                    ElMessage({
                        message: '获取用户帖子失败，请稍后重试',
                        type: 'error',
                    });
                    return;
                }
                // console.log("用户帖子：", resp.data)
                posts.posts = resp.data.data;  
                posts.total_count = resp.data.count
                posts.count = posts.posts.length;
                
                // console.log("my",posts)
                if(posts.posts === null || posts.count === null || posts.count === 0) {
                    if(posts.total_count > 0) {
                        // console.log(11111)
                        store.commit('updateUserPostlistPage', store.state.pagination.user_postlist_page - 1)
                        getUserPosts()
                        return 
                    }
                }

                for (let i = 0; i < posts.count; i++) {
                    let post = posts.posts[i];
                    let imgstr = post.image;
                    posts.posts[i].image = ParseImageUrl(imgstr);
                    posts.posts[i].CreatedAt = FormatDateTime(posts.posts[i].CreatedAt);
                }
                
            } catch (error) {
                ElMessage({
                message: '获取用户帖子失败，请稍后重试',
                type: 'error',
                });
                console.log(error);
            }
        };
        // 我将获取用户信息和获取用户帖子的部分封装为异步函数getUserInfo和getUserPosts。然后，在loadData函数中使用await关键字按顺序执行这两个异步函数。
        const loadData = async () => {
            user_id = parseInt(route.params.userid);  //当前打开这个用户的id，从url上获取
            is_me.value = user_id === store.state.user.id;
            await getUserInfo();
            await getUserPosts();
            
            // 在这里可以继续操作userinfo和posts
        };
        // 监听路由参数的变化
        // 你可以使用watch()函数来监听路由参数的变化，并在参数变化时重新获取用户信息。
        watch(() => router.currentRoute.value.params, () => {
            loadData();
        }, { immediate: true }); // immediate: true 表示在组件初始化时立即执行

        // 展开和收起功能
        let isCollapsed = ref(Array((posts.count ? posts.count : 100)).fill(true));
        

        const changePage = () =>{
            getUserPosts()
            for(let i = 0; i < isCollapsed.value.length; i ++ )
            {
                isCollapsed.value[i] = true
            }
        }


        const showModal = ref(false);
        const modalImage = ref('');
        const showFullImage = (imageUrl) => {
            modalImage.value = imageUrl;
            showModal.value = true;
            // console.log(111)
        };

        const closeModal = () => {
            modalImage.value = '';
            showModal.value = false;
        };
        
        const expandCard = (index) => {
            isCollapsed.value[index] = false;
        };

        const collapseCard = (index) => {
            isCollapsed.value[index] = true;
        };
         // 进入帖子详情页面
         const enterPostDetail = (post_id)=> {
            router.push({
                name:"PostDetail",
                params:{
                    postid:post_id
                },
            })
        }

        return {
            user,
            is_me,
            posts,
            page_size,
            current_page,
            modalImage,
            isCollapsed,
            showModal,
            changePage,
            showFullImage,
            closeModal,
            expandCard,
            collapseCard,
            enterPostDetail,
        }

    },
}

</script>


<style scoped>
.card-out {
    box-shadow: none;
    border: none;
}

.mypost-title{
    font-weight: bold;
}

.single-post{
    margin-top: 10px;
    /* max-height: 200px; */
    overflow: hidden;  /*以便在收起时隐藏超出的内容。  */
}
.single-post {
    box-shadow: none;
    border: none;
}

.avatar {
    border-radius: 50%;
    /* min-width: 30px; */
    /* padding-top: 6px; */
    /* max-width: 50px; */
}

/* 通过将aspect-ratio属性设置为1/1，图片的宽度和高度将保持相等。
然后，使用max-width: 100%确保图片不会超出其父容器的宽度。
最后，使用object-fit: cover;让图片在保持纵横比的同时填充整个区域。 */
.img-thumbnail {
    /* max-width: 100%;
    max-height: 100%; */
    aspect-ratio: 1/1;
    max-width: 100%;
    object-fit: cover;
    cursor: pointer;
    border: none;
}

.modal {
  background: rgba(0, 0, 0, 0.5);
  overflow: auto;
  width: 100%;
  height: 100%;
  cursor: pointer;
}

.modal.show {
  display: flex !important;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
}

.close-button {
  position: absolute;
  top: 60px;
  right: 10px;
  font-size: 24px;
  color: gray;
  background-color: transparent;
  border: none;
  cursor: pointer;
}

.horizontal-line {
  border-bottom: 1px solid lightgray;
}

.expand {
    display: flex;
    justify-content: flex-end;
    padding-top: 10px;
}

.expand button {
  background-color: transparent;
  border: none;
  color: blue;
  cursor: pointer;
  outline: none;
}

.post-details{
    /* float: right; */
    padding-right: 0px;
    font-weight: bold;
    text-align: right;
}

.post-details:hover {
    color: rgb(109, 193, 221);
    cursor: pointer;
}

.show-image {
    /* 根据图片高宽自适应屏幕百分百 */
    object-fit: contain; 
    width: 95%; 
    height: 85%;
}


</style>