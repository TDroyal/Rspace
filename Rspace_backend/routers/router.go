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
	loginGroup := r.Group("/api")
	{
		loginGroup.POST("/token/", controller.LoginHandler)          //获得登录令牌
		loginGroup.POST("/token/refresh/", middleware.RefreshJWT())  // 刷新登录令牌
		loginGroup.POST("/sendemailcode/", controller.SendEmailCode) //给邮箱发送验证码
		loginGroup.POST("/register/", controller.Register)
	}

	// 用户个人信息路由  更新个人信息
	userGroup := r.Group("/myspace")
	{
		userGroup.GET("/getuserinfo/", middleware.JWTAuth(), controller.GetUserInfoByUserIDHandler) //根据user_id返回用户的个人信息
		userGroup.POST("/updateuserinfo/", middleware.JWTAuth(), controller.UpdateUserInfoHandler)
		userGroup.GET("/getuserposts/", middleware.JWTAuth(), controller.GetUserPostsHandler)
		userGroup.POST("/updateavatar/", middleware.JWTAuth(), controller.UpdateAvatar)

		//修改关注状态
		userGroup.POST("/follow/", middleware.JWTAuth(), controller.ChangeFollowStatus)
		//查看用户的关注和粉丝
		userGroup.GET("/getfollowersinfo/", middleware.JWTAuth(), controller.GetFollowersInfo)
		//得到用户的收藏列表
		userGroup.GET("/getstarpostinfo/", middleware.JWTAuth(), controller.GetStarPostINfo)
	}

	// 首页帖子的获取
	postGroup := r.Group("/homepost")
	{
		postGroup.GET("/getposts/", controller.GetHomePost)
		// 根据帖子id获取帖子信息和用户基本信息
		postGroup.GET("/getpostbypostid/", controller.GetPostByPostId)

		// 根据帖子id获取评论(待完成)
		postGroup.GET("/getcommentsbypostid/", controller.GetCommentsByPostId)
		postGroup.DELETE("/deletecommentsbycommentid/", middleware.JWTAuth(), controller.DeleteCommentsByCommentId)
		postGroup.GET("/getiscollectcountbypostid/", controller.GetIscollectCountByPostId)
		// 根据帖子id删除帖子
		postGroup.DELETE("/deletepostbypostid/", middleware.JWTAuth(), controller.DeletePostByPostId)
	}

	// 发帖
	sharePostGroup := r.Group("/post")
	sharePostGroup.Use(middleware.JWTAuth())
	{
		sharePostGroup.POST("/uploadimage/", controller.UploadPostImageHandler)
		// sharePostGroup.POST("/uploadcontent/", controller.UploadPostContentHandler)

		// 根据帖子id和用户id发布评论(待完成)
		sharePostGroup.POST("/uploadcomment/", controller.UploadComment)

		sharePostGroup.POST("/changecollectStatus/", controller.ChangeCollectStatus)
	}
	return r
}
