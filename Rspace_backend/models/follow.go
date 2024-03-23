package models

import (
	"time"
)

// 关注表  （多对多）
// 一个用户有多个粉丝，一个用户可以关注多个用户

type Follow struct {
	ID               uint
	CreatedAt        time.Time // Automatically managed by GORM for creation time
	UpdatedAt        time.Time // Automatically managed by GORM for update time     后续自己的收藏空间显示帖子按照这个时间降序排序，最新收藏时间排序
	IsFollowedUserID uint      `gorm:"column:isfollowed_user_id"`
	FollowedUserID   uint      `gorm:"column:followed_user_id"`
	Status           *int      `gorm:"column:status; type:tinyint(2); default:1;"`
}

// （用户点击关注时需要先判断之前是否关注过，根据IsFollowedUserID和FollowedUserID去数据库找，如果存在，就改Status=1，否则插入新的数据即可）
//用户第一次点击关注时，插入新的数据，status默认为1，关注成功
//用户点取消关注时，直接修改Status为0即可。
//用户再次点击关注时，直接修改Status为1即可。
