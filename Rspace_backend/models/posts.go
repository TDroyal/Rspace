package models

import (
	"gorm.io/gorm"
)

// 帖子表

type Post struct {
	gorm.Model
	Content  string    `gorm:"column:content;type:varchar(500);" form:"content" json:"content" xml:"content"  binding:"required"`
	Image    *string   `gorm:"column:image;type:varchar(500);" form:"image" json:"image" xml:"image"  binding:"required"`     //图片可以为空以@#@分隔开
	Type     int       `gorm:"column:type;type:tinyint(3);default:1;" form:"type" json:"type" xml:"type"  binding:"required"` //默认是1share
	UserID   uint      `gorm:"column:user_id;" form:"user_id" json:"user_id" xml:"user_id"  binding:"required"`               //外键名
	Comments []Comment `gorm:"foreignKey:PostID;references:ID"`
	// 一个帖子有多个评论
	//一个帖子有多个用户收藏
	Collections []Collection `gorm:"foreignKey:PostID;references:ID"`
}
