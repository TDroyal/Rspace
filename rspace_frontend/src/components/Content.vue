<template>
    <div class="mycontainer">
        <div class="container">
            <div class="card">
                <!-- /* 完全透明的白色 */ -->
                <!-- <div class="card-body" style="background-color: rgba(255, 255, 255, 0);"> -->
                    <!-- 灰色 -->
                <div class="card-body" >
                    <slot></slot>
                </div>
            </div>
        </div>
    </div>
</template>


<script>
export default {
    name:"Content",
}

</script>


<style scoped>

/* .mycontainer{
    background-image: url('https://img0.baidu.com/it/u=2705025397,342285900&fm=253&fmt=auto&app=138&f=JPEG?w=973&h=500');
} */

.container {
    margin-top: 60px;
    background-color: #F7F8FA;
    /* 完全透明的白色 */
    /* background-color: rgba(255, 255, 255, 0);   */
    /* max-width: none; */
    max-height: none;
    min-height: 93.5vh;
    
}

.card {
    border: none;
}

.card-body {
    /* 完全透明的白色 */
    /* background-color: rgba(255, 255, 255, 0);   */
    background-color: #F7F8FA;
}



</style>