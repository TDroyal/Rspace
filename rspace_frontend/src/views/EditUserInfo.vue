<template>
    <Content>
        <div class="card card-out">
            <div class="card-body">
                <div class="row justify-content-center">
                    <div class="col-2">
                        <img class="img-fluid my-avatar" :src="$store.state.user.avatar" alt="">
                        <span class="avatar-text">更换头像</span>
                        <!-- <div class="text-center username">{{$store.state.user.username}}</div> -->
                    </div>
                </div>
            </div>
        </div>
    </Content>
</template>

<script>
import Content from '../components/Content.vue'
import {useStore} from 'vuex'

export default {
    name:"EditUserInfo",
    components:{Content},
    setup(){
        const store = useStore()

        return {
            store,
        }
    },
}

</script>

<style scoped>

.card-out {
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

/* 在下面的代码中，我们添加了一个<span>元素作为my-avatar的兄弟元素，
    并设置其内容为"更换头像"。然后，我们使用CSS样式来控制avatar-text的显示和位置。
    初始状态下，我们将avatar-text的display属性设置为none，以隐藏它。
    当鼠标悬停在my-avatar上时，我们使用兄弟选择器(+)选择到avatar-text，
    并将其display属性设置为block，以显示它。 */
.avatar-text {
    /* display不显示 */
    display: none;
    position: absolute;
    top: 50%;
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