<template>
    <!-- 非图片文件也能上传问题的bug还已解决，但是不能从列表中自动删除 -->
    <Content v-if="$store.state.user.is_login === true">
        <div class="row justify-content-center" >
            <div class="col-md-9 col-12">
                <div class="card card-single">
                    <div class="card-body">
                        <div class="row justify-content-center">
                            <div class="col-6 col-md-9">
                                <div class="post-title" >发帖</div>
                            </div>
                            <div class="col-6 col-md-3 ">
                                <!-- <label for="post_type" class="form-label"></label> -->
                                <!-- <select class="form-select form-select-md" aria-label=".form-select-md example" style="font-weight: bold;" id="post_type" v-model="content_type">
                                    <option value="0" selected>帖子类型</option>
                                    <option value="1">日常</option>
                                    <option value="2">新鲜事</option>
                                    <option value="3">笔记</option>
                                    <option value="4">其它</option>
                                </select> -->
                                <el-select
                                    v-model="content_type"
                                    placeholder="帖子类型"
                                    size="large"
                                    class="custom-select"
                                    style="font-weight: bold;"
                                >
                                    <el-option
                                        v-for="item in options"
                                        :key="item.value"
                                        :label="item.label"
                                        :value="item.value"
                                    />
                                </el-select>
                            </div>
                        </div>
                        <br/>
                        <div class="horizontal-line"></div>
                        <!-- 内容，图片，帖子类型 -->

                        <div class="row" style="margin-top: 20px;">
                            <div class="col-12">
                                <div class="mb-3" >
                                    <label for="content" class="form-label">内容</label>
                                    <textarea class="form-control" id="content" rows="10"   placeholder="分享点啥~" v-model="contents" maxlength="400"></textarea>
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
                        <!-- :http-request="uploadPosts"  :on-change="checkImage" :before-upload="beforeImageUpload"-->
                        <div class="row">
                            <div style="font-weight: bold;">图片</div>
                            <div class="col-12">
                                <el-upload
                                    :limit="6"
                                    class="upload-demo"
                                    action="#"  
                                    list-type="picture"
                                    :on-success="handleSuccess"
                                    :on-error="handleError"
                                    :on-exceed="handleExceed"
                                    :on-change="do_checkImage"
                                    :on-remove="handleRemove"
                                    :file-list="fileList"
                                    :auto-upload="false">
                                    <el-button size="large" type="primary" style="width: 100%;">选取图片</el-button>
                                </el-upload>
                            </div>
                            <div class="col-12">
                                <div class="el-upload__tip">只能上传jpg/png/jpeg文件，且不超过2MB</div>
                            </div>
                        </div>
                        <!-- <el-button size="large" type="success" @click="submitUpload">上传到服务器</el-button> -->
                        <div class="row" style="margin-top: 20px;">
                            <div class="col-12 col-md-4" style="margin: 0 auto;">
                                <div class="d-flex justify-content-center" >
                                    <button type="button" style="width: 100%;" class="btn btn-primary position-relative" @click="submitUpload">点击分享</button>
                                </div>
                            </div>
                        </div>
                    </div>
                </div>
                <CopyRight></CopyRight>
            </div>
        </div>
    </Content>
</template>

<script>
import Content from '../components/Content.vue'
import CopyRight from '../components/CopyRight.vue'
import { reactive, ref, computed } from 'vue'
import { useStore} from 'vuex'
import { ElMessage, ElLoading} from 'element-plus' // 
import router from '@/router/index';   //@定位src目录
import $ from 'jquery'
import {BackendRootURL} from '../common_resources/resource';
import { CheckIsLogin,  RefreshToken} from '../utils/MakeAuthenticatedRequest'
export default {
    name: "Post",
    components:{Content, CopyRight},

    setup() {

        const options = [
            {
                value: '1',
                label: '日常',
            },
            {
                value: '2',
                label: '新鲜事',
            },
            {
                value: '3',
                label: '笔记',
            },
            {
                value: '4',
                label: '其它',
            }
        ]
        // <option value="0" selected>帖子类型</option>
        //                             <option value="1">日常</option>
        //                             <option value="2">新鲜事</option>
        //                             <option value="3">笔记</option>
        //                             <option value="4">其它</option>

        const store = useStore()
        // const check_is_login = ()=>{
        //     if(store.state.user.is_login === false) {
        //         router.push({
        //             name:"Login",
        //         })
        //         return false
        //     }
        //     return true
        // }
        // if(check_is_login() === false) {
        //     return
        // }

        if(CheckIsLogin(store) === false) {
            return 
        }

        const uploadRef = ref(null);
        const fileList = ref([]);
        
        const headers = reactive({
            'Authorization': 'Bearer ' + store.state.user.jwt
        })
        const remainingCharacters = computed(() => {
            return 400 - contents.value.length;
        });
        const contents = ref('')
        // let content_type = $('#post_type option:selected').val()
        const content_type = ref('');
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

        // 手动实现上传图片的函数
        const submitUpload = () => {
            if(content_type.value === '') {
                ElMessage({
                    message: '请选择帖子类型',
                    type: 'warning',
                })
                return 
            }
            if(contents.value === '' || /^\s*$/.test(contents.value)) {  //使用正则表达式判断一个字符串中是否只包含空格
                ElMessage({
                    message: '帖子内容不能为空',
                    type: 'warning',
                })
                return 
            }
            // console.log(contents.value, content_type.value, fileList.value)
            
            const loading = ElLoading.service({
                lock: true,
                text: '帖子上传中',
                background: 'rgba(0, 0, 0, 0.7)',
            })
            
            // 得到帖子的相关信息：
            // console.log(contents.value, content_type.value, fileList.value)
            // const post = {
            //     images:[],
            //     content:null,
            //     type:null,
            // }
            // console.log("----------------", fileList.value)
            // return 
            
            const formData = new FormData();
            // console.log(fileList.value)
            fileList.value.forEach((file) => {
                // post.images.push(file.raw)
                formData.append('images', file.raw);
            });
            formData.append('content', contents.value)
            formData.append('type', content_type.value)
            // console.log(formData.getAll('images'), )
            // post.content = contents.value
            // post.type = parseInt(content_type.value)
            // console.log(post)

            // 要将多张照片数据传递到后端，你可以使用 FormData 对象来构建请求，并将每张照片添加到 FormData 对象中的 images 字段。
            // 在下述代码中，我们首先创建了一个新的 FormData 对象。然后，使用 forEach 循环遍历 fileList 数组，将每个文件的 raw 数据添加到 FormData 对象的 images 字段中。
            // 接下来，我们可以通过 formData.append 方法将其他字段（如 content 和 type）添加到 FormData 对象中。
            // 最后，我们使用 $.ajax 发送请求到后端。请注意，我们设置了 processData 和 contentType 选项为 false，以确保 FormData 对象以正确的方式发送。

            $.ajax({
                url: BackendRootURL + "/post/uploadimage/",
                type:"POST",
                data:formData,
                processData: false,
                contentType: false,
                headers:{
                    'Authorization': 'Bearer ' + store.state.user.jwt
                },
                success(resp) {
                    // console.log(resp)
                    if(resp.status != 0) {
                        if(resp.status === 401) {
                            RefreshToken(store)
                            .then((jwt) => {
                                if(jwt) {
                                    submitUpload()
                                }
                            })
                            .catch((error) => {
                                console.error(error);
                            });
                            return
                        }
                        ElMessage.error("分享失败")
                        return
                    }
                    ElMessage.success("分享成功!!!")
                    store.commit("updateHomeCurrentPage", 1)
                    //跳到首页去
                    router.push({
                        name:"Home",
                    })
                    loading.close()
                },
                error(resp) {
                    console.log(resp)
                    loading.close()
                }
            })
            // loading.close()
            // console.log(fileList.value)

        };

        // const handleRemove = (file, fileList) => {
        //     console.log(file, fileList);
        // };
        const handleRemove = (file) => {
            // console.log(file) // 根据图片名称找到第一个同名的
            const index = fileList.value.indexOf(file);
            if (index !== -1) {
                fileList.value.splice(index, 1);
            }
        }

        const handlePreview = (file) => {
            console.log(file);
        };

        const beforeImageUpload = (rawFile) => {
            // console.log(rawFile.type, rawFile.size / 1024/1024)
            if (rawFile.raw.type !== 'image/jpg' && rawFile.raw.type !== 'image/png' && rawFile.raw.type !== 'image/jpeg') {
                ElMessage.error('只能上传 JPG、PNG 或 JPEG 格式的图片！')
                return false
            } else if (rawFile.size / 1024 / 1024 > 2) {
                ElMessage.error('每张图片大小不能超过 2MB！')
                return false
            }
            // fileList.value.push(rawFile);  // Store selected files in the fileList array
            return true
        }

        const do_checkImage = (rawFile) =>{
            checkImage(rawFile)
            filter_pic(checkImage)
        }

        const checkImage = (rawFile) =>{
            // console.log(rawFile)
            if (rawFile.raw.type !== 'image/jpg' && rawFile.raw.type !== 'image/png' && rawFile.raw.type !== 'image/jpeg') {
                // 不是这个类型的就应该删掉该文件
                
                ElMessage.error('只能上传 JPG、PNG 或 JPEG 格式的图片！')
                handleRemove(rawFile)
                return false
            } else if (rawFile.size / 1024 / 1024 > 10) {  //出网带宽峰值小于等于10Mbit/s时，阿里云会分配10Mbit/s入网带宽   出网带宽峰值大于10Mbit/s时，阿里云会分配峰值相等的入网带宽

                ElMessage.error('每张图片大小不能超过 10MB！')
                handleRemove(rawFile)
                return false
            }

            // fileList.value.push(rawFile);  // Store selected files in the fileList array
            fileList.value.splice(fileList.value.length, 0, rawFile); // 添加文件到 fileList 数组末尾
            return true
        }

        const filter_pic = ()=>{
            const filteredFileList = fileList.value.filter(checkImage);
            fileList.value = filteredFileList;
        }

        const handleExceed = ()=>{
            ElMessage.warning("最多只能上传6张图片哦！")
        }

        const progress = (uploadFile)=>{
            console.log('-------------',uploadFile)
        }

        return {
            options,
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
            handleExceed,
            checkImage,
            progress,
            filter_pic,
            do_checkImage,
        };
    }   
}

</script>


<style scoped>
.card-single {
    margin-top: 20px;
    /* box-shadow: none; */
    /* box-shadow: 2px 2px 3px lightgray,  -2px 0 3px lightgray; */
    border: none;
    
}

.horizontal-line {
  border-bottom: 1px solid lightgray;
}

.post-title{
    font-weight: bold;
    font-size: 20px;
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

.upload-demo{
    width: 100%;
}

.custom-select .el-input__inner::placeholder {
    
    color: black;
}



</style>