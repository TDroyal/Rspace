package models

import (
	"gorm.io/gorm"
)

// 评论表
type Comment struct {
	gorm.Model
	Comment string `gorm:"column:Comment;type:varchar(400);"`
	PostID  uint   `gorm:"column:post_id;"`
	UserID  uint   `gorm:"column:user_id;"`
	// 我需要返回对应的发布评论的用户id和用户姓名，用户avatar
}
