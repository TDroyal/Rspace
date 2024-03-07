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

	// 设置静态文件服务
	r.Static("/static", "./static")

	//设置跨域请求问题
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://127.0.0.1:8080", "http://localhost:8080"} // 允许的源地址
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = append(config.AllowHeaders, "Authorization") // 添加允许的请求头
	r.Use(cors.New(config))                                            //注册全局中间件

	//登录的路由
	loginGroup := r.Group("/api/token")
	{
		loginGroup.POST("/", controller.LoginHandler)         //获得登录令牌
		loginGroup.POST("/refresh/", middleware.RefreshJWT()) // 刷新登录令牌
	}

	// 用户个人信息路由  更新个人信息
	userGroup := r.Group("/myspace")
	{
		userGroup.GET("/getuserinfo/", middleware.JWTAuth(), controller.GetUserInfoByUserIDHandler) //根据user_id返回用户的个人信息
		userGroup.POST("/updateuserinfo/", middleware.JWTAuth(), controller.UpdateUserInfo)
	}

	return r
}
