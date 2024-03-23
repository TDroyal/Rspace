// 连接mysql
package dao

import (
	"Rspace_backend/models"

	// "fmt"

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
	err := DB.AutoMigrate(&models.User{}, &models.NormalUser{}, &models.Post{}, &models.Comment{}, &models.Collection{}, &models.Follow{})
	// 这里做插入操作,测试

	// post := []models.Post{
	// 	{Content: "今天很开心,今天很开心,今天很开心,今天很开心今天很开心今天很开心今天很开心今天很开心今天很开心今天很开心今天很开心今天很开心", UserID: 1},
	// 	{Content: "我的Share很有趣！！！", UserID: 1},
	// 	{Content: "Go是最好的语言", UserID: 1},
	// }
	// DB.Create(&post)

	// 做查询操作试一下
	// var normaluser []models.NormalUser
	// DB.Model(&models.NormalUser{}).Preload("Posts").Find(&normaluser)  //把

	// 相当于把每个用户的Posts都查找到，并且放在Posts属性里面了
	// fmt.Printf("--------------%#v-----------\n", *normaluser[0].Posts[1].Image)

	if err != nil {
		panic(err)
	}
}
