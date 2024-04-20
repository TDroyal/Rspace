package controller

import (
	"Rspace_backend/dao"
	"Rspace_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 返回用户名称，用户头像，帖子发布时间，图片和内容，点赞量和评论量
type Search_Info struct {
	PageSize     int    `form:"page_size" json:"page_size" xml:"page_size" binding:"required"`
	CurrentPage  int    `form:"current_page" json:"current_page" xml:"current_page" binding:"required"`
	PostType     int    `form:"post_type" json:"post_type" xml:"post_type" binding:"required"`
	QueryContent string `form:"query_content" json:"query_content" xml:"query_content" binding:"required"`
}

// 搜索的帖子结果返回
type SearchPostResult struct {
	PostID          uint      `form:"post_id" json:"post_id" xml:"post_id"  binding:"required"`
	Name            string    `form:"name" json:"name" xml:"name"  binding:"required"`
	Avatar          string    `form:"avatar" json:"avatar" xml:"avatar"  binding:"required"`
	CreatedAt       time.Time `form:"created_at" json:"created_at" xml:"created_at"  binding:"required"`
	Image           *string   `form:"image" json:"image" xml:"image"  binding:"required"`
	Content         string    `form:"content" json:"content" xml:"content"  binding:"required"`
	CollectionCount int64     `form:"collection_count" json:"collection_count" xml:"collection_count"  binding:"required"`
	CommentCount    int64     `form:"comment_count" json:"comment_count" xml:"comment_count"  binding:"required"`
}

func GetPostList(c *gin.Context) {
	var search_info Search_Info
	if err := c.ShouldBind(&search_info); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "获取搜索信息失败",
			"data":    err.Error(),
		})
		return
	}

	var search_result []SearchPostResult

	//获取当前请求的帖子列表 (先按点赞降序排序，再按评论，最后按时间)
	if err := dao.DB.Model(&models.Post{}).Where("type = ? and posts.deleted_at is null and content like ?", search_info.PostType, "%"+search_info.QueryContent+"%").
		Order("posts.created_at desc").
		InnerJoins("join normal_users on normal_users.id = posts.user_id").
		Select("posts.id as post_id, normal_users.name, normal_users.avatar, posts.created_at, posts.image, posts.content").
		Limit(search_info.PageSize).
		Offset(search_info.PageSize * (search_info.CurrentPage - 1)).
		Scan(&search_result).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "获取搜索结果失败",
			"data":    err,
		})
		return
	}

	// 成功之后，又遍历每个搜索的结果，去查找其对应的点赞数量和
	n := len(search_result)
	for i := 0; i < n; i++ {
		// fmt.Println(search_result[i].PostID)
		if err := dao.DB.Model(&models.Collection{}).
			Where("post_id = ? and status = 1", search_result[i].PostID).Count(&search_result[i].CollectionCount).Error; err != nil {
			panic(err)
		}

		if err := dao.DB.Model(&models.Comment{}).
			Where("post_id = ? and deleted_at is null", search_result[i].PostID).Count(&search_result[i].CommentCount).Error; err != nil {
			panic(err)
		}
	}

	// fmt.Println(search_result)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "搜索成功",
		"data":    search_result,
	})
}

type Search_User_Info struct {
	PageSize     int    `form:"page_size" json:"page_size" xml:"page_size" binding:"required"`
	CurrentPage  int    `form:"current_page" json:"current_page" xml:"current_page" binding:"required"`
	QueryContent string `form:"query_content" json:"query_content" xml:"query_content" binding:"required"`
}

// 搜索的用户列表结果返回
type SearchUserResult struct {
	UserID    uint   `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Name      string `form:"name" json:"name" xml:"name"  binding:"required"`
	Avatar    string `form:"avatar" json:"avatar" xml:"avatar"  binding:"required"`
	FansCount int64  `form:"fanscount" json:"fanscount" xml:"fanscount"  binding:"required"`
}

func GetUserList(c *gin.Context) {
	var search_info Search_User_Info
	if err := c.ShouldBind(&search_info); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "获取搜索信息失败",
			"data":    err.Error(),
		})
		return
	}

	var search_result []SearchUserResult
	if err := dao.DB.Model(&models.NormalUser{}).Where("name like ?", "%"+search_info.QueryContent+"%").
		Select("id as user_id, name, avatar").
		Limit(search_info.PageSize).
		Offset(search_info.PageSize * (search_info.CurrentPage - 1)).
		Scan(&search_result).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "获取搜索结果失败",
			"data":    err,
		})
		return
	}

	n := len(search_result)
	for i := 0; i < n; i++ {
		if err := dao.DB.Model(&models.Follow{}).
			Where("isfollowed_user_id = ? and status = 1", search_result[i].UserID).Count(&search_result[i].FansCount).Error; err != nil {
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "搜索成功",
		"data":    search_result,
	})
}
