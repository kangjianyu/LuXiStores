package common

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var (
	ErrIsNil            = errors.New("操作成功")
	ErrParam            = errors.New("请求参数错误")
	ErrInternal         = errors.New("内部系统错误")
	ErrSessionExpire    = errors.New("会话已过期")
	ErrTransIdDuplicate = errors.New("trans_id重复")
	ErrSetLockFailed    = errors.New("获取lock失败")
	ErrRedisKeyNotExist = errors.New("redis key not exist")
	ErrCaptcha          = errors.New("验证码错误")
	ErrAuth             = errors.New("用户不存在或密码不正确")
)

var errMap = map[error]int{
	nil:              0,
	ErrIsNil:         0,
	ErrParam:         499,
	ErrInternal:      500,
	ErrSessionExpire: 604,
	ErrAuth:          1001,
	ErrCaptcha:       1005,
}

type BaseRes struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
}

func BuildResp(c *gin.Context, data interface{}, err error) map[string]interface{} { // 连错误码也一起返回, 简化代码
	if err == nil {
		err = ErrIsNil
	}
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
	_, exist := errMap[err]
	if !exist {
		err = ErrInternal
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
