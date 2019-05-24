package common

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrIsNil            = errors.New("操作成功")
	ErrParamError       = errors.New("请求参数错误")
	ErrInternalError    = errors.New("内部系统错误")
	ErrSessionExpire    = errors.New("会话已过期")
	ErrTransIdDuplicate = errors.New("trans_id重复")
	ErrSetLockFailed    = errors.New("获取lock失败")
	ErrRedisKeyNotExist = errors.New("redis key not exist")
)

var errMap = map[error]int{
	nil:                 0,
	ErrIsNil:            0,
	ErrParamError:       499,
	ErrInternalError:    500,
	ErrSessionExpire:    604,
	ErrTransIdDuplicate: 20005,
}

type BaseRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func BuildHttpResp(c *gin.Context, data interface{}, err error) map[string]interface{} { // 连错误码也一起返回, 简化代码
	if err == nil {
		err = ErrIsNil
	}
	_, exist := errMap[err]
	if !exist {
		err = ErrInternalError
	}
	dmError := errMap[err]
	errorMsg := err.Error()
	resp := gin.H{
		"code": dmError,
		"msg":  errorMsg,
	}
	if data != nil {
		resp["data"] = data
	}
	c.JSON(200, resp)
	return resp
}
