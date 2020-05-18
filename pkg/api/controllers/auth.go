package controllers

import "github.com/gin-gonic/gin"

type AuthController struct {
	BaseController
}

func (u *AuthController) Captcha(c *gin.Context) {
	ok(c, "hahaha")
}
