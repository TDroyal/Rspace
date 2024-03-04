package controller

import (
	"Rspace_backend/dao"
	"Rspace_backend/middleware"
	"Rspace_backend/models"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// 登录handler函数
func LoginHandler(c *gin.Context) {
	var logininfo middleware.LoginInfo
	// 得到前端传过来的账号和密码
	if err := c.ShouldBind(&logininfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	fmt.Printf("-------%#v\n", logininfo)
	// 去数据库里面查看是否存在此用户，并且该用户状态正常
	var user models.User
	cnt := dao.DB.Where("username = ? and password = ? and status = ?", logininfo.UserName, logininfo.Password, 1).Find(&user).RowsAffected
	// cnt := models.SearchUserByUserNameAndPassword(&login, &user)
	// fmt.Printf("-------%#v %d\n", user, cnt)

	// 存在此用户就返回登录成功的信息
	if cnt != 0 {
		//用户账号密码正确，就创建token
		middleware.GenerateToken(c, logininfo, user.ID)

	} else { //统一为账户密码错误
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "username or password error",
			"data":    nil,
		})
	}
}

func GetUserInfoByUserID(c *gin.Context) {
	// 得到中间件jwt认证的claims
	claims := c.MustGet("claims").(*middleware.Myclaims)
	user_id := c.Query("user_id")
	var normal_userinfo models.NormalUser

	if claims != nil {
		if strconv.Itoa(int(claims.ID)) == user_id {
			// 去数据库中查找用户的基本信息。
			// fmt.Println("--------------------", user_id)
			if err := dao.DB.Where("id = ?", user_id).Find(&normal_userinfo).Error; err != nil {
				panic(err)
			}
			c.JSON(http.StatusOK, gin.H{
				"status":  0,
				"message": "get userinfo ok",
				"data":    normal_userinfo,
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "token not match",
			"data":    nil,
		})
	}

}
