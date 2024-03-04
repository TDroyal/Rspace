package routers

import (
	"Rspace_backend/controller"
	"Rspace_backend/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// var R *gin.Engine

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//设置跨域请求问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8080", "http://localhost:8080"} // 允许的源地址
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	r.Use(cors.New(config)) //注册全局中间件

	//登录的路由
	userGroup := r.Group("/api/token")
	{
		userGroup.POST("/", controller.LoginHandler)         //获得登录令牌
		userGroup.POST("/refresh/", middleware.RefreshJWT()) // 刷新登录令牌
	}

	return r
}
