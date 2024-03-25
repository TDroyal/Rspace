<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-md-9 col-12">
                <div v-for="(post, index) in posts.posts" :key="post.ID" >
                    <div class="card card-single">
                        <div class="card-body">
                            <div class="row justify-content-center">
                                <div class="col-9">
                                    <div class="post-type">{{post_type_map[post.type]}}</div>
                                </div>
                                <div class="col-3" >
                                    <div class="post-details" @click="enterPostDetail(post.ID)">详情</div>
                                </div>
                            </div>
                            <!-- <br/> -->
                            <div class="horizontal-line"></div>
                            <div class="card card-single-inner" :style="{ maxHeight: isCollapsed[index] ? '500px' : 'none'}">
                                <div class="card-body">
                                    <!-- 展示头像和姓名 -->
                                    <div class="row">
                                        <div class="col-3 col-md-1">
                                            <img class="img-fluid avatar" :src="post.avatar" @click="enterUserProfile(post.user_id)" alt="">
                                        </div>
                                        <!-- align-items: center; -->
                                        <div class="col-8 col-md-5" style="display: flex;  padding-left: 0px; ">
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
                            <div class="horizontal-line"></div>
                            <!-- 左边是点赞/收藏  右边是评论 -->
                            <div class="row">
                                <div class="col-6 text-center">
                                    <div class="collect" v-if="iscollect[index] === false" @click="changeCollectStatusForAPost(index, post.ID)">
                                        <!-- 旁边还要显示被点赞/收藏的数量，还要判断是否被我收藏 -->
                                        <svg t="1710466679468" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5759" width="200" height="200"><path d="M512 901.746939c-13.583673 0-26.122449-4.179592-37.093878-13.061225-8.881633-7.314286-225.697959-175.020408-312.424489-311.379592C133.746939 532.37551 94.040816 471.24898 94.040816 384.522449c0-144.718367 108.146939-262.269388 240.326531-262.269388 67.395918 0 131.657143 30.82449 177.632653 84.636735 45.453061-54.334694 109.191837-84.636735 177.110204-84.636735 132.702041 0 240.326531 117.55102 240.326531 262.269388 0 85.159184-37.093878 143.673469-67.395919 191.216327l-1.044898 1.567346c-86.726531 136.359184-303.542857 304.587755-312.424489 311.379592-10.44898 8.359184-22.987755 13.061224-36.571429 13.061225z" fill="#8a8a8a" p-id="5760" data-spm-anchor-id="a313x.search_index.0.i15.73333a81m6RRCQ" class="selected"></path></svg>
                                        <span v-if="post.iscollect_count > 0">{{post.iscollect_count}}</span>
                                    </div>
                                    <div class="collect" v-else @click="changeCollectStatusForAPost(index, post.ID)">
                                        <svg t="1710466089661" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3591" width="200" height="200"><path d="M512 901.746939c-13.583673 0-26.122449-4.179592-37.093878-13.061225-8.881633-7.314286-225.697959-175.020408-312.424489-311.379592C133.746939 532.37551 94.040816 471.24898 94.040816 384.522449c0-144.718367 108.146939-262.269388 240.326531-262.269388 67.395918 0 131.657143 30.82449 177.632653 84.636735 45.453061-54.334694 109.191837-84.636735 177.110204-84.636735 132.702041 0 240.326531 117.55102 240.326531 262.269388 0 85.159184-37.093878 143.673469-67.395919 191.216327l-1.044898 1.567346c-86.726531 136.359184-303.542857 304.587755-312.424489 311.379592-10.44898 8.359184-22.987755 13.061224-36.571429 13.061225z" fill="#d81e06" p-id="3592"></path></svg>
                                        <span v-if="post.iscollect_count > 0">{{post.iscollect_count}}</span>
                                    </div>
                                </div>
                                <div class="col-6 text-center">
                                    <!-- <div>评论</div> -->
                                    <div class="comment" v-if="commentisCollapsed[index] === true" @click="expandComment(index, post.ID)">
                                        <!-- 评论旁边也要显示有多少条评论 -->
                                        <svg t="1710466380019" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4748" width="200" height="200"><path d="M853.333333 768c35.413333 0 64-20.650667 64-55.978667V170.581333A63.978667 63.978667 0 0 0 853.333333 106.666667H170.666667C135.253333 106.666667 106.666667 135.253333 106.666667 170.581333v541.44C106.666667 747.285333 135.338667 768 170.666667 768h201.173333l110.016 117.44a42.752 42.752 0 0 0 60.586667 0.042667L651.904 768H853.333333z m-219.029333-42.666667h-6.250667l-115.797333 129.962667c-0.106667 0.106667-116.010667-129.962667-116.010667-129.962667H170.666667c-11.776 0-21.333333-1.621333-21.333334-13.312V170.581333A21.205333 21.205333 0 0 1 170.666667 149.333333h682.666666c11.776 0 21.333333 9.536 21.333334 21.248v541.44c0 11.754667-9.472 13.312-21.333334 13.312H634.304zM341.333333 490.666667a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z m170.666667 0a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z m170.666667 0a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z" fill="#3D3D3D" p-id="4749"></path></svg>
                                        <span v-if="post.comments.count > 0">{{post.comments.count}}</span>
                                    </div>
                                    <div class="comment" v-else @click="collapseComment(index)">
                                        <!-- 评论旁边也要显示有多少条评论 -->
                                        <svg t="1710466380019" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4748" width="200" height="200"><path d="M853.333333 768c35.413333 0 64-20.650667 64-55.978667V170.581333A63.978667 63.978667 0 0 0 853.333333 106.666667H170.666667C135.253333 106.666667 106.666667 135.253333 106.666667 170.581333v541.44C106.666667 747.285333 135.338667 768 170.666667 768h201.173333l110.016 117.44a42.752 42.752 0 0 0 60.586667 0.042667L651.904 768H853.333333z m-219.029333-42.666667h-6.250667l-115.797333 129.962667c-0.106667 0.106667-116.010667-129.962667-116.010667-129.962667H170.666667c-11.776 0-21.333333-1.621333-21.333334-13.312V170.581333A21.205333 21.205333 0 0 1 170.666667 149.333333h682.666666c11.776 0 21.333333 9.536 21.333334 21.248v541.44c0 11.754667-9.472 13.312-21.333334 13.312H634.304zM341.333333 490.666667a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z m170.666667 0a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z m170.666667 0a42.666667 42.666667 0 1 0 0-85.333334 42.666667 42.666667 0 0 0 0 85.333334z" fill="#3D3D3D" p-id="4749"></path></svg>
                                        <span v-if="post.comments.count > 0">{{post.comments.count}}</span>
                                    </div>    
                                </div>
                            </div>
                            <div class="horizontal-line"></div>
                            <div class="row comment-card" :style="{ maxHeight: commentisCollapsed[index] ? '0px' : 'none', marginTop: commentisCollapsed[index] ? '0px' : '20px'}">
                                <div class="col-12">
                                    <div class="mb-3">
                                        <!-- <label for="exampleFormControlTextarea1" class="form-label">Example textarea</label> -->
                                        <textarea class="form-control" id="comment" rows="3" placeholder="在这里写评论..." maxlength="200" v-model="comments[index]"></textarea>
                                        <!-- 提交评论到对应的数据库表中 -->
                                        <div class="postAComment" @click="postAComment(index, post.ID)">
                                            提交评论
                                        </div>
                                    </div>
                                </div>
                                <div v-for="(comment) in post.comments.comments" :key="comment.id">
                                    <div class="horizontal-line" style="margin-top: 10px;"></div>
                                    <div class="comment-card-list" >
                                        <div class="row" style="margin-top: 10px;">
                                            <div class="col-md-1 col-3">
                                                <img class="img-fluid avatar" :src="comment.avatar" @click="enterUserProfile(comment.user_id)" alt="">
                                            </div>
                                                <!-- align-items: center; -->
                                            <div class="col-md-5 col-9" style="display: flex;  padding-left: 0px; ">
                                                <div class="row">
                                                    <div class="col-12 username" style="font-weight: bold;" @click="enterUserProfile(comment.user_id)">
                                                        {{comment.name}}
                                                    </div>
                                                    <div class="col-12 post-time" style="color: gray;">
                                                        {{comment.CreatedAt}}
                                                    </div>
                                                </div>
                                            </div>
                                        </div>
                                        <div class="row">
                                            <div class="col-md-11 col-9">
                                                {{comment.comment}}
                                            </div>
                                            <div class="col-md-1 col-3 delete-comment-button" v-if="comment.user_id === $store.state.user.id" @click="deleteAComment(index, post.ID, comment.id)">删除</div>
                                        </div>
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
import {BackendRootURL} from '../common_resources/resource';
export default {
    name:"Home",
    components: { Content },

    setup()
    {
        const store = useStore()
        const post_type_map = reactive({
            '1':'日常',
            '2':'新鲜事',
            '3':'笔记',
            '4':'其它'
        })
        // const user = reactive({
        //     ...store.state.user
        // })
        const posts = reactive({
            count:0,
            posts:[],
        })
        // 从数据库中读取最新的5条博客进行展示，返回的信息有哪些要仔细考虑
        $.ajax({
            url: BackendRootURL + "/homepost/getposts/",
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
                    // console.log(post)
                    // console.log(imageurl)
                    posts.posts[i].image = ParseImageUrl(imgstr)
                    posts.posts[i].CreatedAt = FormatDateTime(posts.posts[i].CreatedAt)
                    posts.posts[i].avatar = BackendRootURL + "/static/avatar/" + posts.posts[i].avatar
                    posts.posts[i].type = post.type
                    posts.posts[i].comments = {
                        count:0,
                        comments:[],
                    }
                    posts.posts[i].iscollect_count = 0    // 每个作品被收藏（点赞）的数量
                }
                // console.log(posts)
                
                for(let index = 0; index < posts.count; index ++ ) {
                    // 再获取当前每个post的评论数
                    get_comments_by_post_id(index, posts.posts[index].ID)
                    // 获得当前每个post的点赞数，以及每个帖子是否被当前用户收藏
                    get_iscollect_count_by_post_id(index, posts.posts[index].ID, store.state.user.id)
                }

                
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
        // 默认都是收起的
        let isCollapsed = ref(Array((posts.count ? posts.count : 100)).fill(true));
        let commentisCollapsed = ref(Array((posts.count ? posts.count : 100)).fill(true))
        const expandCard = (index) => {
            isCollapsed.value[index] = false;
        };
        const collapseCard = (index) => {
            isCollapsed.value[index] = true;
        };

        

        const get_comments_by_post_id = (index, post_id)=>{
            const comment_list = reactive({
                count:0,
                comments:[],
            })
            $.ajax({
                url: BackendRootURL + "/homepost/getcommentsbypostid/",
                type:"GET",
                data:{
                    post_id:post_id,
                },
                success(resp) {
                    if(resp.status !== 0) {
                        ElMessage.error("获取评论失败")
                        return 
                    }
                    if(resp.data === null) {
                        return 
                    }

                    // console.log(resp.data)
                    comment_list.count = resp.data.length
                    comment_list.comments = resp.data
                    for(let i = 0; i < comment_list.count; i ++ ) {
                        comment_list.comments[i].avatar = BackendRootURL + "/static/avatar/" + comment_list.comments[i].avatar
                        comment_list.comments[i].CreatedAt = FormatDateTime(comment_list.comments[i].CreatedAt)
                    } 
                    //把comment_list放到对应的index下
                    if(comment_list.count === 0) {
                        // console.log(55555555555)
                        posts.posts[index].comments = {
                            count:0,
                            comments:[],
                        }
                    }else {
                        // console.log(66666666666666666666, comment_list.count)
                        posts.posts[index].comments = comment_list
                    }
                    
                    // console.log(posts.posts[index].comments)
                    // console.log(posts)
                    // console.log(index, posts.posts[index].comments)
                },
                error(resp) {
                    console.log(resp)
                }
            })
        }

        const expandComment = (index, post_id)=>{  //点击展开评论，才去后端获取，根据帖子id返回该帖子对应的评论和评论用户的id，头像，name，不需要jwt认证
            //先获取数据
            get_comments_by_post_id(index, post_id)
            // console.log(post_id)

            //再展开
            commentisCollapsed.value[index] = false
        }

        const collapseComment = (index)=>{
            commentisCollapsed.value[index] = true
        }

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
        // 用户提交的评论
        // const comments = ref('')
        const comments = ref(Array((posts.count ? posts.count : 100)).fill(''))
        // 评论提交
        const postAComment = (index, post_id)=>{
            if(store.state.user.is_login === false) {
                router.push({
                    name:"Login",
                })
                comments.value[index] = ''
                return 
            }
            // 登录的用户就可以存评论了

            if(comments.value[index] === '') {
                ElMessage({
                    message: '评论不能为空',
                    type: 'warning',
                })
                return 
            }

            // console.log("帖子id=", post_id, store.state.user.id, comments.value, store.state.user.jwt)

            //评论应该先显示在前端（把新的评论push到评论列表中即可），然后提醒评论成功，然后将评论存入数据库，然后清空comments.value = ''
            
            $.ajax({
                url: BackendRootURL + "/post/uploadcomment/",
                type:"POST",
                data:{
                    post_id:post_id,
                    user_id:store.state.user.id,
                    comment:comments.value[index],
                },
                headers:{
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp) {
                    if(resp.status !== 0) {
                        ElMessage({
                            message: '评论失败，请稍后重试',
                            type: 'error',
                        })
                        return 
                    }
                    ElMessage({
                        message: '评论成功',
                        type: 'success',
                    })
                    comments.value[index] = ""
                    //渲染每个帖子对应的评论列表, 成功后。
                    get_comments_by_post_id(index, post_id)
                },
                error(resp) {
                    ElMessage({
                        message: '评论失败，请稍后重试',
                        type: 'error',
                    })
                    console.log(resp)
                }
            })

        }

        const deleteAComment = (index, post_id,comment_id)=>{  //获得comment的id删
            // console.log(comment_id)
            //删完后，重新查询一次评论
            $.ajax({
                url: BackendRootURL + "/homepost/deletecommentsbycommentid/",
                type:"DELETE",
                // data:{
                //     comment_id: comment_id,
                // },
                data:JSON.stringify({ comment_id: comment_id }),  //必须这样写？delete???
                // params:{
                //     comment_id:comment_id,
                // },
                headers:{
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp){
                    if(resp.status === 0) {
                        ElMessage.success("成功删除评论")
                        get_comments_by_post_id(index, post_id)
                        return 
                    }
                    ElMessage.error("删除评论失败")
                },
                error(resp) {
                    ElMessage.error("删除评论失败")
                    console.log(resp)
                }
            })
        }
        
        //判断每个作品当前用户是否collect，默认是未收藏
        const iscollect = ref(Array((posts.count ? posts.count : 100)).fill(false))

        const changeCollectStatusForAPost = (index, post_id)=>{  //点击改变对该帖子的收藏状态
            if(store.state.user.is_login === false) {
                router.push({
                    name:"Login",
                })
                return 
            }
            // 登录的用户就可以收藏了
            // console.log(index, post_id, iscollect.value[index])
            //根据post_id和user_id先判断数据库里面是否存在
            $.ajax({
                url: BackendRootURL + "/post/changecollectStatus/",
                type:"POST",
                data:{
                    post_id:post_id,
                    user_id:store.state.user.id,
                    status: iscollect.value[index] === true ? 0 : 1,
                },
                headers:{
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp) {
                    // console.log(resp)
                    if(resp.status !== 0) {
                        ElMessage.error("收藏失败，请稍后重试")
                        return 
                    }
                    //渲染每个帖子对应的评论列表, 成功后。
                    // ElMessage.success("收藏成功")
                    // 将当前帖子收藏数量+-1，不像后端请求，节约请求次数
                    if(iscollect.value[index] === true) {
                        iscollect.value[index] = false
                        posts.posts[index].iscollect_count -= 1
                    }else if(iscollect.value[index] === false) {
                        iscollect.value[index] = true
                        posts.posts[index].iscollect_count += 1
                    }
                },
                error(resp) {
                    ElMessage({
                        message: '收藏失败，请稍后重试',
                        type: 'error',
                    })
                    console.log(resp)
                }
            })

        }

        const get_iscollect_count_by_post_id = (index, post_id, user_id) => {  //不需要jwt认证
            $.ajax({
                url: BackendRootURL + "/homepost/getiscollectcountbypostid/",
                type:"GET",
                data:{
                    post_id:post_id,
                    user_id:user_id,
                },
                success(resp) {
                    if(resp.status !== 0) {
                        ElMessage.error("获取点赞数量失败")
                        return 
                    }
                    // console.log(resp)
                    posts.posts[index].iscollect_count = resp.data.count
                    
                    //更新iscollect的值
                    if(resp.data.is_collected === 1) {
                        iscollect.value[index] = true
                    }else{
                        iscollect.value[index] = false
                    }
                    
                },
                error(resp) {
                    console.log(resp)
                }
            })
        }
        
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
            posts,
            comments,
            iscollect,
            // user,
            showModal,
            modalImage,
            showFullImage,
            closeModal,
            isCollapsed,
            commentisCollapsed,
            get_comments_by_post_id,
            expandCard,
            collapseCard,
            enterUserProfile,
            expandComment,
            collapseComment,
            postAComment,
            post_type_map,
            deleteAComment,
            changeCollectStatusForAPost,
            get_iscollect_count_by_post_id,
            enterPostDetail,
        }
    }
}

</script>


<style scoped>

svg {
    max-height: 20px;
    max-width: 20px;
}

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
    /* min-width: 30px; */
    /* max-width: 50px; */
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
    color: lightblue;
}

.comment:hover {
    color: blue;
    cursor: pointer;
}

.collect:hover {
    color:blue;
    cursor: pointer;

}

.comment-card {
    overflow: hidden;  /*以便在收起时隐藏超出的内容。  */
}

.postAComment {
    margin-top: 10px;
    float: right;  /* 这会使 <div> 元素向右浮动，与其父元素的右边缘对齐。 */
    color: gray;
}

.postAComment:hover {
    color: blue;
    cursor: pointer;
}

.comment-card-list{
    /* background-color: rgb(250, 246, 246); */
    padding-left: 20px;
}

.delete-comment-button{
    float: right;
    color: gray;
}

.delete-comment-button:hover {
    cursor: pointer;
    color: blue;
}

.post-details{
    float: right;
    font-weight: bold;
}

.post-details:hover {
    color: rgb(109, 193, 221);
    cursor: pointer;
}

</style>