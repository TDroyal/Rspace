package controller

import (
	"Rspace_backend/dao"
	"Rspace_backend/middleware"
	"Rspace_backend/models"
	"fmt"
	"net/http"
	"strconv"

	"os"

	"github.com/gin-gonic/gin"
)

// 获取帖子的内容
// Content，Type
type UploadPost struct {
	Content string `form:"content" json:"content" xml:"content"  binding:"required"`
	Type    int    `form:"type" json:"type" xml:"type"  binding:"required"`
}

func UploadPostContentHandler(c *gin.Context) {

	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// Multipart form

	// 先得到用户的id
	claims := c.MustGet("claims").(*middleware.Myclaims)
	id := claims.ID //得到用户id
	fmt.Println(id)

	// 首先得到内容和类型
	var uploadresult UploadPost
	if err := c.ShouldBind(&uploadresult); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	fmt.Printf("-----------%#v\n", uploadresult)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "upload post success",
		"data":    nil,
	})
}

func UploadPostImageHandler(c *gin.Context) {
	// 先得到用户的id
	claims := c.MustGet("claims").(*middleware.Myclaims)
	id := claims.ID //得到用户id
	// fmt.Println(id)

	form, _ := c.MultipartForm() //上传多个文件
	fmt.Printf("------------form=%#v\n", form)
	files := form.File["images"]
	content := form.Value["content"][0]
	content_type := form.Value["type"][0]
	content_type_int, _ := strconv.Atoi(content_type)

	var images []string

	fmt.Println(content, content_type, files, "***************")

	// 保存照片在static文件夹下，保存失败的次数
	err_cnt := 0

	for index, file := range files {
		// log.Println(file.Filename)
		// fmt.Printf("%#v---------------\n", file.Filename)
		img_name := fmt.Sprintf("rspace_%d_%s", index, file.Filename)
		images = append(images, img_name)
		dst := fmt.Sprintf("./static/posts/%s", img_name)
		// 上传文件至指定目录
		// fmt.Println("+++++++++++++++++++++", dst)  // ./static/posts/rspace_3_资质认证材料.jpeg
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			err_cnt++
			fmt.Println("==========================", err)
		}
		// fmt.Println(index, images)
	}

	// 如果没有完全保存到后端，那么我们将保存成功的图片先删除
	if err_cnt != 0 {

		// 先查找保存成功的图片，然后删掉
		// 使用 os.Stat 函数判断文件是否存在

		n := len(images)
		for i := 0; i < n; i++ {
			img_name := images[i]
			filename := fmt.Sprintf("./static/posts/%s", img_name)
			_, err := os.Stat(filename)
			if err == nil {
				// fmt.Println("文件存在")
				err_ := os.Remove(filename)
				if err_ != nil {
					// 删除文件失败
					// 处理错误
					fmt.Println("删除文件失败", err_)
				}
			} else if os.IsNotExist(err) {
				fmt.Println("文件不存在")
			} else {
				fmt.Println("发生错误:", err)
			}
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "upload images error",
			"data":    nil,
		})
		return
	}

	// 所有图片都完美保存
	// 1. 先将图片预处理
	n := len(images)
	var image_str string
	for i := 0; i < n; i++ {
		image_str += images[i] + "@#@"
	}
	// fmt.Printf("***************************%s\n", image_str)

	//接下来就存入数据库啦
	post := models.Post{Content: content, Image: &image_str, UserID: id, Type: content_type_int}
	if err := dao.DB.Create(&post).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "upload post error",
			"data":    nil,
		})
		return
	}
	// fmt.Println(images)

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "upload post success",
		"data":    nil,
	})
}
