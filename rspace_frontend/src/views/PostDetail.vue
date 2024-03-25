<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-md-9 col-12">
                <div class="card card-out">
                    <div class="card-body">
                        <!-- 左边：帖子详情  + 右边：编辑+删除按钮 -->
                        <div class="row">
                            <div class="col-12">
                                <div class="row">
                                    <div class="col-md-4 col-6 title-detail">
                                        帖子详情
                                    </div>
                                    <div class="col-md-8 col-5">
                                        <!-- display: flex; justify-content: flex-end; div右靠齐-->
                                        <div class="row" style="display: flex; justify-content: flex-end;">
                                            <div class="col-1" >
                                                <div class="edit-post" v-if="$store.state.user.is_login === true && post.user_id === $store.state.user.id">
                                                    <svg t="1711075212319" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="2362" width="200" height="200"><path d="M540.441638 592.438751c-19.293125 19.455016-44.045054 28.15164-66.790069 7.413537-19.587472-17.858851-17.393247-40.138263 1.612221-61.376086l401.108349-414.262327c20.005579-19.048951 49.489809-28.332931 69.572988-9.365595 24.082958 22.745016 13.809569 50.635087-5.238712 70.634646l-400.339032 406.876217 0.074255 0.079608z m-6.690379 0" p-id="2363"></path><path d="M939.847456 473.915813c23.911701 0.705764 35.618694 23.114957 34.6527 48.03279v337.313932c0 63.401731-51.397045 114.798107-114.798107 114.798107H170.894005c-63.401731 0-114.798776-51.397045-114.798776-114.798776V170.45516c0-63.401731 51.397045-114.798107 114.798107-114.798107H610.028007c31.7012 0 57.399054 6.296355 57.399053 37.997555s-25.697854 35.990642-57.399053 35.990643H177.381018c-27.229797 0-49.303835 22.074038-49.303835 49.303835v675.165048c0 27.229797 22.074038 49.303835 49.303835 49.303835h671.819524c27.229797 0 49.303835-22.074038 49.303835-49.303835V521.948603c-3.311407-23.915046 12.743898-47.327026 36.659613-48.03279h4.683466z" p-id="2364"></path></svg>
                                                </div>
                                            </div>
                                            <div class="col-1">
                                                <div class="remove-post" v-if="$store.state.user.is_login === true && post.user_id === $store.state.user.id" @click="removeApost(post.ID)">
                                                    <!-- <svg t="1711075304107" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3393" width="200" height="200"><path d="M106.666667 213.333333h810.666666v42.666667H106.666667z" fill="#3D3D3D" p-id="3394"></path><path d="M640 128v42.666667h42.666667V128c0-23.573333-19.093333-42.666667-42.538667-42.666667H383.872A42.496 42.496 0 0 0 341.333333 128v42.666667h42.666667V128h256z" fill="#3D3D3D" p-id="3395"></path><path d="M213.333333 896V256H170.666667v639.957333C170.666667 919.552 189.653333 938.666667 213.376 938.666667h597.248C834.218667 938.666667 853.333333 919.68 853.333333 895.957333V256h-42.666666v640H213.333333z" fill="#3D3D3D" p-id="3396"></path><path d="M320 341.333333h42.666667v384h-42.666667zM490.666667 341.333333h42.666666v384h-42.666666zM661.333333 341.333333h42.666667v384h-42.666667z" fill="#3D3D3D" p-id="3397"></path></svg> -->
                                                    <svg t="1711084187003" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3915" width="200" height="200"><path d="M106.666667 213.333333h810.666666v42.666667H106.666667z" fill="#2c2c2c" p-id="3916"></path><path d="M640 128v42.666667h42.666667V128c0-23.573333-19.093333-42.666667-42.538667-42.666667H383.872A42.496 42.496 0 0 0 341.333333 128v42.666667h42.666667V128h256z" fill="#2c2c2c" p-id="3917"></path><path d="M213.333333 896V256H170.666667v639.957333C170.666667 919.552 189.653333 938.666667 213.376 938.666667h597.248C834.218667 938.666667 853.333333 919.68 853.333333 895.957333V256h-42.666666v640H213.333333z" fill="#2c2c2c" p-id="3918"></path><path d="M320 341.333333h42.666667v384h-42.666667zM490.666667 341.333333h42.666666v384h-42.666666zM661.333333 341.333333h42.666667v384h-42.666667z" fill="#2c2c2c" p-id="3919"></path></svg>
                                                </div>
                                            </div>
                                        </div>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div class="horizontal-line"></div>
                        <div class="row" style="margin-top: 20px;">
                            <div class="col-12">
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
                                    <div>{{post.content}}</div>
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
                    </div>
                </div>
            </div>
            <div class="col-md-3 col-12">
                <div class="card card-out">
                    <div class="card-body">
                        <div class="row">
                            <div class="col-12">
                                <div class="star-title" v-if="iscollect_count > 0">
                                    点赞{{iscollect_count}}
                                </div>
                                <div class="star-title" v-else>
                                    点赞
                                </div>
                            </div>
                        </div>
                        <div class="horizontal-line"></div>
                        <div class="row" style="margin-top: 10px">
                            <div class="col-12 text-center">
                                <div class="collect" v-if="iscollect === false" @click="changeCollectStatusForAPost(post.ID)">
                                    <!-- 旁边还要显示被点赞/收藏的数量，还要判断是否被我收藏 -->
                                    <!-- <svg t="1710466679468" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="5759" width="200" height="200"><path d="M512 901.746939c-13.583673 0-26.122449-4.179592-37.093878-13.061225-8.881633-7.314286-225.697959-175.020408-312.424489-311.379592C133.746939 532.37551 94.040816 471.24898 94.040816 384.522449c0-144.718367 108.146939-262.269388 240.326531-262.269388 67.395918 0 131.657143 30.82449 177.632653 84.636735 45.453061-54.334694 109.191837-84.636735 177.110204-84.636735 132.702041 0 240.326531 117.55102 240.326531 262.269388 0 85.159184-37.093878 143.673469-67.395919 191.216327l-1.044898 1.567346c-86.726531 136.359184-303.542857 304.587755-312.424489 311.379592-10.44898 8.359184-22.987755 13.061224-36.571429 13.061225z" fill="#8a8a8a" p-id="5760" data-spm-anchor-id="a313x.search_index.0.i15.73333a81m6RRCQ" class="selected"></path></svg> -->
                                    <svg t="1711087782127" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4993" width="200" height="200"><path d="M64 483.04V872c0 37.216 30.144 67.36 67.36 67.36H192V416.32l-60.64-0.64A67.36 67.36 0 0 0 64 483.04zM857.28 344.992l-267.808 1.696c12.576-44.256 18.944-83.584 18.944-118.208 0-78.56-68.832-155.488-137.568-145.504-60.608 8.8-67.264 61.184-67.264 126.816v59.264c0 76.064-63.84 140.864-137.856 148L256 416.96v522.4h527.552a102.72 102.72 0 0 0 100.928-83.584l73.728-388.96a102.72 102.72 0 0 0-100.928-121.824z" p-id="4994" fill="#8a8a8a"></path></svg>
                                    <span v-if="iscollect_count > 0">{{iscollect_count}}</span>
                                </div>
                                <div class="collect" v-else @click="changeCollectStatusForAPost(post.ID)">
                                    <!-- <svg t="1710466089661" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="3591" width="200" height="200"><path d="M512 901.746939c-13.583673 0-26.122449-4.179592-37.093878-13.061225-8.881633-7.314286-225.697959-175.020408-312.424489-311.379592C133.746939 532.37551 94.040816 471.24898 94.040816 384.522449c0-144.718367 108.146939-262.269388 240.326531-262.269388 67.395918 0 131.657143 30.82449 177.632653 84.636735 45.453061-54.334694 109.191837-84.636735 177.110204-84.636735 132.702041 0 240.326531 117.55102 240.326531 262.269388 0 85.159184-37.093878 143.673469-67.395919 191.216327l-1.044898 1.567346c-86.726531 136.359184-303.542857 304.587755-312.424489 311.379592-10.44898 8.359184-22.987755 13.061224-36.571429 13.061225z" fill="#d81e06" p-id="3592"></path></svg> -->
                                    <svg t="1711087782127" class="icon" viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="4993" width="200" height="200"><path d="M64 483.04V872c0 37.216 30.144 67.36 67.36 67.36H192V416.32l-60.64-0.64A67.36 67.36 0 0 0 64 483.04zM857.28 344.992l-267.808 1.696c12.576-44.256 18.944-83.584 18.944-118.208 0-78.56-68.832-155.488-137.568-145.504-60.608 8.8-67.264 61.184-67.264 126.816v59.264c0 76.064-63.84 140.864-137.856 148L256 416.96v522.4h527.552a102.72 102.72 0 0 0 100.928-83.584l73.728-388.96a102.72 102.72 0 0 0-100.928-121.824z" p-id="4994" fill="#d81e06"></path></svg>
                                    <span v-if="iscollect_count > 0">{{iscollect_count}}</span>
                                </div>
                            </div>
                        </div>

                    </div>
                </div>
                <div class="card card-out">
                    <div class="card-body">
                        <div class="row">
                            <div class="col-12">
                                <div class="comment-title" v-if="comment_list.count > 0">
                                    评论{{comment_list.count}}
                                </div>
                                <div class="comment-title" v-else>
                                    评论
                                </div>
                            </div>
                        </div>
                        <div class="horizontal-line"></div>
                        <!-- 显示评论列表 -->
                        <div class="row" style="margin-top: 10px">
                            <div class="col-12">
                                <div class="mb-3">
                                        <!-- <label for="exampleFormControlTextarea1" class="form-label">Example textarea</label> -->
                                    <textarea class="form-control" id="comment" rows="3" placeholder="在这里写评论..." maxlength="200" v-model="comments"></textarea>
                                    <!-- 提交评论到对应的数据库表中 -->
                                    <div class="postAComment" @click="postAComment(post.ID)">
                                        提交评论
                                    </div>
                                </div>
                            </div>
                            
                            <div class="col-12" v-if="comment_list.count === 0">
                                <div class="horizontal-line" style="margin-top: 10px;"></div>
                                <div style="color: lightgray;">
                                    暂时还没有评论，快来评论吧~
                                </div>
                            </div>
                            <div v-else v-for="(comment) in comment_list.comments" :key="comment.id">
                                <div class="horizontal-line" style="margin-top: 10px;"></div>
                                <div class="comment-card-list" >
                                    <div class="row" style="margin-top: 10px;">
                                        <div class="col-3">
                                            <img class="img-fluid avatar" :src="comment.avatar" @click="enterUserProfile(comment.user_id)" alt="">
                                        </div>
                                        <div class="col-9" style="display: flex;  padding-left: 0px; ">
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
                                        <div class="col-9">
                                            {{comment.comment}}
                                        </div>
                                        <div class="col-3 delete-comment-button" v-if="comment.user_id === $store.state.user.id" @click="deleteAComment(post.ID, comment.id)">删除</div>
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
import router from '@/router/index';   //@定位src目录 
import $ from 'jquery'
import {BackendRootURL} from '../common_resources/resource';
import { ElMessage,  ElMessageBox } from 'element-plus';
import { ref, reactive } from 'vue';
import { useStore } from 'vuex';
import FormatDateTime from '../utils/DateTime'
import ParseImageUrl from '../utils/ParseImageUrl'
// 这个界面应该有(如果这个帖子是当前登录用户的那么)编辑帖子按钮（进入一个新的页面编辑帖子/updatepost/:postid）和删除帖子按钮
export default {
    name:"PostDetail",
    components:{Content},
    setup()
    {
        //读取路由上的post_id，然后通过post_id读取post的信息。
        const store = useStore()
        // console.log(router.currentRoute.value.params.postid)
        let post = reactive({
            "CreatedAt":"",
            "ID":"",
            "UpdatedAt":"",
            "avatar":"",
            "content":"",
            "image":"",
            "name":"",
            "type":"",
            "user_id":'',
        })

        // 获得帖子的基本信息
        $.ajax({
            url: BackendRootURL + "/homepost/getpostbypostid/",
            type:"GET",
            data: {
                postid: router.currentRoute.value.params.postid,
            },
            success(resp) {
                // console.log(resp)
                if(resp.status !== 0) {
                    ElMessage.error("获取帖子信息失败")
                    return 
                }
                post.user_id = resp.data.user_id
                post.type = resp.data.type
                post.name = resp.data.name
                post.image = ParseImageUrl(resp.data.image)
                post.content = resp.data.content
                post.avatar = BackendRootURL + "/static/avatar/" + resp.data.avatar
                post.UpdatedAt = FormatDateTime(resp.data.UpdatedAt)
                post.ID = resp.data.ID
                post.CreatedAt = FormatDateTime(resp.data.CreatedAt)
                // console.log(post.user_id, store.state.user.id) //这个得到的post数据进行展示
                // if(store.state.user.id === post.user_id) {
                //     console.log(11111111111)
                // }
            },
            error(resp) {
                console.log(resp)
                ElMessage.error("获取帖子信息失败")
            }
        })
        // 还需要获得帖子的评论和点赞信息，封装为函数，每次点击删除评论或者提交评论就再请求一次评论信息。
        

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

        //评论列表
        const comments = ref('')
        const comment_list = reactive({
            count:0,
            comments:[],
        })
        const get_comments_by_post_id = (post_id)=>{
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
                },
                error(resp) {
                    console.log(resp)
                }
            })
        }
        get_comments_by_post_id(router.currentRoute.value.params.postid)
        // console.log(comment_list)

        // 评论提交
        const postAComment = (post_id)=>{
            if(store.state.user.is_login === false) {
                router.push({
                    name:"Login",
                })
                comments.value = ''
                return 
            }
            // 登录的用户就可以存评论了

            if(comments.value === '') {
                ElMessage({
                    message: '评论不能为空',
                    type: 'warning',
                })
                return 
            }
            $.ajax({
                url: BackendRootURL + "/post/uploadcomment/",
                type:"POST",
                data:{
                    post_id:post_id,
                    user_id:store.state.user.id,
                    comment:comments.value,
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
                    comments.value = ""
                    //渲染每个帖子对应的评论列表, 成功后。
                    get_comments_by_post_id(post_id)
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

        const deleteAComment = (post_id,comment_id)=>{  //获得comment的id删
            //删完后，重新查询一次评论
            $.ajax({
                url: BackendRootURL + "/homepost/deletecommentsbycommentid/",
                type:"DELETE",
                data:JSON.stringify({ comment_id: comment_id }),  //必须这样写？delete???
                headers:{
                    'Authorization': "Bearer " + store.state.user.jwt,
                },
                success(resp){
                    if(resp.status === 0) {
                        ElMessage.success("成功删除评论")
                        get_comments_by_post_id(post_id)
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
        
        //获取该作品被收藏的次数
        const iscollect_count = ref(0)
        //判断作品当前用户是否collect，默认是未收藏
        const iscollect = ref(false)
        const get_iscollect_count_by_post_id = (post_id, user_id) => {  //不需要jwt认证
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
                    iscollect_count.value = resp.data.count
                    
                    //更新iscollect的值
                    if(resp.data.is_collected === 1) {
                        iscollect.value = true
                    }else{
                        iscollect.value = false
                    }
                    
                },
                error(resp) {
                    console.log(resp)
                }
            })
        }
        get_iscollect_count_by_post_id(router.currentRoute.value.params.postid, store.state.user.id)

        const changeCollectStatusForAPost = (post_id)=>{  //点击改变对该帖子的收藏状态
            if(store.state.user.is_login === false) {
                router.push({
                    name:"Login",
                })
                return 
            }
            $.ajax({
                url: BackendRootURL + "/post/changecollectStatus/",
                type:"POST",
                data:{
                    post_id:post_id,
                    user_id:store.state.user.id,
                    status: iscollect.value === true ? 0 : 1,
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
                    if(iscollect.value === true) {
                        iscollect.value = false
                        iscollect_count.value -= 1
                    }else if(iscollect.value === false) {
                        iscollect.value = true
                        iscollect_count.value += 1
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

        // 删除帖子，删除成功后，返回到首页。
        const removeApost = (post_id)=>{
            console.log(post_id)
            //先是弹出一个提醒框，是否删除。
            ElMessageBox.confirm(
                '确认是否删除该帖子？',
                '警告',
                {
                    confirmButtonText: '确定',
                    cancelButtonText: '取消',
                    type: '提示',
                }
            )
            .then(() => {
                // 删除成功然后跳转到首页去。
                $.ajax({
                    url: BackendRootURL + "/homepost/deletepostbypostid/",
                    type:"DELETE",
                    data:JSON.stringify({ post_id: post_id }),  
                    headers:{
                        'Authorization': "Bearer " + store.state.user.jwt,
                    },
                    success(resp){
                        if(resp.status === 0) {
                            ElMessage.success("删除成功")
                            router.push({name:"Home"})
                            return 
                        }
                        ElMessage.error("删除失败")
                    },
                    error(resp) {
                        ElMessage.error("删除失败")
                        console.log(resp)
                    }
                })
            })
            .catch(() => {
                ElMessage({
                    type: 'info',
                    message: '取消删除',
                })
            })
        }

        return {
            store,
            post,
            showModal,
            modalImage,
            comments,
            comment_list,
            iscollect_count,
            iscollect,
            postAComment,
            enterUserProfile,
            showFullImage,
            closeModal,
            get_comments_by_post_id,
            deleteAComment,
            get_iscollect_count_by_post_id,
            changeCollectStatusForAPost,
            removeApost,
        }
    }
}

</script>


<style scoped>
.card-out {
    margin-top: 20px;
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
    border: none;
}

.title-detail{
    font-weight: bold;
    font-size: 20px;
}

svg{
    max-height: 20px;
    max-width: 20px;
}

.edit-post{
    width: 100%;
}

.remove-post {
    width: 100%;
}

.edit-post:hover {
    cursor: pointer;
}

.remove-post:hover {
    cursor: pointer;
}

.horizontal-line {
  border-bottom: 1px solid lightgray;
}

.avatar {
    border-radius: 50%;
}

.avatar:hover {
    cursor: pointer;
}

.username:hover {
    cursor: pointer;
    color: lightblue;
}

.img-thumbnail {
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


.comment-title{
    font-weight: bold;
}

.star-title {
    font-weight: bold;
}

.postAComment:hover {
    color: blue;
    cursor: pointer;
}

.comment-card-list{
    padding-left: 1
    
    0px;
}

.delete-comment-button{
    float: right;
    color: gray;
}

.delete-comment-button:hover {
    cursor: pointer;
    color: blue;
}

.postAComment {
    margin-top: 10px;
    float: right;  /* 这会使 <div> 元素向右浮动，与其父元素的右边缘对齐。 */
    color: gray;
}

.collect:hover {
    color:blue;
    cursor: pointer;

}

</style>