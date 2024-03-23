package controller

import (
	"Rspace_backend/dao"
	"Rspace_backend/middleware"
	"Rspace_backend/models"
	"fmt"
	"net/http"
	"sort"

	// "strconv"

	"github.com/gin-gonic/gin"
)

// 登录handler函数
func LoginHandler(c *gin.Context) {
	var logininfo middleware.LoginInfo
	// 得到前端传过来的账号和密码
	if err := c.ShouldBind(&logininfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
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

type Userinfo struct {
	NormalUserinfo models.NormalUser `form:"normal_userinfo" json:"normal_userinfo" xml:"normal_userinfo"`
	IsFollowed     bool              `form:"is_followed" json:"is_followed" xml:"is_followed"`
	FansCount      int64             `form:"fanscount" json:"fanscount" xml:"fanscount"`
	FollowerCount  int64             `form:"followercount" json:"followercount" xml:"followercount"`
}

func GetUserInfoByUserIDHandler(c *gin.Context) {
	// 得到中间件jwt认证的claims
	claims := c.MustGet("claims").(*middleware.Myclaims)
	//登录的用户id
	me_id := claims.ID
	// fmt.Print(me_id)
	//当前查看的用户id
	user_id := c.Query("user_id")

	// 需要判断user_id 是否被我me_id关注

	var user_info Userinfo

	// if claims != nil {
	// if strconv.Itoa(int(claims.ID)) == user_id {
	// 去数据库中查找用户的基本信息。
	// fmt.Println("--------------------", user_id)
	if err := dao.DB.Where("id = ?", user_id).Find(&user_info.NormalUserinfo).Error; err != nil {
		panic(err)
	}

	//查找是否关注，关注数，粉丝数

	//粉丝数
	if err := dao.DB.Model(&models.Follow{}).Where("isfollowed_user_id = ? AND status = ?", user_id, 1).Count(&user_info.FansCount).Error; err != nil {
		panic(err)
	}
	//关注数
	if err := dao.DB.Model(&models.Follow{}).Where("followed_user_id = ? AND status = ?", user_id, 1).Count(&user_info.FollowerCount).Error; err != nil {
		panic(err)
	}
	//查找是me_id是否是user_id的粉丝
	var count int64
	if err := dao.DB.Model(&models.Follow{}).Where("isfollowed_user_id = ? AND followed_user_id = ? AND status = ?", user_id, me_id, 1).Count(&count).Error; err != nil {
		panic(err)
	}
	if count == 1 {
		user_info.IsFollowed = true
	} else {
		user_info.IsFollowed = false
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0, //0表示成功
		"message": "get userinfo ok",
		"data":    user_info,
	})
	// return
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  -1,
	// 	"message": "token not match",
	// 	"data":    nil,
	// })
	// }
	// c.JSON(http.StatusOK, gin.H{
	// 	"status":  -1,
	// 	"message": "token is null",
	// 	"data":    nil,
	// })
}

// 绑定修改的用户信息
type UpdatedUserInfo struct {
	Name         string  `form:"name" json:"name" xml:"name"  binding:"required"` //姓名不能为空
	Gender       int     `form:"gender" json:"gender" xml:"gender"  binding:"required"`
	Age          *int    `form:"age" json:"age" xml:"age"  binding:"required"`
	Address      *string `form:"address" json:"address" xml:"address"  binding:"required"` // 设置为指针类型，即使有binding:"required"，也可以为空
	Introduction *string `form:"introduction" json:"introduction" xml:"introduction"  binding:"required"`
}

func UpdateUserInfoHandler(c *gin.Context) {
	// 得到中间件jwt认证的claims
	claims := c.MustGet("claims").(*middleware.Myclaims)
	id := claims.ID //得到用户id
	var updateduserinfo UpdatedUserInfo
	if err := c.ShouldBind(&updateduserinfo); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "error",
			"error":   err.Error(),
		})
		return
	}
	// fmt.Println(id)
	// fmt.Println("----------------------", *updateduserinfo.Address, *updateduserinfo.Introduction)
	//根据id插入数据{royal 2 2 成都 5555555555}
	if err_ := dao.DB.Model(&models.NormalUser{}).Where("id = ?", id).Updates(models.NormalUser{Name: updateduserinfo.Name, Age: updateduserinfo.Age, Gender: updateduserinfo.Gender, Address: updateduserinfo.Address, Introduction: updateduserinfo.Introduction}).Error; err_ != nil {
		panic(err_)
	}

	var normal_userinfo models.NormalUser
	if err := dao.DB.Where("id = ?", id).Find(&normal_userinfo).Error; err != nil {
		panic(err)
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0, //0表示成功
		"message": "update userinfo ok",
		"data":    normal_userinfo, //返回更新后的用户信息，那边马上更新全局的userinfo
	})
}

func GetUserPostsHandler(c *gin.Context) {
	// 得到中间件jwt认证的claims
	claims := c.MustGet("claims").(*middleware.Myclaims)
	// id := claims.ID //得到用户id
	if claims == nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "token is null",
			"data":    nil,
		})
		return
	}

	user_id := c.Query("user_id")

	// if strconv.Itoa(int(id)) != user_id {
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"status":  -1,
	// 		"message": "token not match",
	// 		"data":    nil,
	// 	})
	// 	return
	// }

	// 查询根据用户id用户的所有帖子
	var normaluser models.NormalUser
	if err := dao.DB.Model(&models.NormalUser{}).Where("id = ?", user_id).Preload("Posts").Find(&normaluser).Error; err != nil {
		panic(err)
	}

	// 自定义排序函数（降序排序desc）
	sort.Slice(normaluser.Posts, func(i, j int) bool {
		return normaluser.Posts[i].CreatedAt.After(normaluser.Posts[j].CreatedAt)
	})

	// fmt.Printf("----------------%#v", normaluser.Posts)
	//post以切片形式返回
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "get userposts success",
		"data":    normaluser.Posts,
	})
}

func UpdateAvatar(c *gin.Context) {
	// 先得到用户的id
	claims := c.MustGet("claims").(*middleware.Myclaims)
	id := claims.ID //得到用户id
	fmt.Printf("%#v--------------\n", id)
	file, _ := c.FormFile("avatar")
	// fmt.Printf("------------avatar=%#v\n", file.Filename)
	// 获得头像资源
	dst := fmt.Sprintf("./static/avatar/%s", file.Filename)
	fmt.Println(dst)
	err := c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"messgae": "update avatar error",
			"data":    nil,
		})
		return
	}
	// 存入数据库，并且返回file.Filename
	if err := dao.DB.Model(&models.NormalUser{}).Where("id = ?", id).Update("avatar", file.Filename).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"messgae": "update avatar error",
			"data":    nil,
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "update avatar success",
		"data":    file.Filename,
	})
}

// 更改用户之间的关注关系

type FollowUserID struct {
	UserID uint `form:"user_id" json:"user_id" xml:"user_id" binding:"required"`
}

func ChangeFollowStatus(c *gin.Context) {
	// 得到中间件jwt认证的claims
	claims := c.MustGet("claims").(*middleware.Myclaims)
	//登录的用户id
	me_id := claims.ID

	//被关注的用户id
	var user FollowUserID
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	user_id := user.UserID

	fmt.Printf("%#v-----%#v-----\n", user_id, me_id)

	//先查数据库，有没有此条记录，如果没有肯定是关注
	var follow models.Follow
	res := dao.DB.Where("isfollowed_user_id = ? AND followed_user_id = ?", user_id, me_id).Find(&follow)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "error",
			"data":    nil,
		})
		return
	}
	if res.RowsAffected != 0 { //如果存在，就把status改成相反的状态
		status := 0
		if *follow.Status == 0 {
			status = 1
		}
		if err := dao.DB.Model(&follow).Update("status", status).Error; err != nil {
			panic(err)
		}
	} else { //不存在就创建
		follow_info := models.Follow{IsFollowedUserID: user_id, FollowedUserID: me_id}
		if err := dao.DB.Create(&follow_info).Error; err != nil {
			panic(err)
		}
	}
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "change follow success",
		"data":    nil,
	})
}
