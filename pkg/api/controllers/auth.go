package controllers

import (
	"cmsgo/pkg/api/middleware"
	"cmsgo/pkg/api/model/request"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	BaseController
}

func (u *AuthController) Captcha(c *gin.Context) {
	ok(c)
}

// @Summary 用户登陆
func (u *AuthController) JwtAuthLogin(c *gin.Context) {
	token := JwtToken(c)
	okWithData(token, c)
}

// @Summary 用户refresh-token接口
func (u *AuthController) JwtAuthRefreshLogin(c *gin.Context) {
	//jwtAuth.RefreshHandler(c)
}

//签发jwt
func JwtToken(c *gin.Context) string {
	j := &middleware.JWT{
		[]byte("qmPlus"), // 唯一签名
	}

	clams := request.CustomClaims{
		ID:          2,
		NickName:    "haha",
		AuthorityId: "authorityid",
		StandardClaims: jwt.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000),       // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 60*60*24*7), // 过期时间 一周
			Issuer:    "qmPlus",                              //签名的发行者
		},
	}

	token, err := j.CreateToken(clams)
	if err != nil {
		fmt.Println("获取token失败")
	} else {
		fmt.Println(token)
	}
	return token

}
