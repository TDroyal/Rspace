package models

// 创建用户结构体模型
// import (
// 	"Rspace_backend/dao"
// )

// 绑定 JSON
type Login struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

type User struct {
	// gorm.Model
	ID       uint   `gorm:"column:id"`
	UserName string `gorm:"column:username;type:varchar(50);not null;unique;" form:"username" json:"username" xml:"username" binding:"required"`
	Password string `gorm:"column:password;type:varchar(50);not null;" form:"password" json:"password" xml:"password" binding:"required"` //存入数据库前加密
	UserType int    `gorm:"column:usertype;type:tinyint(10);default:1"`                                                                   // 用户类型  默认是1，默认是普通用户  0是管理员
	Status   int    `gorm:"column:status;type:tinyint(10);default:1"`                                                                     // 账号状态  默认是1,激活状态
}

//创建表  自动迁移（把结构体和数据库表进行对应，就是帮你自动创建，修改数据库表（根据你对结构体的改变））

//用户的增删改查都放在这里

// 根据用户账号和密码查询用户是否存在
// func SearchUserByUserNameAndPassword(login *Login, user *User) (is_exist int) {
// 	is_exist = dao.DB.Where("username = ? and password = ? and status = ?", login.UserName, login.Password, 1).Find(user).RowsAffected
// 	return
// }
