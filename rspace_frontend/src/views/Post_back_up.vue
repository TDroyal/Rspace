<template>
    <Content>
        <div class="row justify-content-center">
            <div class="col-9">
                <div class="card card-single">
                    <div class="card-body">
                        <div class="row justify-content-center">
                            <div class="col-9">
                                <div class="post-title">发帖</div>
                            </div>
                            <div class="col-3">
                                <!-- <label for="post_type" class="form-label"></label> -->
                                <select class="form-select form-select-md" aria-label=".form-select-md example" style="font-weight: bold;" id="post_type" v-model="content_type">
                                    <option value="0" selected>帖子类型</option>
                                    <option value="1">日常</option>
                                    <option value="2">新鲜事</option>
                                    <option value="3">笔记</option>
                                    <option value="4">其它</option>
                                </select>
                            </div>
                        </div>
                        <br/>
                        <div class="horizontal-line"></div>
                        <!-- 内容，图片，帖子类型 -->

                        <div class="row" style="margin-top: 20px;">
                            <div class="col-12">
                                <div class="mb-3" >
                                    <label for="content" class="form-label">内容</label>
                                    <textarea class="form-control" id="content" rows="5" placeholder="分享点啥~" v-model="contents" maxlength="200"></textarea>
                                    <div class="left-words-count text-end">{{ remainingCharacters }} 字剩余</div>
                                </div>
                            </div>
                        </div>

                        <!-- <template>
                            <el-upload
                                ref="uploadRef"
                                class="upload-demo"
                                action="https://run.mocky.io/v3/9d059bf9-4660-45f2-925d-ce80ad6c4d15"
                                :auto-upload="false"
                            >
                                <template #trigger>
                                    <el-button type="primary">select file</el-button>
                                </template>

                                <el-button class="ml-3" type="success" @click="submitUpload">
                                    upload to server
                                </el-button>

                                <template #tip>
                                    <div class="el-upload__tip">
                                        jpg/png files with a size less than 500kb
                                    </div>
                                </template>
                            </el-upload>
                        </template> -->
                        <!-- action必选参数，上传的地址 -->
                        <div class="row">
                            <div style="font-weight: bold;">图片</div>
                            <div class="col-12">
                                <el-upload
                                    :limit="6"
                                    class="upload-demo"
                                    ref="uploadRef"
                                    action="http://127.0.0.1:9090/post/uploadimage/"  
                                    :on-preview="handlePreview"
                                    :on-remove="handleRemove"
                                    :headers="headers"
                                    v-model:file-list="fileList"
                                    :before-upload="beforeImageUpload"
                                    :on-success="handleSuccess"
                                    :on-error="handleError"
                                    :data="contentinfo"
                                    :auto-upload="false">
                                    <el-button  size="large" type="primary" style="width: 100%;">选取文件</el-button>
                                </el-upload>
                            </div>
                            <div class="col-12">
                                <div class="el-upload__tip">只能上传jpg/png/jpeg文件，且不超过2MB</div>
                            </div>
                        </div>
                        <!-- <el-button size="large" type="success" @click="submitUpload">上传到服务器</el-button> -->
                        <div class="row">
                            <div class="col-12">
                                <div class="d-flex justify-content-center">
                                    <button type="button" class="btn btn-primary position-relative" @click="submitUpload">点击分享</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </Content>
</template>

<script>
import Content from '../components/Content.vue'
import { reactive, ref, computed } from 'vue'
import { useStore} from 'vuex'
import { ElMessage} from 'element-plus'
// import $ from 'jquery'

export default {
    name: "Post",
    components:{Content},

    setup() {
        const uploadRef = ref(null);
        const fileList = ref([]);
        const store = useStore()
        const headers = reactive({
            'Authorization': 'Bearer ' + store.state.user.jwt
        })
        const remainingCharacters = computed(() => {
            return 200 - contents.value.length;
        });
        const contents = ref('')
        // let content_type = $('#post_type option:selected').val()
        const content_type = ref('0');
        const contentinfo = reactive({})
        const handleSuccess = (response) => {
            console.log("上传成功:", response);
            if(response.status != 0) {
                ElMessage({
                    message: '图片上传失败！！！',
                    type: 'error',
                });
                return 
            }
            ElMessage({
                message: '图片上传成功！！！',
                type: 'success',
            });
        };

        // 上传失败的回调函数
        const handleError = (error) => {
            console.error("上传失败:", error);
            ElMessage({
                message: '图片上传失败！！！',
                type: 'error',
            });
        };

        const submitUpload = () => {
            if(content_type.value === '0') {
                ElMessage({
                    message: '请选择帖子类型',
                    type: 'warning',
                })
                return 
            }
            if(contents.value === '') {
                ElMessage({
                    message: '帖子内容不能为空',
                    type: 'warning',
                })
                return 
            }
            // console.log(content_type.value, contents.value)
            contentinfo.content = contents.value
            contentinfo.type = content_type.value
            console.log(fileList.value)
            uploadRef.value.submit();
            
            // 这个只传content和content_type

            // $.ajax({
            //     url:"http://127.0.0.1:9090/post/uploadcontent/",
            //     type:"POST",
            //     data: {
            //         content:contents.value,
            //         type:content_type.value,
            //     },
            //     headers:{
            //         'Authorization': 'Bearer ' + store.state.user.jwt
            //     },
            //     success(resp) {
            //         // console.log('........................'+resp.status)
            //         if(resp.status != 0) {
            //             ElMessage({
            //                 message: '分享失败！！！',
            //                 type: 'error',
            //             })
            //             return
            //         }
            //         return
            //     },
            //     error(resp) {
            //         console.log(resp)
            //     }
            // })

            // ElMessage({
            //     message: '分享成功！！！',
            //     type: 'success',
            // })
            // console.log(content_type.value, contents.value)
        };

        const handleRemove = (file, fileList) => {
            console.log(file, fileList);
        };

        const handlePreview = (file) => {
            console.log(file);
        };

        const beforeImageUpload = (rawFile) => {
            // console.log(rawFile.type, rawFile.size / 1024/1024)
            if (rawFile.type !== 'image/jpg' && rawFile.type !== 'image/png' && rawFile.type !== 'image/jpeg') {
                ElMessage.error('Image must be JPG or PNG or JPEG format!')
                return false
            } else if (rawFile.size / 1024 / 1024 > 2) {
                ElMessage.error('Each image size can not exceed 2MB!')
                return false
            }
            fileList.value.push(rawFile);  // Store selected files in the fileList array
            return true
        }


        return {
            contents,
            remainingCharacters,
            content_type,
            uploadRef,
            fileList,
            headers,
            contentinfo,
            submitUpload,
            handleRemove,
            handlePreview,
            beforeImageUpload,
            handleSuccess,
            handleError,
        };
    }   
}

</script>


<style scoped>
.card-single {
    margin-top: 20px;
    /* box-shadow: none; */
    box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray;
    border: none;
    
}

.horizontal-line {
  border-bottom: 1px solid lightgray;
}

.post-title{
    
    font-weight: bold;
}

.form-label{
    font-weight: bold;
}

button{
    width: 30%;
    margin: auto;
}

.left-words-count{
    left: 100px;
    color: gray;
}

</style>