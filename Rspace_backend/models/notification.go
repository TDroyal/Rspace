package models

import (
	"time"
)

type Notification struct {
	ID                 uint
	CreatedAt          time.Time // Automatically managed by GORM for creation time
	UpdatedAt          time.Time // Automatically managed by GORM for update time     后续自己的收藏空间显示帖子按照这个时间降序排序，最新收藏时间排序
	GenerateUserID     uint      `gorm:"column:generate_user_id"`
	ReceiveUserID      uint      `gorm:"column:receive_user_id"`
	PostID             uint      `gorm:"column:post_id"`
	NotificationStatus *int      `gorm:"column:notification_status; type:tinyint(3); default:1;"` //通知的状态，默认为1   0：消息已被ReceiveUserID读 1：消息未读
	NotificationType   int       `gorm:"column:notification_type; type:tinyint(3);"`              //1是评论 2是点赞
}
