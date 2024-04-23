<template>
    <Content v-if="$store.state.user.is_login === true">
        <div class="row" >
            <div class="col-md-3 col-12 mycol-left">
                <!-- 头像 写成一个子组件  -->
                <UserProfileInfo @changeNavbar="changeNavbar"  @follow="follow" @unfollow="unfollow"  v-bind:userinfo="user" :is_me="is_me"></UserProfileInfo>
            </div>

            <div class="col-md-9 col-12 mycol-right" >
                <!-- 写一个导航栏 -->
                <!-- @changePage="changePage"  -->
                <UserProfileNavbar  :is_me="is_me" @changeNavbar="changeNavbar" :navSelected="navSelected"></UserProfileNavbar>
                <!-- 下面写一个router切换渲染的东西：<router-view/> -->
                <!-- <router-view :is_followed="user.is_followed"/> -->
                <router-view v-if="navSelected === '3'" :is_followed="user.is_followed"/>
                <router-view v-else/>
                <!-- <UserPostLists @changePage="changePage" :posts="posts" :userinfo="user" :is_me="is_me"  v-if="navSelected === '0'"></UserPostLists>
                <UserProfileStar :is_me="is_me"  v-if="navSelected === '1'"></UserProfileStar>
                <UserProfileFollowList  v-if="navSelected === '2'"></UserProfileFollowList>
                <UserProfileFansList :is_followed="user.is_followed" v-if="navSelected === '3'"></UserProfileFansList> -->
            </div>
        </div>
        <CopyRight></CopyRight>
    </Content>
</template>

<script>
import { ref, reactive, watch, } from 'vue' //computed
import { useStore } from 'vuex'
import { useRoute } from 'vue-router';
import router from '@/router/index';   //@定位src目录
import Content from '../components/Content.vue'
import CopyRight from '../components/CopyRight.vue';
import UserProfileInfo from '../components/UserProfileInfo.vue'
// import UserPostLists from '../components/UserPostLists.vue'
import $ from 'jquery'
// import { ElMessage } from 'element-plus'
// import {FormatDateTime} from '../utils/DateTime'
// import ParseImageUrl from '../utils/ParseImageUrl'
import UserProfileNavbar from '../components/UserProfileNavbar.vue'
// import UserProfileStar from '../components/UserProfileStar.vue';
// import UserProfileFansList from '../components/UserProfileFansList.vue';
// import UserProfileFollowList from '../components/UserProfileFollowList.vue';

import {BackendRootURL} from '../common_resources/resource';
export default {
    name: "Profile",
    // ,UserProfileStar, UserProfileFansList, UserProfileFollowList, UserPostLists,
    components: {Content,UserProfileInfo,UserProfileNavbar, CopyRight},
    setup() {
        const store = useStore()
        // store.commit('updateUserPostlistPage', 1)
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
        // console.log(user_id)
        // let is_me = computed(() => user_id === store.state.user.id)
        const is_me = ref(user_id === store.state.user.id)
        // console.log(user_id, store.state.user.id, is_me.value)
        // 如果当前查看的userprofile不是我自己的话，user数据应该去数据库重新获取
        // 是自己的话就不重新获取
        const user = reactive({
            // ...store.state.user
        })

        // const page_size = ref(3)
        // const current_page = ref(1)
        // const posts = reactive({
        //     total_count:0,
        //     count:0,
        //     posts:[],
        // })
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

        // const getUserPosts = async () => {
        //     try {
        //         const resp = await $.ajax({
        //             url:  BackendRootURL + "/myspace/getuserposts/",
        //             type: "GET",
        //         data: {
        //             user_id: user_id,
        //             page_size: page_size.value,
        //             current_page: store.state.pagination.user_postlist_page,
        //         },
        //         headers: {
        //             'Authorization': "Bearer " + store.state.user.jwt,
        //         },
        //         });

        //         if (resp.status != 0) {
        //             ElMessage({
        //                 message: '获取用户帖子失败，请稍后重试',
        //                 type: 'error',
        //             });
        //             return;
        //         }
        //         // console.log("用户帖子：", resp.data)
        //         posts.posts = resp.data.data;  
        //         posts.total_count = resp.data.count
        //         posts.count = posts.posts.length;

        //         // console.log("my",posts)
        //         if(posts.posts === null || posts.count === null || posts.count === 0) {
        //             if(posts.total_count > 0) {
        //                 // console.log(11111)
        //                 store.commit('updateUserPostlistPage', store.state.pagination.user_postlist_page - 1)
        //                 getUserPosts()
        //                 return 
        //             }
        //         }

        //         for (let i = 0; i < posts.count; i++) {
        //         let post = posts.posts[i];
        //         let imgstr = post.image;
        //         posts.posts[i].image = ParseImageUrl(imgstr);
        //         posts.posts[i].CreatedAt = FormatDateTime(posts.posts[i].CreatedAt);
        //         }
        //     } catch (error) {
        //         ElMessage({
        //         message: '获取用户帖子失败，请稍后重试',
        //         type: 'error',
        //         });
        //         console.log(error);
        //     }
        // };
        // 我将获取用户信息和获取用户帖子的部分封装为异步函数getUserInfo和getUserPosts。然后，在loadData函数中使用await关键字按顺序执行这两个异步函数。
        const loadData = async () => {
            user_id = parseInt(route.params.userid);  //当前打开这个用户的id，从url上获取
            is_me.value = user_id === store.state.user.id;
            await getUserInfo();
            // await getUserPosts();
            // console.log(user)
            // 在这里可以继续操作userinfo和posts
        };

        // 监听路由参数的变化
        // 你可以使用watch()函数来监听路由参数的变化，并在参数变化时重新获取用户信息。
        // loadData();
        // console.log(router.currentRoute.value.params)  // 捕捉userid这个params

        watch(() => router.currentRoute.value.params, () => {
            // console.log(posts.count, posts.total_count)
            // if((posts.posts === null || posts.count === null || posts.count === 0) && posts.total_count > 0) {
            //     store.commit('updateUserPostlistPage', store.state.pagination.user_postlist_page - 1)
            //     getUserPosts()
            //     return 
            // }
            loadData();
            
        }, { immediate: true }); // immediate: true 表示在组件初始化时立即执行

        // loadData();
        
        const follow = () => {
            // console.log(555555)
            if (user.is_followed) return;
            // console.log(55555555555555555555)
            user.is_followed = true;
            user.fanscount ++ ;
        };

        const unfollow = () => {
            // console.log(666666)
            if (!user.is_followed) return;
            user.is_followed = false;
            user.fanscount -- ;
        };

        const nav_map = reactive({
            "ProfilePostLists":'0',
            'ProfileStar':'1',
            'ProfileFollowList':'2',
            'ProfileFanList':'3',
        })

        // 声明 navSelected 为响应式引用
        const navSelected = ref(nav_map[route.name]);

        // 监督route.name如果发生变化就更新navSelected
        // 监视 route.name 的变化
        watch(() => route.name, (newName) => {   // 刷新时，navSelected的值应该匹配url
            navSelected.value = nav_map[newName];
        });

        // const navSelected = ref(nav_map[route.name])
        // console.log(route.name)
        // if(route.name)

        const changeNavbar = (navSelect)=>{  //子组件往父组件传值
            navSelected.value = navSelect
        }

        // const changePage = () =>{
        //     getUserPosts()
        // }

        

        return {
            // store,
            // prefix,
            user,
            is_me,
            // userinfo,
            // posts,
            navSelected,
            follow,
            unfollow,
            changeNavbar,
            // page_size,
            // current_page,
            // changePage,
        }
    }
}

</script>


<style scoped>
.mycol-left{
    /* style="background-color: #F7F8FA; " */
    /* background-color: transparent; */
    /* position: relative; */
    z-index: 4;
}

.mycol-right {
    /* background-color: transparent; */
    /* position: relative; */
    z-index: 4;
}
</style>