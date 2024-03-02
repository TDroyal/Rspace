package routers

import (
	"Rspace_backend/controller"

	"github.com/gin-gonic/gin"
)

// var R *gin.Engine

func SetupRouter() *gin.Engine {
	r := gin.Default()

	//登录的路由
	r.POST("/login", controller.LoginHandler)

	return r
}
