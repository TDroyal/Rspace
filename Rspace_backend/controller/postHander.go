package controller

import (
	"Rspace_backend/dao"
	"Rspace_backend/middleware"
	"Rspace_backend/models"
	"fmt"
	"net/http"
	"strconv"
	"time"

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

type GetComment struct {
	Comment string `form:"comment" json:"comment" xml:"comment"  binding:"required"`
	PostID  uint   `form:"post_id" json:"post_id" xml:"post_id" binding:"required"`
	UserID  uint   `form:"user_id" json:"user_id" xml:"user_id" binding:"required"`
}

func UploadComment(c *gin.Context) {
	// 先得到用户的id
	claims := c.MustGet("claims").(*middleware.Myclaims)
	id := claims.ID //得到用户id
	fmt.Println(id)

	var get_comment GetComment

	if err := c.ShouldBind(&get_comment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	// 再存入数据库
	comment := models.Comment{Comment: get_comment.Comment, PostID: get_comment.PostID, UserID: get_comment.UserID}

	if err := dao.DB.Create(&comment).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "create to db error",
			"data":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "comment success",
		"data":    nil,
	})

	// fmt.Printf("%#v-----------\n", get_comment)
}

// 数据库查询需要返回的评论数据
type PostCommentData struct {
	ID        uint      `form:"id" json:"id" xml:"id"` //评论id
	CreatedAt time.Time //评论创建时间
	Comment   string    `form:"comment" json:"comment" xml:"comment"`
	UserID    uint      `form:"user_id" json:"user_id" xml:"user_id"`
	Name      string    `form:"name" json:"name" xml:"name"`
	Avatar    string    `form:"avatar" json:"avatar" xml:"avatar"`
}

func GetCommentsByPostId(c *gin.Context) {
	post_id := c.Query("post_id")
	// fmt.Printf("%#v-------------\n", post_id)

	var post_comment_data []PostCommentData

	dao.DB.Model(&models.Comment{}).
		Select("comments.id, comments.created_at, comments.comment, comments.user_id,normal_users.name ,normal_users.avatar").
		Where("comments.post_id = ?", post_id).
		Joins("left join normal_users on normal_users.id = comments.user_id").
		Order("comments.created_at desc").
		Scan(&post_comment_data)
	// fmt.Printf("%#v===============\n", post_comment_data)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "get posts success",
		"data":    post_comment_data,
	})
}

type GetCommendID struct {
	CommendID uint `form:"comment_id" json:"comment_id" xml:"comment_id" binding:"required"`
}

func DeleteCommentsByCommentId(c *gin.Context) {
	var id GetCommendID
	// delete前端传json过来，我才能获取到数据
	if err := c.BindJSON(&id); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "delete comment unsuccess",
			"data":    nil,
		})
		return
	}
	if err := dao.DB.Delete(&models.Comment{}, id.CommendID).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "delete comment unsuccess",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "delete comment success",
		"data":    nil,
	})
}
