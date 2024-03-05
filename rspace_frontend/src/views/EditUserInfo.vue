<template>
    <Content>
        <div class="card card-out">
            <div class="card-body">
                <div class="row justify-content-center">
                    <div class="col-2">
                        <img class="img-fluid my-avatar" :src="$store.state.user.avatar" alt="" >
                        <span class="avatar-text">更换头像</span>
                        <!-- <div class="text-center username">{{$store.state.user.username}}</div> -->
                    </div>
                </div>
                <form @submit.prevent="modify_userinfo" style="margin-top: 20px;">
                    <div class="row justify-content-center" >
                        <div class="col-5">
                            <div class="mb-3">
                                <label for="username" class="form-label">昵称</label>
                                <input type="text" class="form-control" id="username" placeholder="你的昵称" v-model="userinfo.username">
                            </div>
                        </div>
                        <div class="col-5">
                            <div class="mb-3">
                                <label for="exampleFormControlInput1" class="form-label">性别</label>
                                <input type="text" class="form-control" id="exampleFormControlInput1" placeholder="请选择你的性别" v-model="userinfo.gender" >
                            </div>
                        </div>
                        <div class="col-5">
                            <div class="mb-3">
                                <label for="exampleFormControlInput1" class="form-label">年龄</label>
                                <input type="number" class="form-control" id="exampleFormControlInput1" placeholder="你的年龄" v-model="userinfo.age">
                            </div>
                        </div>
                        <div class="col-5">
                            <div class="mb-3">
                                <label for="exampleFormControlInput1" class="form-label">地址</label>
                                <input type="text" class="form-control" id="exampleFormControlInput1" placeholder="你所在地" v-model="userinfo.address">
                            </div>
                        </div>
                        <div class="col-10">
                            <div class="mb-3">
                                <label for="exampleFormControlTextarea1" class="form-label">个人介绍</label>
                                <textarea class="form-control" id="exampleFormControlTextarea1" rows="3" placeholder="关于你的个性、兴趣..." v-model="userinfo.introduction"></textarea>
                            </div>
                        </div>
                    </div>
                    <div class="row justify-content-center" style="margin-top: 20px;">
                        <div class="col-3 d-flex">
                            <button type="submit" class="btn position-relative">保存</button>
                        </div>    
                    </div>
                </form>
                
            </div>
        </div>
    </Content>
</template>

<script>
import { reactive } from 'vue';
import Content from '../components/Content.vue'
import {useStore} from 'vuex'

export default {
    name:"EditUserInfo",
    components:{Content},
    setup(){
        const store = useStore()

        const userinfo = reactive({
            ...store.state.user
        })
        const modify_userinfo = ()=>{
            console.log(userinfo)
        }
        return {
            // store,
            userinfo,
            modify_userinfo,
        }
    },
}

</script>

<style scoped>

.card-out {
    margin-top: 20px;
    /* box-shadow: none; */
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
    border: none;
}

.my-avatar{
    border-radius: 50%;
}

.my-avatar:hover{
    cursor: pointer;
}

.username{
    margin-top:5px;
    font-weight: bold;
}

button{
    width: 80%;
    margin: auto;
    background-color: #EFF9F2;
    color: #2DB55D;
    border: none;
}

button:hover {
    background-color: #92e4aa;
    color: white;
}

input {
    background-color: #F7F7F7;
}

textarea {
    background-color: #F7F7F7;
}


/* 在下面的代码中，我们添加了一个<span>元素作为my-avatar的兄弟元素，
    并设置其内容为"更换头像"。然后，我们使用CSS样式来控制avatar-text的显示和位置。
    初始状态下，我们将avatar-text的display属性设置为none，以隐藏它。
    当鼠标悬停在my-avatar上时，我们使用兄弟选择器(+)选择到avatar-text，
    并将其display属性设置为block，以显示它。 */
.avatar-text {
    /* display不显示 */
    display: none;
    position: absolute;
    top: 20%;
    left: 50%;
    transform: translate(-50%, -50%);
    background-color: rgba(0, 0, 0, 0.8);
    color: white;
    padding: 6px 12px;
    border-radius: 4px;
}

.my-avatar:hover + .avatar-text {
    display: block;
}

</style>