package controller

import (
	"Rspace_backend/dao"
	"Rspace_backend/middleware"
	"Rspace_backend/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ANotification struct {
	GenerateUserID   uint `form:"generate_user_id" json:"generate_user_id" xml:"generate_user_id" binding:"required"`
	ReceiveUserID    uint `form:"receive_user_id" json:"receive_user_id" xml:"receive_user_id" binding:"required"`
	PostID           uint `form:"post_id" json:"post_id" xml:"post_id" binding:"required"`
	NotificationType int  `form:"notification_type" json:"notification_type" xml:"notification_type" binding:"required"`
}

// 发起通知
func PostNotification(c *gin.Context) {
	var notification ANotification
	if err := c.ShouldBind(&notification); err != nil {
		// c.JSON(http.StatusOK, gin.H{
		// 	"status":  -1,
		// 	"message": "获取通知消息失败",
		// 	"data":    err.Error(),
		// })
		panic(err)
	}

	// fmt.Println("-----------------------", notification, notification.PostID)
	// 先判断数据库有没有这个点赞记录，评论通知是评论一次通知一次，点赞和关注通知是初次点赞和关注才会通知
	if notification.NotificationType == 2 {
		var count int64
		if err := dao.DB.Model(&models.Notification{}).Where("generate_user_id = ? and receive_user_id = ? and post_id = ? and notification_type = ?", notification.GenerateUserID, notification.ReceiveUserID, notification.PostID, notification.NotificationType).Count(&count).Error; err != nil {
			panic(err)
		}
		// fmt.Println(count)
		//数据库中此条点赞通知已经存在
		if count >= 1 {
			return
		}
	} else if notification.NotificationType == 3 {
		var count int64
		if err := dao.DB.Model(&models.Notification{}).Where("generate_user_id = ? and receive_user_id = ? and notification_type = ?", notification.GenerateUserID, notification.ReceiveUserID, notification.NotificationType).Count(&count).Error; err != nil {
			panic(err)
		}
		// fmt.Println(count)
		//数据库中此条点赞通知已经存在
		if count >= 1 {
			return
		}
	}

	//将通知存入数据库
	var notification_info = models.Notification{
		GenerateUserID:   notification.GenerateUserID,
		ReceiveUserID:    notification.ReceiveUserID,
		PostID:           notification.PostID,
		NotificationType: notification.NotificationType,
	}
	if err := dao.DB.Create(&notification_info).Error; err != nil {
		panic(err)
	}
}

// 暂时获取所有通知(后续在分页)

type BackToNotifications struct { //返回给前端的通知
	ID                 uint      `form:"id" json:"id" xml:"id"  binding:"required"`
	UserID             uint      `form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`
	Name               string    `form:"name" json:"name" xml:"name"  binding:"required"`
	Avatar             string    `form:"avatar" json:"avatar" xml:"avatar"  binding:"required"`
	NotificationStatus *int      `form:"notification_status" json:"notification_status" xml:"notification_status"  binding:"required"`
	NotificationType   int       `form:"notification_type" json:"notification_type" xml:"notification_type"  binding:"required"`
	CreatedAt          time.Time `form:"created_at" json:"created_at" xml:"created_at"  binding:"required"`
	PostID             uint      `form:"post_id" json:"post_id" xml:"post_id"  binding:"required"`
	PostContent        string    `form:"post_content" json:"post_content" xml:"post_content"  binding:"required"`
}

// 分页
type PaginationInfo struct {
	PageSize    int `form:"page_size" json:"page_size" xml:"page_size" binding:"required"`
	CurrentPage int `form:"current_page" json:"current_page" xml:"current_page" binding:"required"`
}

func GetNotifications(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.Myclaims)
	id := claims.ID //得到当前用户id

	// 获得分页消息
	var pagination PaginationInfo
	if err := c.ShouldBind(&pagination); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "error",
			"data":    err.Error(),
		})
		return
	}

	// 先获得一共有多少条通知消息
	var count1, count2 int64
	if err := dao.DB.Model(&models.Notification{}).Where("receive_user_id = ? and notification_type = 3", id).Count(&count1).Error; err != nil {
		panic(err)
	}

	if err := dao.DB.Model(&models.Notification{}).Where("receive_user_id = ? and notification_type in (1, 2)", id).
		InnerJoins("join posts on posts.id = notifications.post_id and posts.deleted_at is null ").
		Count(&count2).Error; err != nil {
		panic(err)
	}

	var get_notifications []BackToNotifications
	if err := dao.DB.Model(&models.Notification{}).
		Where("receive_user_id = ?", id).
		Order("notifications.created_at desc").
		InnerJoins("join normal_users on normal_users.id = notifications.generate_user_id").
		InnerJoins("join posts on posts.id = notifications.post_id and posts.deleted_at is null or notifications.post_id = posts.id and notifications.notification_type = 3").
		Select("notifications.id, notifications.generate_user_id as user_id, normal_users.name,normal_users.avatar, notification_status, notification_type, notifications.created_at, notifications.post_id, posts.content as post_content").
		Limit(pagination.PageSize).
		Offset(pagination.PageSize * (pagination.CurrentPage - 1)).
		Scan(&get_notifications).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "获取通知失败",
			"data":    err,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "获得通知成功",
		"data": gin.H{
			"data":  get_notifications,
			"count": count1 + count2, // 总的数量
		},
	})
}

// 将通知设为已读，得到通知的id，根据通知的id将notification_status改为0
type PutNoticeStatus struct {
	NoticeID           uint `form:"notice_id" json:"notice_id" xml:"notice_id" binding:"required"`
	NotificationStatus *int `form:"notification_status" json:"notification_status" xml:"notification_status"  binding:"required"`
}

func ChangeNotificationStatus(c *gin.Context) {
	var changeInfo PutNoticeStatus
	if err := c.ShouldBind(&changeInfo); err != nil {
		panic(err)
	}
	if err := dao.DB.Model(&models.Notification{}).Where("id = ?", changeInfo.NoticeID).Update("notification_status", changeInfo.NotificationStatus).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "ChangeNotificationStatus success",
		"data":    nil,
	})

}

func GetUnreadNotificationsCount(c *gin.Context) {
	claims := c.MustGet("claims").(*middleware.Myclaims)
	id := claims.ID //得到当前用户id

	// 根据用户id查找未读的消息的条数
	// 关注的notification_type是3，对帖子的评论和点赞notification_type分别是1和2  1和2的记数需要关联post看此post有无删除
	var follow_notification_count int64
	if err := dao.DB.Model(&models.Notification{}).Where("receive_user_id = ? and notification_status = 1 and notification_type = 3", id).
		Count(&follow_notification_count).Error; err != nil {
		panic(err)
	}

	var post_notification_count int64
	if err := dao.DB.Model(&models.Notification{}).Where("receive_user_id = ? and notification_status = 1 and notification_type in (1, 2)", id).
		InnerJoins("join posts on posts.id = notifications.post_id and posts.deleted_at is null").Count(&post_notification_count).Error; err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "获得未读通知条数成功",
		"data":    follow_notification_count + post_notification_count,
	})
}
