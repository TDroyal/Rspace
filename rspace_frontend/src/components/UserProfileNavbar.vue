<template>
    <div class="card card-out" style="background-color:transparent; z-index: 3;">
        <div class="card-body" style="padding-bottom: 0px;">
            <div class="card">
                <div class="card-body" style="padding: 0px 16px;">
                    <div class="row justify-content-center">
                        <div class="col-3 mynav" :class="{ 'active': activeNav === '0' }" @click="setActiveNav('0')">
                            动态
                        </div>
                        <div class="col-3 mynav" :class="{ 'active': activeNav === '1' }" @click="setActiveNav('1')">
                            收藏
                        </div>
                        <div class="col-3 mynav" v-if="is_me === true" :class="{ 'active': activeNav === '2' }" @click="setActiveNav('2')">
                            我的关注
                        </div>
                        <div class="col-3 mynav" v-else :class="{ 'active': activeNav === '2' }" @click="setActiveNav('2')">
                            Ta的关注
                        </div>
                        <div class="col-3 mynav" v-if="is_me === true" :class="{ 'active': activeNav === '3' }" @click="setActiveNav('3')">
                            我的粉丝
                        </div>
                        <div class="col-3 mynav" v-else :class="{ 'active': activeNav === '3' }" @click="setActiveNav('3')">
                            Ta的粉丝
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import { watchEffect, ref} from 'vue' //computed
import { useStore } from 'vuex'
import router from '../router'
import { useRoute } from 'vue-router'
export default {
    name:"UserProfileNavbar",
    props:{
        is_me:{
            type:Boolean,
            requried:true,
        },
        navSelected:{
            type:String,
            requried:true,
        },
    },
    setup(props, context)
    {
        const store = useStore()
        const route = useRoute()
        const activeNav = ref(props.navSelected)  //选中的导航栏，传给父组件，父组件根据不同的导航栏显示不同的内容。

        // 如果希望 activeNav 随着父组件传递的 navSelected 改变而更新，可以使用 watchEffect 监听 props.navSelected 的变化，并在变化时更新 activeNav 的值。
        // 我们使用了 watchEffect 来监听 props.navSelected 的变化。当 props.navSelected 的值发生变化时，
        // watchEffect 的回调函数会被触发，然后我们将 activeNav 的值更新为最新的 props.navSelected 的值。
        watchEffect(() => {
            activeNav.value = props.navSelected;
        });
        
        const setActiveNav = (nav)=>{
            activeNav.value = nav
            if(activeNav.value === '0') {
                store.commit("updateUserPostlistPage", 1)
                // context.emit("changePage")
            }
            context.emit('changeNavbar', activeNav.value) // 子组件向父组件传递数据可以通过自定义事件来实现。
            if(nav === '0') {
                router.push({
                    name:"ProfilePostLists",  // ProfilePostLists
                    params:{
                        userid: parseInt(route.params.userid)
                    }
                })
            }else if(nav === '1') {
                router.push({
                    name:"ProfileStar",
                    params:{
                        userid: parseInt(route.params.userid)
                    }
                })
            }else if(nav === '2') {
                router.push({
                    name:"ProfileFollowList",
                    params:{
                        userid: parseInt(route.params.userid)
                    }
                })
            }else if(nav === '3') {
                router.push({
                    name:"ProfileFanList",
                    params:{
                        userid: parseInt(route.params.userid)
                    }
                })
            }
        }

        return {
            activeNav,
            setActiveNav,
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

.mynav {
    /* text-align: center; */
    display: flex;
    align-items: center; /* 垂直居中 */
    justify-content: center; /* 水平居中 */
    
    font-size: 20px;
    padding: 6px;
    color: black;
}

.mynav:hover {
    cursor: pointer;
    background-color: lightgray;
}

.active {
    color: #337AB7;
    
    box-shadow: 2px 0px 3px lightgray,  -2px 0px 3px lightgray;
}
</style>