// 连接mysql
package dao

import (
	"Rspace_backend/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	DB *gorm.DB
)

func InitMySQL() error {
	var err error

	dsn := "root:123456@tcp(127.0.0.1:3306)/rspace?charset=utf8mb4&parseTime=True&loc=Local"
	if DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{}); err != nil {
		return err
	}

	return err
}

// 关闭数据库
func CloseMySQL() {
	sqlDB, err := DB.DB() //通过db.DB()获取底层的*sql.DB对象，然后调用其Close方法来关闭数据库连接

	if err != nil {
		panic(err)
	}
	sqlDB.Close()

}

// 初始化模型
func InitModel() {
	err := DB.AutoMigrate(&models.User{}, &models.NormalUser{})
	if err != nil {
		panic(err)
	}
}
