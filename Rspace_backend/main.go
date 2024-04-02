package main

import (
	"Rspace_backend/dao"
	"Rspace_backend/routers"
	// "github.com/gin-gonic/gin"
)

func main() {
	//连接mysql数据库
	err := dao.InitMySQL()
	if err != nil {
		panic(err)
	}
	//连接redis
	dao.InitRedis()

	//持续退出时mysql数据库
	defer dao.CloseMySQL()

	// 模型绑定
	dao.InitModel()

	// 注册路由
	r := routers.SetupRouter()

	r.Run(":9090")
}
