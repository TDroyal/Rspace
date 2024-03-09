<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-9">
                <div v-for="(post, index) in posts.posts" :key="post.ID" >
                    <div class="card card-single">
                        <div class="card-body">
                            <div class="row justify-content-center">
                                <div class="col-12">
                                    <div class="post-type">日常</div>
                                </div>
                            </div>
                            <!-- <br/> -->
                            <div class="horizontal-line"></div>
                            <div class="card card-single-inner" :style="{ maxHeight: isCollapsed[index] ? '500px' : 'none'}">
                                <div class="card-body">
                                    <!-- 展示头像和姓名 -->
                                    <div class="row">
                                        <div class="col-1">
                                            <img class="img-fluid avatar" :src="post.avatar" @click="enterUserProfile(post.user_id)" alt="">
                                        </div>
                                        <!-- align-items: center; -->
                                        <div class="col-5" style="display: flex;  padding-left: 0px; ">
                                            <div class="row">
                                                <div class="col-12 username" style="font-weight: bold;" @click="enterUserProfile(post.user_id)">
                                                    {{post.name}}
                                                </div>
                                                <div class="col-12 post-time" style="color: gray;">
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
                            <div class="row">
                                <div class="col-12">
                                    <div class="expand">
                                        <button v-if="isCollapsed[index] === true" @click="expandCard(index)">展开</button>
                                        <button v-if="isCollapsed[index] === false" @click="collapseCard(index)">收起</button>
                                    </div> 
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

        <!-- 模态框 -->
        <div class="modal" :class="{ 'show': showModal }" v-if="modalImage" >
            <div class="modal-dialog modal-dialog-centered">
                <div class="modal-content">
                <button class="close-button" @click="closeModal">&times;</button>
                <img :src="modalImage" class="img-fluid" alt="Full Image">
                </div>
            </div>
        </div>

    </Content>
</template>

<script>
import Content from '../components/Content.vue'
import { useStore } from 'vuex'
import { ref, reactive } from 'vue'
import $ from 'jquery'
import { ElMessage } from 'element-plus'
import FormatDateTime from '../utils/DateTime'
import ParseImageUrl from '../utils/ParseImageUrl'
// import prefix from '../utils/ParseImageUrl'
import router from '@/router/index';   //@定位src目录
export default {
    name:"Home",
    components: { Content },

    setup()
    {
        const store = useStore()
        // const user = reactive({
        //     ...store.state.user
        // })
        const posts = reactive({
            count:0,
            posts:[],
        })
        // 从数据库中读取最新的5条博客进行展示，返回的信息有哪些要仔细考虑
        $.ajax({
            url:"http://127.0.0.1:9090/homepost/getposts/",
            type:"GET",
            success(resp) {
                // console.log(resp)
                if(resp.status != 0) {
                    ElMessage({
                        message: '获取首页信息失败，请稍后重试',
                        type: 'error',
                    })
                    return 
                }
                posts.posts = resp.data
                posts.count = posts.posts.length
                // console.log(posts)
                // 遍历每个posts的Image，对其进行分割并拼接，组成一个数组
                for(let i = 0; i < posts.count; i ++ ) {
                    let post = posts.posts[i]
                    let imgstr = post.image
                    
                    // console.log(imageurl)
                    posts.posts[i].image = ParseImageUrl(imgstr)
                    posts.posts[i].CreatedAt = FormatDateTime(posts.posts[i].CreatedAt)
                    posts.posts[i].avatar = "http://127.0.0.1:9090/static/avatar/" + posts.posts[i].avatar

                }
                // console.log(posts)
            },
            error(resp) {
                ElMessage({
                    message: '获取用户帖子失败，请稍后重试',
                    type: 'error',
                })
                console.log(resp)
            },
        })

        // 点击放大图片和关闭功能
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

        // 展开和收起功能
        let isCollapsed = ref(Array((posts.count ? posts.count : 100)).fill(true));
        
        const expandCard = (index) => {
            isCollapsed.value[index] = false;
        };
        const collapseCard = (index) => {
            isCollapsed.value[index] = true;
        };

        // 进入用户的个人空间
        const enterUserProfile = (user_id)=>{
            if(store.state.user.is_login) {
                router.push({
                    name:"Profile",
                    params:{
                        userid:user_id,
                    }
                })
            }else {
                router.push({
                    name:"Login",
                })
            }
            
        }


        return {
            posts,
            // user,
            showModal,
            modalImage,
            showFullImage,
            closeModal,
            isCollapsed,
            expandCard,
            collapseCard,
            enterUserProfile,
        }
    }
}

</script>


<style scoped>


.post-type{
    font-weight:bold;
    font-size: 20px;
}

.card-single {
    margin-top: 20px;
    /* box-shadow: none; */
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
    border: none;
    
}

.card-single-inner{
    border: none;
    overflow: hidden;  /*以便在收起时隐藏超出的内容。  */
    transition:all 2s ease-in ;
    -moz-transition: 2s;
    -webkit-transition: height 2s;
    -o-transition: height 2s;
}

.avatar {
    border-radius: 50%;
}

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
}

.modal.show {
  display: flex !important;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
}

.close-button {
  position: absolute;
  top: 10px;
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

.avatar:hover {
    cursor: pointer;
}

.username:hover {
    cursor: pointer;
}

</style>