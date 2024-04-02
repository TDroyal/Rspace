package controller

import (
	"Rspace_backend/dao"
	"Rspace_backend/logic"
	"Rspace_backend/middleware"
	"Rspace_backend/models"
	"fmt"
	"net/http"
	"sort"
	"time"

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
	// fmt.Printf("-------%#v\n", logininfo)
	// 去数据库里面查看是否存在此用户，并且该用户状态正常
	var user models.User
	var cnt int64
	if logininfo.LoginMethod == "login_with_username" {
		cnt = dao.DB.Where("username = ? and password = ? and status = ?", logininfo.UserName, logic.GetMd5(logininfo.Password), 1).Find(&user).RowsAffected
	} else if logininfo.LoginMethod == "login_with_email" {
		cnt = dao.DB.Where("email = ?", logininfo.UserName).Find(&user).RowsAffected
	}
	// cnt := models.SearchUserByUserNameAndPassword(&login, &user)
	// fmt.Printf("-------%#v %d\n", user, cnt)

	// 存在此用户就返回登录成功的信息
	if cnt != 0 {
		//根据邮箱去查找redis中存的验证码
		if logininfo.LoginMethod == "login_with_email" {
			redis_email_code, err := logic.RedisGet(logininfo.UserName)
			if redis_email_code == "" || err != nil {
				c.JSON(http.StatusOK, gin.H{
					"status":  -2,
					"message": "验证码已经过期，请重新获取",
					"data":    err,
				})
				return
			}
			if redis_email_code != logininfo.Password {
				c.JSON(http.StatusOK, gin.H{
					"status":  -2,
					"message": "邮箱验证码不正确",
					"data":    err,
				})
				return
			}
			middleware.GenerateToken(c, middleware.LoginInfo{UserName: user.UserName, Password: user.Password}, user.ID)
		} else if logininfo.LoginMethod == "login_with_username" {
			//用户账号密码正确，就创建token
			middleware.GenerateToken(c, logininfo, user.ID)
		}
	} else {
		//统一为账户密码错误
		if logininfo.LoginMethod == "login_with_username" {
			c.JSON(http.StatusOK, gin.H{
				"status":  -1,
				"message": "username or password error",
				"data":    nil,
			})
		} else if logininfo.LoginMethod == "login_with_email" {
			c.JSON(http.StatusOK, gin.H{
				"status":  -2,
				"message": "此邮箱未注册",
				"data":    nil,
			})
		}
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
	// fmt.Println("-----------, id = ", id, updateduserinfo)
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

type SearchInfo struct {
	UserID     uint `form:"user_id" json:"user_id" xml:"user_id" binding:"required"`
	SearchType int  `form:"search_type" json:"search_type" xml:"search_type" binding:"required"`
}

type SearchResult struct {
	ID     uint   `form:"id" json:"id" xml:"id"`
	Name   string `form:"name" json:"name" xml:"name"`
	Avatar string `form:"avatar" json:"avatar" xml:"avatar"`
	Gender int    `form:"gender" json:"gender" xml:"gender"`
}

func GetFollowersInfo(c *gin.Context) {
	// 1表示查询关注，2表示查询粉丝
	var searchInfo SearchInfo
	if err := c.ShouldBind(&searchInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "error",
			"data":    err.Error(),
		})
		return
	}
	var userInfo []SearchResult
	if searchInfo.SearchType == 1 { //查询user_id关注的用户信息
		if err := dao.DB.Model(&models.Follow{}).
			Where("followed_user_id = ? AND status = ?", searchInfo.UserID, 1).
			Order("updated_at DESC").
			Select("normal_users.id, normal_users.name, normal_users.avatar, normal_users.gender").
			Joins("left join normal_users on normal_users.id = follows.isfollowed_user_id").
			Scan(&userInfo).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  -1,
				"message": "get follow error",
				"data":    err,
			})
			return
		}
		// fmt.Printf("%#v-------\n", userInfo)
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "get follow success",
			"data":    userInfo,
		})
		return
	} else if searchInfo.SearchType == 2 {
		if err := dao.DB.Model(&models.Follow{}).
			Where("isfollowed_user_id = ? AND status = ?", searchInfo.UserID, 1).
			Order("updated_at DESC").
			Select("normal_users.id, normal_users.name, normal_users.avatar, normal_users.gender").
			Joins("left join normal_users on normal_users.id = follows.followed_user_id").
			Scan(&userInfo).Error; err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"status":  -1,
				"message": "get fans error",
				"data":    err,
			})
			return
		}
		// fmt.Printf("%#v-------\n", userInfo)
		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "get fans success",
			"data":    userInfo,
		})
		return
	}

}

type StarPostInfo struct { //gorm匹配Scan从数据库查出来的字段
	PostID    uint      `form:"post_id" json:"post_id" xml:"post_id" gorm:"column:id"`
	PostType  int       `form:"post_type" json:"post_type" xml:"post_type" gorm:"column:type"`
	Content   string    `form:"content" json:"content" xml:"content" gorm:"column:content"`
	CreatedAt time.Time `form:"push_time" json:"push_time" xml:"push_time" gorm:"column:created_at"`
	UserID    uint      `form:"user_id" json:"user_id" xml:"user_id" gorm:"column:user_id"`
	Name      string    `form:"name" json:"name" xml:"name" gorm:"column:name"`
	Avatar    string    `form:"avatar" json:"avatar" xml:"avatar" gorm:"column:avatar"`
}

// 获取用户user_id收藏的帖子，包含用户头像，id，name, 帖子id, 帖子content，类型，帖子发布日期createAt,
func GetStarPostINfo(c *gin.Context) {
	user_id := c.Query("user_id")
	// fmt.Printf("------------------%#v\n", user_id)
	// 三表联查
	var starpost_info []StarPostInfo

	if err := dao.DB.Model(&models.Collection{}).
		Where("collections.user_id = ? AND collections.status = ?", user_id, 1).
		Order("collections.updated_at DESC").
		Select("posts.id, posts.type, posts.content, posts.created_at, posts.user_id, normal_users.name, normal_users.avatar").
		Joins("join posts on posts.id = collections.post_id AND posts.deleted_at is NULL").
		Joins("join normal_users on normal_users.id = posts.user_id").
		Scan(&starpost_info).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  -1,
			"message": "get starpostlist error",
			"data":    err,
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "get starpostlist success",
		"data":    starpost_info,
	})

	// fmt.Printf("%#v-------\n", starpost_info)
}

// 给用户邮箱发送验证码
type Emials struct {
	Email string `form:"email" json:"email" xml:"email"  binding:"required"`
	Type  string `form:"type" json:"type" xml:"type"  binding:"required"`
}

// [GIN] 2024/04/01 - 21:40:26 | 200 |    1.4075853s |       127.0.0.1 | POST     "/api/token/sendemailcode/"
func SendEmailCode(c *gin.Context) {
	// 判断是注册时获取邮箱，还是在登录时获取邮箱

	//获得用户的邮箱
	var e Emials
	if err := c.ShouldBind(&e); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "get email failed" + err.Error(),
			"data":    err,
		})
		return
	}
	// fmt.Println("邮箱", e.Email)
	// 先判断邮箱是否注册过，没有注册过再
	var count int64
	if err := dao.DB.Model(&models.User{}).Where("email = ?", e.Email).Count(&count).Error; err != nil {
		panic(err)
	}
	//注册时，邮箱已经存在
	if count != 0 && e.Type == "register" {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": "此邮箱已注册",
			"data":    nil,
		})
		return
	}

	// 用邮箱登录时，邮箱不存在
	if count == 0 && e.Type == "login" {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": "此邮箱还未注册",
			"data":    nil,
		})
		return
	}

	//发之前先判断此邮箱在60s内是否已经发送过验证码
	redis_email_code, err_ := logic.RedisGet(e.Email)
	if redis_email_code != "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": "验证码发送频繁，请稍后再再试",
			"data":    err_,
		})
		return
	}

	code := logic.GenerateRandCode()
	if err := logic.SendCodeByEmail(e.Email, code); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "send email code error" + err.Error(),
			"data":    err,
		})
		return
	}

	//发送成功后把code存入redis里面存一分钟
	if err := logic.RedisSet(e.Email, code, time.Second*60); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "redis set code error" + err.Error(),
			"data":    err,
		})
		return
	}
	// code, _ = logic.RedisGet(e.Email)
	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "send email code success",
		"data":    nil,
	})
}

// 绑定 JSON  注册的请求信息
type RegisterInfo struct {
	UserName  string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password  string `form:"password" json:"password" xml:"password" binding:"required"`
	Email     string `form:"email" json:"email" xml:"email" binding:"required"`
	EmailCode string `form:"email_code" json:"email_code" xml:"email_code" binding:"required"`
}

// 注册
func Register(c *gin.Context) {
	var register_info RegisterInfo
	if err := c.ShouldBind(&register_info); err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "get register info err" + err.Error(),
			"data":    err,
		})
		return
	}
	var count int64
	if err := dao.DB.Model(&models.User{}).Where("username = ?", register_info.UserName).Count(&count).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "search username error" + err.Error(),
			"data":    err,
		})
		return
	}
	if count != 0 {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": "此用户名已经存在",
			"data":    nil,
		})
		return
	}
	redis_email_code, err_ := logic.RedisGet(register_info.Email)
	if err_ != nil || redis_email_code == "" {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": "验证码已经过期，请重新获取",
			"data":    err_,
		})
		return
	}
	if redis_email_code != register_info.EmailCode {
		c.JSON(http.StatusOK, gin.H{
			"status":  -2,
			"message": "邮箱验证码不正确",
			"data":    err_,
		})
		return
	}

	//满足注册的条件
	//创建新的账号，密码（MD5加密），邮箱，用户名默认为账号，头像默认为灰色头像
	var userinfo = models.User{
		UserName: register_info.UserName,
		Password: logic.GetMd5(register_info.Password),
		Email:    register_info.Email,
	}
	if err := dao.DB.Create(&userinfo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "register error" + err.Error(),
			"data":    err,
		})
		return
	}
	var normal_userinfo = models.NormalUser{Name: register_info.UserName}
	if err := dao.DB.Create(&normal_userinfo).Error; err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": "register error" + err.Error(),
			"data":    err,
		})
		return
	}
	//表示注册成功，应该产生token并且直接进入系统
	// var logininfo = middleware.LoginInfo{UserName: register_info.UserName, Password: register_info.Password}
	var user models.User
	dao.DB.Where("username = ? AND email = ?", register_info.UserName, register_info.Email).Find(&user)
	middleware.GenerateToken(c, middleware.LoginInfo{UserName: user.UserName, Password: user.Password}, user.ID)
}
