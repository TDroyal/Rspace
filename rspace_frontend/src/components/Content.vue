<template>
    <div class="mycontainer">
        <div class="background"></div>
        <div class="container">
            <div class="card">
                <!-- /* 完全透明的白色 */ -->
                <!-- <div class="card-body" style="background-color: rgba(255, 255, 255, 0);"> -->
                    <!-- 灰色 -->
                <div class="card-body" >
                    <slot></slot>
                    <!-- 回到顶部 -->
                    <el-backtop :right="100" :bottom="100" style="width: 60px; height: 60px;"/>
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

.mycontainer{
    
    /* position: absolute; */
    
    /* background-size: 100vw 50vh;
    z-index: 10000; */
    /* position: fixed; */

    /* background-image: url('../assets/background.jpeg');
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: -1; */
    position: relative;
    overflow: hidden;
}

.background {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background-image: url('../assets/background.jpeg');
    background-size: cover;
    z-index: 2;
}


.container {
    margin-top: 60px;
    background-color: #F7F8FA;
    /* background-image: url('https://img0.baidu.com/it/u=2705025397,342285900&fm=253&fmt=auto&app=138&f=JPEG?w=973&h=500'); */
    /* 完全透明的白色 */
    /* background-color: rgba(255, 255, 255, 0);   */
    /* max-width: none; */
    max-height: none;
    min-height: 93.5vh;
    position: relative;
    background-color: transparent;  /* 设置为透明 */
    z-index: 3;  /*比.background大，并且是position: relative; 故在它的上面显示 */
}

.card {
    border: none;
    background-color: transparent;
    /* z-index: -2; */
}

.card-body {
    /* 完全透明的白色 */
    /* background-color: rgba(255, 255, 255, 0);   */
    /* background-image: url('https://img0.baidu.com/it/u=2705025397,342285900&fm=253&fmt=auto&app=138&f=JPEG?w=973&h=500'); */
    /* background-color: #F7F8FA; */
    background-color: transparent;
    /* z-index: -2; */
}

/* 使用 CSS 的 position 属性将 slot 的内容定位到最上层，例如设置 position: relative; 并且使用 z-index 将其置于最上层。 */
.slot-content {
  position: relative;
  z-index: 100000;
}

</style>