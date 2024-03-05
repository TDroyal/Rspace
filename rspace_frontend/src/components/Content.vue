<template>
    <div class="container">
        <div class="card">
            <div class="card-body" style="background-color: #F7F8FA;">
                <slot></slot>
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
.container {
    margin-top: 10px;
    background-color: #F7F8FA;
    /* height: 92.8vh; */
    max-height: none;
    min-height: 92.8vh;
}

.card {
    border: none;
}


</style>