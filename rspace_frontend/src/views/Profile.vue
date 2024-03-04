<template>
    <Content>
        <div class="row">
            <div class="col-3" style="background-color: #F7F8FA;">
                <!-- 头像 写成一个子组件 -->
                <UserProfileInfo v-bind:userinfo="user"></UserProfileInfo>
            </div>

            <div class="col-9">
                body
            </div>
        </div>
    </Content>
</template>

<script>
import { reactive } from 'vue'
import { useStore } from 'vuex'
import Content from '../components/Content.vue'
import UserProfileInfo from '../components/UserProfileInfo.vue'

export default {
    name: "Profile",
    components: {Content,UserProfileInfo,},
    setup() {
        const store = useStore()
        const user = reactive({
            ...store.state.user
        })
        console.log(user)

        return {
            store,
            user,
        }
    }
}

</script>


<style scoped>

</style>