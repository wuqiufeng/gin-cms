package controllers

import (
	"cmsgo/pkg/api/dto"
	"net/http"

	"github.com/gin-gonic/gin"
)

type ControllerError struct {
	Code    int         `json:"code"`
	Message string      `json:"msg"`
	Data    interface{} `json:"data"`
}

const (
	SUCCESS = 0
	ERROR   = 10000
)

var (
	Err404          = &ControllerError{404, "页面没有找到", ""}
	ErrExpired      = &ControllerError{10011, "登录已过期", ""}
	ErrPermission   = &ControllerError{10012, "没有权限", ""}
	ErrChkJwt       = &ControllerError{10013, "无效的令牌", ""}
	ErrUnauthorized = &ControllerError{10010, "未登录或非法访问", ""}

	ErrInputData     = &ControllerError{10001, "数据输入错误", ""}
	ErrDatabase      = &ControllerError{10002, "服务器错误", ""}
	ErrValidation    = &ControllerError{13011, "请求参数验证失败", ""}
	ErrAddFail       = &ControllerError{11000, "创建失败", ""}
	ErrEditFail      = &ControllerError{11001, "更新失败", ""}
	ErrDelFail       = &ControllerError{11002, "删除失败", ""}
	ErrInvalidParams = &ControllerError{11003, "验证失败", ""}
	ErrIdData        = &ControllerError{10016, "此ID无数据记录", ""}
	OtherTaskRunning = &ControllerError{12000, "有其他任务在执行", ""}
	ErrUploadCover   = &ControllerError{12001, "上传封面失败", ""}
)

type BaseController struct {
}

func (bc *BaseController) BindAndValidate(c *gin.Context, obj interface{}) bool {
	if err := dto.Bind(c, obj); err != nil {
		failValidate(c, err.Error())
		return false
	}
	return true
}

func result(code int, data interface{}, msg string, c *gin.Context) {
	// 开始时间
	c.JSON(http.StatusOK, ControllerError{
		code,
		msg,
		data,
	})
}

func ok(c *gin.Context) {
	result(SUCCESS, map[string]interface{}{}, "操作成功", c)
}

func okWithMessage(message string, c *gin.Context) {
	result(SUCCESS, map[string]interface{}{}, message, c)
}

func okWithData(data interface{}, c *gin.Context) {
	result(SUCCESS, data, "操作成功", c)
}

func okDetailed(data interface{}, message string, c *gin.Context) {
	result(SUCCESS, data, message, c)
}

func fail(c *gin.Context) {
	result(ERROR, map[string]interface{}{}, "操作失败", c)
}

func failWithMessage(message string, c *gin.Context) {
	result(ERROR, map[string]interface{}{}, message, c)
}

func failWithDetailed(code int, data interface{}, message string, c *gin.Context) {
	result(code, data, message, c)
}

func failController(errs *ControllerError, c *gin.Context) {
	result(errs.Code, map[string]interface{}{}, errs.Message, c)
}

func failValidate(c *gin.Context, msg string) {
	errs := ErrValidation
	c.AbortWithStatusJSON(http.StatusOK, gin.H{
		"code": errs.Code,
		"msg":  errs.Message,
		"data": msg,
	})
}
