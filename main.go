package main

import (
	"LuXiStores/user/dao"
	"LuXiStores/user/handler"
	"github.com/gin-gonic/gin"
)

//import "LuXiStores/user/handler"
func main() {
	r := gin.Default()
	dao.Init()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/account/login", handler.UserLogin)
	r.POST("/account/signup", handler.UserSignup)
	r.GET("/exit/:session", handler.UserExit)
	r.GET("/account/captcha", handler.GenerateCaptcha)
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	r.Run() // listen and serve on 0.0.0.0:8080
}
