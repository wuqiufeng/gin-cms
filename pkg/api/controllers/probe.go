package controllers

import "github.com/gin-gonic/gin"

func Healthy(c *gin.Context) {
	fail(c, Err404)
}
