<template>
    <Content>
        <div class="row">
            <div class="col-3" style="background-color: #F7F8FA;">
                <!-- 头像 写成一个子组件 -->
                <UserProfileInfo v-bind:userinfo="user"></UserProfileInfo>
            </div>

            <div class="col-9">
                <UserPostLists :posts="posts" :userinfo="user"></UserPostLists>
            </div>
        </div>
    </Content>
</template>

<script>
import { reactive } from 'vue'
import { useStore } from 'vuex'
import Content from '../components/Content.vue'
import UserProfileInfo from '../components/UserProfileInfo.vue'
import UserPostLists from '../components/UserPostLists.vue'

export default {
    name: "Profile",
    components: {Content,UserProfileInfo,UserPostLists,},
    setup() {
        const store = useStore()
        const user = reactive({
            ...store.state.user
        })


        //作品应该从后端读取
        const posts = reactive({
            count:3,
            posts:[
                {
                    id:1,
                    userId:store.state.user.id,
                    content:"今天很开心,今天很开心,今天很开心,今天很开心今天很开心今天很开心今天很开心今天很开心今天很开心今天很开心今天很开心今天很开心",
                    image:"https://img0.baidu.com/it/u=1949733848,3040104377&fm=253&fmt=auto&app=120&f=JPEG?w=1280&h=800",
                },
                {
                    id:2,
                    userId:store.state.user.id,
                    content:"我的Share很有趣",
                    image:"https://img2.baidu.com/it/u=3644231109,3907585499&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=500",
                },
                {
                    id:3,
                    userId:store.state.user.id,
                    content:"Go是最好的语言",
                    image:"https://img1.baidu.com/it/u=2714579055,4205218818&fm=253&fmt=auto&app=138&f=JPEG?w=500&h=313",   
                }
            ]
        });
        // console.log(user)

        return {
            // store,
            user,
            posts,
        }
    }
}

</script>


<style scoped>

</style>