<template>
    <Content>
        <div class="row">
            <div class="col-3" style="background-color: #F7F8FA;">
                <!-- 头像 写成一个子组件 -->
                <UserProfileInfo v-bind:userinfo="user" :is_me="is_me"></UserProfileInfo>
            </div>

            <div class="col-9">
                <UserPostLists :posts="posts" :userinfo="user" :is_me="is_me"></UserPostLists>
            </div>
        </div>
    </Content>
</template>

<script>
import { ref, reactive, watch } from 'vue' //computed
import { useStore } from 'vuex'
import { useRoute } from 'vue-router';
import router from '@/router/index';   //@定位src目录
import Content from '../components/Content.vue'
import UserProfileInfo from '../components/UserProfileInfo.vue'
import UserPostLists from '../components/UserPostLists.vue'
import $ from 'jquery'
import { ElMessage } from 'element-plus'
import FormatDateTime from '../utils/DateTime'
import ParseImageUrl from '../utils/ParseImageUrl'

export default {
    name: "Profile",
    components: {Content,UserProfileInfo,UserPostLists,},
    setup() {
        const store = useStore()
        const route = useRoute();
        let user_id = parseInt(route.params.userid);  //当前打开这个用户的id，从url上获取
        // console.log(user_id)
        // let is_me = computed(() => user_id === store.state.user.id)
        const is_me = ref(user_id === store.state.user.id)
        // console.log(user_id, store.state.user.id, is_me.value)
        // 如果当前查看的userprofile不是我自己的话，user数据应该去数据库重新获取
        // 是自己的话就不重新获取
        const user = reactive({
            // ...store.state.user
        })
        const posts = reactive({
            count:0,
            posts:[],
        })
        
        //先获取用户个人基本信息
        const getUserInfo = async () => {
            // if (is_me.value === false) {
                console.log(is_me.value)
                try {
                    const resp = await $.ajax({
                        url: "http://127.0.0.1:9090/myspace/getuserinfo/",
                        type: "GET",
                        data: {
                            user_id: user_id,
                        },
                        headers: {
                            'Authorization': "Bearer " + store.state.user.jwt,
                        },
                    });

                    if (resp.status != 0) {
                        return;
                    }
                    user.id = resp.data.id,
                    user.username = resp.data.name,
                    user.age = resp.data.age,
                    user.avatar = 'http://127.0.0.1:9090/static/avatar/' + resp.data.avatar,
                    user.gender = resp.data.gender,
                    user.address = resp.data.address,
                    user.introduction = resp.data.introduction
                    // console.log("111111111111111", user);
                } catch (error) {
                    console.log(error);
                }
            // }
        };

        const getUserPosts = async () => {
            try {
                const resp = await $.ajax({
                    url: "http://127.0.0.1:9090/myspace/getuserposts/",
                    type: "GET",
                data: {
                    user_id: user_id,
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

                posts.posts = resp.data;
                posts.count = posts.posts.length;

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
        // loadData();
        // console.log(router.currentRoute.value.params)  // 捕捉userid这个params
        watch(() => router.currentRoute.value.params, () => {
            loadData();
        }, { immediate: true }); // immediate: true 表示在组件初始化时立即执行

        // loadData();

        return {
            // store,
            // prefix,
            user,
            is_me,
            // userinfo,
            posts,
        }
    }
}

</script>


<style scoped>

</style>