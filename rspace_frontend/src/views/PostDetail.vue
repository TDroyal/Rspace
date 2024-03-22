<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-10">
                <div class="card card-out">
                    <div class="card-body">
                        details
                    </div>
                </div>
            </div>
        </div>
    </Content>

</template>


<script>
import Content from '../components/Content.vue'
import router from '@/router/index';   //@定位src目录 
import $ from 'jquery'
import BackendRootURL from '../common_resources/resource'
import { ElMessage } from 'element-plus';
import { reactive } from 'vue';
// 这个界面应该有(如果这个帖子是当前登录用户的那么)编辑帖子按钮（进入一个新的页面编辑帖子/updatepost/:postid）和删除帖子按钮
export default {
    name:"PostDetail",
    components:{Content},
    setup()
    {
        //读取路由上的post_id，然后通过post_id读取post的信息。

        // console.log(router.currentRoute.value.params.postid)
        let post = reactive({})

        // 获得帖子的基本信息
        $.ajax({
            url: BackendRootURL + "/homepost/getpostbypostid/",
            type:"GET",
            data: {
                postid: router.currentRoute.value.params.postid,
            },
            success(resp) {
                console.log(resp)
                if(resp.status !== 0) {
                    ElMessage.error("获取帖子信息失败")
                    return 
                }
                post = resp.data
                // console.log(post) 这个得到的post数据进行展示
            },
            error(resp) {
                console.log(resp)
                ElMessage.error("获取帖子信息失败")
            }
        })
        // 还需要获得帖子的评论和点赞信息

        return {
            post,
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


</style>