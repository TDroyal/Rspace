package models

// 普通用户表

type NormalUser struct {
	ID           uint   `gorm:"column:id" form:"id" json:"id" xml:"id"`                                  //默认是主键   同时也是参考User表的外键
	Name         string `gorm:"column:name;type:varchar(50);unique;" form:"name" json:"name" xml:"name"` //昵称唯一  默认是账号名username
	Gender       int    `gorm:"column:gender;type:tinyint(2);" form:"gender" json:"gender" xml:"gender"`
	Age          int    `gorm:"column:age;type:tinyint(10);" form:"age" json:"age" xml:"age"`
	Avatar       string `gorm:"column:avatar;type:varchar(300);" form:"avatar" json:"avatar" xml:"avatar"` //默认是一个空头像  （可以在后端放一个空头像，大家进来默认这个url）
	Address      string `gorm:"column:address;type:varchar(100);" form:"address" json:"address" xml:"address"`
	Introduction string `gorm:"column:introduction;type:varchar(500);" form:"introduction" json:"introduction" xml:"introduction"`
	// User         User   `gorm:"foreignKey:ID;references:ID"` // GORM同时提供自定义外键名字的方式  参考User表的ID
}
