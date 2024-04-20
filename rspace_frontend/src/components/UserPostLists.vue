<template>
    <div class="card card-out" style="background-color:transparent; z-index: 3; ">
        <div class="card-body" >
            <div class="card" >
                <div class="card-body">
                    <div class="row justify-content-center">
                        <div class="col-12">
                            <div class="mypost-title" v-if="is_me === true">我的帖子</div>
                            <div class="mypost-title" v-else>{{userinfo.username}}的帖子</div>
                        </div>
                    </div>
                    <br/>
                    <div v-if="posts && userinfo" >
                        <div v-for="(post, index) in posts.posts" :key="post.ID">
                            <div class="horizontal-line"></div>
                            <!-- :class="{ expanded: index === expandedIndex }" -->
                            <!-- 动态绑定style -->
                            <div class="card single-post"   :style="{ maxHeight: isCollapsed[index] ? '300px' : 'none'}">
                                <div class="card-body">
                                    <!-- 展示头像和姓名 -->
                                    <div class="row">
                                        <div class="col-md-1 col-3">
                                            <img class="img-fluid avatar" :src="userinfo.avatar" alt="">
                                        </div>
                                        <!-- align-items: center; -->
                                        <div class="col-md-11 col-9" style="display: flex;  padding-left: 0px; ">
                                            <div class="row" style="width: 100%;">
                                                <div class="col-8 col-md-9" style="font-weight: bold; padding-right: 0px;">
                                                    {{userinfo.username}}
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
                                    <!-- <div class="row justify-content-center" style="padding-top: 0px; ">
                                        <div class="col-3">
                                            <img class="img-thumbnail" :src="post.image" alt="" @click="showFullImage(post.image)">
                                        </div>
                                        <div class="col-3">
                                            <img class="img-thumbnail" :src="post.image" alt="" @click="showFullImage(post.image)">
                                        </div>
                                        <div class="col-3">
                                            <img class="img-thumbnail" :src="post.image" alt="" @click="showFullImage(post.image)">
                                        </div>
                                    </div>
                                    <div class="row justify-content-center" style="padding-top: 0px; ">
                                        <div class="col-3">
                                            <img class="img-thumbnail" :src="post.image" alt="" @click="showFullImage(post.image)">
                                        </div>
                                        <div class="col-3">
                                            <img class="img-thumbnail" :src="post.image" alt="" @click="showFullImage(post.image)">
                                        </div>
                                    </div> -->
                                    <!-- <div class="expand">
                                        <button v-if="isCollapsed[index]" @click="expandCard(index)">展开</button>
                                        <button v-else @click="collapseCard(index)">收起</button>
                                    </div> -->
                                    <!-- <template v-if="isPostCollapsed(post.id)">
                                        <div class="expand">
                                            <button @click="expandCard(post.id)">展开</button>
                                        </div>
                                    </template>
                                    <template v-else>
                                        <div class="expand">
                                            <button @click="collapseCard(post.id)">收起</button>
                                        </div>
                                    </template> -->
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

</template>


<script>
import { ref } from 'vue'
import router from '@/router/index';   //@定位src目录
// import {useStore} from 'vuex'
export default {
    name:"UserPostLists",
    components:{},
    props:{  //父组件向子组件传递信息props  从父组件获取的信息
        posts:{
            type:Object,
            requried:true,
        },
        userinfo:{
            type:Object,
            requried:true,
        },
        is_me:{
            type:Boolean,
            requried:true,
        },
    },
    // setup(props){  //在setup()函数中，将props引入到函数的参数中，以便在setup()函数中使用
    //     console.log(props.posts)  // // 可以通过props.posts来访问父组件传递的posts对象
    // }
    emits: ['changePage'], // 声明 changePage 事件
    setup(props, context) {
        // const store = useStore()
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
        let isCollapsed = ref(Array((props.posts.count ? props.posts.count : 100)).fill(true));
        // const isCollapsed = ref({});
        // const expandedIndex = ref(-1);

        // console.log(props.posts)

        // 在Vue的生命周期中，setup()函数在组件实例被创建时执行，但是在模板渲染之前执行。
        // 因此，在初始渲染时，console.log(isCollapsed.value[0])语句被执行时，isCollapsed数组的值可能还没有被设置。
        // onMounted(() => {
        //     // isCollapsed = ref(Array(props.posts.count).fill(true));
        //     console.log(isCollapsed.value[1])
        // })
        // console.log(isCollapsed.value[0])
        const expandCard = (index) => {
            isCollapsed.value[index] = false;
            // console.log(isCollapsed.value)
            // console.log(index)
            // expandedIndex.value = index;
        };

        const collapseCard = (index) => {
            isCollapsed.value[index] = true;
            // expandedIndex.value = -1;
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

        //
        const page_size = ref(3)
        // const current_page = ref(1)
        const changePage = ()=>{
            
            context.emit('changePage')
        }
        return {
            showModal,
            modalImage,
            showFullImage,
            closeModal,
            isCollapsed,
            // expandedIndex,
            expandCard,
            collapseCard,
            enterPostDetail,
            changePage,
            page_size,
            // current_page,
            // isPostCollapsed,
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

.mypost-title{
    font-weight: bold;
}

.single-post{
    margin-top: 10px;
    /* max-height: 200px; */
    overflow: hidden;  /*以便在收起时隐藏超出的内容。  */
}

/* .single-post.expanded {
  max-height: none;
} */

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
