package controller

import (
	"Rspace_backend/dao"
	"Rspace_backend/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 登录handler函数
func LoginHandler(c *gin.Context) {
	var login models.Login
	// 得到前端传过来的账号和密码
	if err := c.ShouldBind(&login); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	fmt.Printf("-------%#v\n", login)
	// 去数据库里面查看是否存在此用户，并且该用户状态正常
	var user models.User
	cnt := dao.DB.Where("username = ? and password = ? and status = ?", login.UserName, login.Password, 1).Find(&user).RowsAffected
	// cnt := models.SearchUserByUserNameAndPassword(&login, &user)
	// fmt.Printf("-------%#v %d\n", user, cnt)

	// 存在此用户就返回登录成功的信息
	if cnt != 0 {
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"is_login": 1,
			"userid":   user.ID,
		})
	} else { //统一为账户密码错误
		c.JSON(http.StatusOK, gin.H{
			"message":  "failed",
			"is_login": 0,
		})
	}
}
