package middleware

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

// 绑定 JSON  登录请求信息
type LoginInfo struct {
	UserName string `form:"username" json:"username" xml:"username"  binding:"required"`
	Password string `form:"password" json:"password" xml:"password" binding:"required"`
}

// 定义荷载payload
type Myclaims struct {
	ID        uint   `json:"user_id"`
	UserName  string `json:"username"`
	TokenType string `json:"token_type"` //jwt, refresh_jwt
	// StandardClaims结构体实现了Claims接口(Valid()函数)
	jwtgo.StandardClaims
}

// 定义密钥
type JWT struct {
	SigningKey []byte
}

// refresh_jwt
type Refresh_JWT struct {
	Refresh_Token string `form:"refresh_jwt" json:"refresh_jwt" xml:"refresh_jwt"  binding:"required"`
}

// 登陆结果
type LoginResult struct {
	Token         string `json:"token"`
	Refresh_Token string `json:"refresh_token"`
	// UserName      string `json:"username"`
	// UserID        uint   `json:"user_id"`
}

var (
	secret              = "iamsecret" //私钥
	ErrTokenExpired     = errors.New("token is expired")
	ErrTokenNotValidYet = errors.New("token not active yet")
	ErrTokenMalformed   = errors.New("that's not even a token")
	ErrTokenInvalid     = errors.New("couldn't handle this token")
)

func NewJWT() *JWT {
	return &JWT{
		SigningKey: []byte(secret),
	}
}

// 创建Token(基于用户的基本信息claims)
// 使用HS256算法进行token生成
// 使用用户基本信息claims以及签名key(signkey)生成token
func (j *JWT) CreateToken(claims Myclaims) (string, error) {
	// 返回一个token的结构体指针
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// 生成token
func GenerateToken(c *gin.Context, logininfo LoginInfo, user_id uint) { //用户登录成功后，会生成jwt和refresh_jwt两个令牌
	// 构造SignKey: 签名和解签名需要使用一个值
	j := NewJWT()

	// 构造用户claims信息(负荷) jwt信息
	claims := Myclaims{
		user_id,
		logininfo.UserName,
		"jwt",
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 600),  // 签名过期时间  10分钟  60 * 10
			Issuer:    "royal_111",                     // 签名颁发者
		},
	}

	// 构造用户refresh_claims信息，refresh_jwt   有效期时间一般为一周
	refresh_claims := Myclaims{
		user_id,
		logininfo.UserName,
		"refresh_jwt",
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),   // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 604800), // 签名过期时间  一周
			Issuer:    "royal_111",                       // 签名颁发者
		},
	}

	// 根据claims生成token对象
	token, err := j.CreateToken(claims)
	refresh_token, err_ := j.CreateToken(refresh_claims)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": err.Error(),
			"data":    nil,
		})
	}

	if err_ != nil {
		c.JSON(http.StatusOK, gin.H{
			"status":  -1,
			"message": err_.Error(),
			"data":    nil,
		})
	}

	log.Println(token)
	// 返回用户相关数据，
	data := LoginResult{
		// UserName:      logininfo.UserName,
		Token:         token,         // 其中token也返回给用户，用户每次请求需要认证的api时带上token进行身份认证。
		Refresh_Token: refresh_token, //refresh_jwt
		// UserID:   user_id,
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  0,
		"message": "登陆成功",
		"data":    data,
	})

	// return
}

// 解析token
func (j *JWT) ParserToken(tokenstr string) (*Myclaims, error) {
	// 输入token
	// 输出自定义函数来解析token字符串为jwt的Token结构体指针
	// Keyfunc是匿名函数类型: type Keyfunc func(*Token) (interface{}, error)
	// func ParseWithClaims(tokenString string, claims Claims, keyFunc Keyfunc) (*Token, error) {}
	token, err := jwtgo.ParseWithClaims(tokenstr, &Myclaims{}, func(token *jwtgo.Token) (interface{}, error) {
		return j.SigningKey, nil
	})

	fmt.Println(token, err)
	if err != nil {
		// jwt.ValidationError 是一个无效token的错误结构
		if ve, ok := err.(*jwtgo.ValidationError); ok {
			// ValidationErrorMalformed是一个uint常量，表示token不可用
			if ve.Errors&jwtgo.ValidationErrorMalformed != 0 {
				return nil, ErrTokenMalformed
				// ValidationErrorExpired表示Token过期
			} else if ve.Errors&jwtgo.ValidationErrorExpired != 0 {
				return nil, ErrTokenExpired
				// ValidationErrorNotValidYet表示无效token
			} else if ve.Errors&jwtgo.ValidationErrorNotValidYet != 0 {
				return nil, ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}

		}
	}

	// 将token中的claims信息解析出来和用户原始数据进行校验
	// 做以下类型断言，将token.Claims转换成具体用户自定义的Claims结构体
	if claims, ok := token.Claims.(*Myclaims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("token NotValid")
}

// 刷新令牌中间件
func RefreshJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var refresh_jwt Refresh_JWT
		if err := c.ShouldBind(&refresh_jwt); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  -1,
				"message": "未得到refresh_token",
				"error":   err.Error(),
			})
			c.Abort()
			return
		}

		// fmt.Println("refresh_jwt = ", refresh_jwt)

		j := NewJWT() //得到私钥
		refresh_claims, err := j.ParserToken(refresh_jwt.Refresh_Token)

		if err != nil {
			// token过期
			if err == ErrTokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "token授权已过期, 请重新申请授权",
					"data":   nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   nil,
			})
			c.Abort()
			return
		}

		// fmt.Println("--------", refresh_claims.ID)
		// 根据解析出来的用户信息构造当前用户的专属令牌
		// 构造用户claims信息(负荷) jwt信息
		claims := Myclaims{
			refresh_claims.ID,
			refresh_claims.UserName,
			"jwt",
			jwtgo.StandardClaims{
				NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
				ExpiresAt: int64(time.Now().Unix() + 600),  // 签名过期时间  10分钟  60 * 10
				Issuer:    "royal_111",                     // 签名颁发者
			},
		}

		// 根据claims生成token对象
		token, err := j.CreateToken(claims)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"status":  -1,
				"message": err.Error(),
				"data":    nil,
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  0,
			"message": "refresh jwt success",
			"jwt":     token,
		})
	}
}

// 身份认证中间件
func JWTAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		// fmt.Println("88888888888888888888")
		//拿到token
		authHeader := c.Request.Header.Get("Authorization")
		// fmt.Println("************", authHeader)
		// if authHeader == "Bearer " {
		// 	c.JSON(http.StatusOK, gin.H{
		// 		"status": -1,
		// 		"msg":    "token为空, 请携带token",
		// 		"data":   nil,
		// 	})
		// 	c.Abort()
		// 	return
		// }

		token := strings.TrimPrefix(authHeader, "Bearer ") //去除前缀
		if token == "" {
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    "token为空, 请携带token",
				"data":   nil,
			})
			c.Abort()
			return
		}

		fmt.Println("token = ", token)

		//解析出实际的载荷
		j := NewJWT() //得到私钥

		claims, err := j.ParserToken(token)
		if err != nil {
			// token过期
			if err == ErrTokenExpired {
				c.JSON(http.StatusOK, gin.H{
					"status": -1,
					"msg":    "token授权已过期, 请重新申请授权",
					"data":   nil,
				})
				c.Abort()
				return
			}
			// 其他错误
			c.JSON(http.StatusOK, gin.H{
				"status": -1,
				"msg":    err.Error(),
				"data":   nil,
			})
			c.Abort()
			return
		}

		// 解析到具体的claims相关信息  Set is used to store a new key/value pair exclusively for this context.
		c.Set("claims", claims) // 身份认证成功后，下游可以handler函数可以使用该参数
	}
}
