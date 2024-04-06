package controller

import (
	"Rspace_backend/dao"
	"net/http"

	// "Rspace_backend/middleware"
	"Rspace_backend/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 获得最新的5个帖子，联表查询，将用户的name和id，avatar返回

// 2024-4-4根据分页来返回对应的帖子信息
type Pagination struct {
	PageSize    int `form:"page_size" json:"page_size" xml:"page_size" binding:"required"`
	CurrentPage int `form:"current_page" json:"current_page" xml:"current_page" binding:"required"`
}

type PostsWithUserInfo struct {
	gorm.Model
	Content string  `form:"content" json:"content" xml:"content"  binding:"required"`
	Image   *string `form:"image" json:"image" xml:"image"  binding:"required"`
	Type    int     `form:"type" json:"type" xml:"type"  binding:"required"`
	UserID  uint    `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Name    string  `form:"name" json:"name" xml:"name"  binding:"required"`
	Avatar  string  `form:"avatar" json:"avatar" xml:"avatar"`
}

func GetHomePost(c *gin.Context) {
	var result []PostsWithUserInfo

	//获取分页信息
	var pagination_info Pagination
	if err := c.ShouldBind(&pagination_info); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "获取分页信息失败",
			"data":    err.Error(),
		})
		return
	}
	//还需要返回总的post数量
	var count int64
	if err := dao.DB.Model(&models.Post{}).Where("deleted_at is NULL").Count(&count).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "获取帖子总数失败",
			"data":    err.Error(),
		})
		return
	}

	// 连表查询 查找第CurrentPage的PageSize
	// Limit指定要检索的最大记录数Offset指定开始返回记录之前要跳过的记录数
	if err := dao.DB.Model(&models.Post{}).Order("posts.created_at desc").Limit(pagination_info.PageSize).Offset(pagination_info.PageSize * (pagination_info.CurrentPage - 1)).Select("posts.id, posts.created_at, posts.updated_at, posts.deleted_at, posts.content, posts.image, posts.type, posts.user_id, normal_users.name, normal_users.avatar").Joins("JOIN normal_users ON normal_users.id = posts.user_id").Scan(&result).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "获取帖子数据失败",
			"data":    err.Error(),
		})
		return
	}
	// fmt.Printf("---------%#v------------", result)

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "get posts success",
		"data":    gin.H{"count": count, "data": result},
	})
}

// 根据帖子id获取帖子信息
func GetPostByPostId(c *gin.Context) {
	var result PostsWithUserInfo
	post_id := c.Query("postid")
	// fmt.Printf("%#v------\n", post_id)
	if err := dao.DB.Model(&models.Post{}).Where("posts.id = ?", post_id).Select("posts.id, posts.created_at, posts.updated_at, posts.deleted_at, posts.content, posts.image, posts.type, posts.user_id, normal_users.name, normal_users.avatar").Joins("JOIN normal_users ON normal_users.id = posts.user_id").Scan(&result).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "get post error",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "get post success",
		"data":    result,
	})
}
