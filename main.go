package main

import (
	"LuXiStores/user/dao"
	"LuXiStores/user/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hashicorp/go-uuid"
	"net/http"
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

	Uuid, err := uuid.GenerateUUID()
	if err != nil {
		fmt.Println("uuid error")
	}
	fmt.Println(Uuid)
	//
	r.GET("/login/:way", func(c *gin.Context) {
		way := c.Param("way")
		username := c.DefaultQuery("username", "")
		if way == "cookie" {
			sessionid, err := c.Cookie("sessionid")
			user, err := handler.CheckCookie(sessionid)
			if err != nil || user == "" {
				c.JSON(200, gin.H{
					"valid": "cookie验证失败",
				})

			} else if username == user {
				c.JSON(200, gin.H{
					"valid": "cookie验证成功",
				})
			} else {
				c.JSON(200, gin.H{
					"valid": "cookie验证失败",
				})
			}
		}
		username := c.DefaultQuery("username", "")
		password := c.DefaultQuery("password", "")
		if err != nil {
			if valid, err := handler.UserLogin(username, password); err == nil && valid == true {
				c.SetCookie("sessionid", Uuid, 0, "/", "", false, false)
				err := handler.SetCookie(Uuid, username)
				if err != nil {

				}
				c.JSON(200, gin.H{
					"valid": true,
				})
			} else {
				c.String(http.StatusOK, "用户名或密码错误")

			}
		} else {
			user, err := handler.CheckCookie(sessionid)
			if err != nil || user == "" {
				c.JSON(200, gin.H{
					"valid": "cookie验证失败",
				})

			} else if username == user {
				c.JSON(200, gin.H{
					"valid": "cookie验证成功",
				})
			} else {
				c.JSON(200, gin.H{
					"valid": "cookie验证失败",
				})
			}
		}
	})
	r.GET("/login/:cookie")
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	r.Run() // listen and serve on 0.0.0.0:8080
}
