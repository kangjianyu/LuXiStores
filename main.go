package main

import (
	"LuXiStores/common"
	"LuXiStores/goods/handler"
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
	r.POST("/blacklist/add",user_handler.BlackListAdd)
	r.POST("/blacklist/del",user_handler.BlackListDel)
	r.GET("/blacklist/check",user_handler.BlackListCheck)
	r.GET("/user_profile/list",user_handler.GetProfile)
	r.POST("/user_profile/update",user_handler.UpdateProfile)
	r.POST("/account/login", user_handler.Login)
	r.POST("/account/signup", user_handler.UserSignUp)
	r.GET("/account/exit", user_handler.Logout)
	r.GET("/account/captcha", user_handler.GenerateCaptcha)
	r.POST("/account/forget",user_handler.ForgetPassword)
	r.POST("/account/update",user_handler.UpdatePassword)
	r.GET("/goods/category",goods_handler.CategoryForNext)
	r.POST("/goods/category/update",goods_handler.CategoryUpdate)
	r.POST("/goods/category/add",goods_handler.CategoryAdd)
	r.GET("/goods/list",goods_handler.GetGoodsInfo)
	r.GET("/goods/list/detail",goods_handler.GetGoodsInfoDetail)
	r.POST("/goods/add",goods_handler.AddGoodsInfo)
	r.POST("/goods/update",goods_handler.UpdateGoodsInfo)
	r.POST("/goods/status",goods_handler.UpdateGoodsStatus)
	r.POST("/goods/del",goods_handler.DelGoodsInfo)
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	fmt.Println(user_dao.Rds.GetUserToken("foo"))
	fmt.Println(user_dao.DB.GetUserInfoByUid(1))
	r.Run() // listen and serve on 0.0.0.0:8080
}
