<template>
    <div class="horizontal-line"></div>
    <div class="row row-font">
        <div class="col-md-10 col-12 copyright">
            © 2024-2024 Rspace 版权所有 &nbsp; | &nbsp; 
            <span>
                <a href="http://beian.miit.gov.cn/publish/query/indexFirst.action" target="_black">蜀ICP备2024072244号</a>
            </span>
        </div>
        <div class="col-md-2 col-12 connection">
            <!-- 联系我们 &nbsp; | &nbsp;  -->
            <span>
                <a href="mailto:royalwangchao@gmail.com" target="_black">联系我们</a>
            </span>
        </div>
    </div>
    
</template>


<script>

export default{
    name:"CopyRight",
    components:{},
    setup(){

    },
}
</script>


<style scoped>

.horizontal-line {
    margin-top: 20px;
    border-bottom: 1px solid lightgray;
    margin-bottom: 5px;
}

.connection{
    display: flex;
    justify-content: right;
    color: #8c8c8c;
}

.copyright{
    color: #8c8c8c;
}

a {
    text-decoration: none;
}

/* .row-font{
    font-size: 100%;
} */

</style>