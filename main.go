package main

import (
	"LuXiStores/common"
	user_dao "LuXiStores/user/dao"
	user_handler "LuXiStores/user/handler"
	"fmt"
	"github.com/gin-gonic/gin"
)

//import "LuXiStores/user/handler"
func main() {
	r := gin.Default()
	common.Init()
	r.GET("/ping", user_handler.Hello)
	r.POST("/account/login", user_handler.Login)
	//r.POST("/account/signup", user_handler.UserSignup)
	//r.GET("/exit/:session", user_handler.Logout)
	//r.GET("/account/captcha", user_handler.GenerateCaptcha)
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	fmt.Println(user_dao.DB.GetUserInfoByUid(1))
	r.Run() // listen and serve on 0.0.0.0:8080
}
