package models

import "time"

// 收藏关系表
// 一个用户有多个收藏   一个帖子也有多个收藏

type Collection struct {
	ID        uint
	CreatedAt time.Time // Automatically managed by GORM for creation time
	UpdatedAt time.Time // Automatically managed by GORM for update time     后续自己的收藏空间显示帖子按照这个时间降序排序，最新收藏时间排序
	PostID    uint      `gorm:"column:post_id;"`
	UserID    uint      `gorm:"column:user_id;"`
	Status    *int      `gorm:"column:status; type:tinyint(3); default:1;"`
}

// （用户点击收藏时需要先判断之前是否收藏过，根据PostID和UserID去数据库找，如果存在，就改Status=1，否则插入新的数据即可）
//用户第一次点击时收藏，插入新的数据，status默认为1，则红星
//用户取消收藏时，直接修改Status为0即可。
//用户再次点击收藏时，直接修改Status为1即可。
