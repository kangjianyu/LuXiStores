package main

import (
	"LuXiStores/backed"
	cart_handler "LuXiStores/cart/handler"
	"LuXiStores/category/handler"
	"LuXiStores/common"
	goods_handler "LuXiStores/goods/handler"
	order_handler "LuXiStores/order/handler"
	"LuXiStores/receiver/handler"
	"LuXiStores/shorturl/handler"
	user_handler "LuXiStores/user/handler"
	"github.com/gin-gonic/gin"
)

//import "LuXiStores/user/handler"
func main() {
	r := gin.Default()
	common.Init()
	backed.InitBackend()
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
	r.GET("/goods/category", category_handler.CategoryForNext)
	r.POST("/goods/category/update", category_handler.CategoryUpdate)
	r.POST("/goods/category/add", category_handler.CategoryAdd)
	r.GET("/goods/list",goods_handler.GetGoodsInfo)
	r.GET("/goods/list/detail",goods_handler.GetGoodsInfoDetail)
	r.POST("/goods/add",goods_handler.AddGoodsInfo)
	r.POST("/goods/update",goods_handler.UpdateGoodsInfo)
	r.POST("/goods/status",goods_handler.UpdateGoodsStatus)
	r.POST("/goods/del",goods_handler.DelGoodsInfo)
	r.GET("/goods/receiver/list", receiver_handler.GetGoodsReceiverAddressList)
	r.POST("/goods/receiver/del", receiver_handler.DelGoodsReceiverAddress)
	r.POST("/goods/receiver/add", receiver_handler.AddGoodsReceiverAddress)
	r.POST("/goods/receiver/update", receiver_handler.UpdateGoodsReceiverAddress)
	r.POST("/goods/receiver/default", receiver_handler.SetGoodsReceiverAddress)
	r.POST("/goods/cart/list",cart_handler.GetGoodsCartList)
	r.POST("/goods/cart/del",cart_handler.DelGoodsCartList)
	r.POST("/goods/cart/add",cart_handler.AddGoodsCartList)
	r.POST("/goods/cart/update",cart_handler.UpdateGoodsCartList)
	r.POST("/order/add",order_handler.AddOrder)
	r.GET("/order/check",order_handler.CheckAliPay)
	r.POST("/order/comment/add",order_handler.AddOrderComment)
	r.POST("/order/comment/del",order_handler.DelOrderComment)
	r.POST("/trade/add",order_handler.AddTrade)
	r.GET("/m78/:ww",shorturl_handler.GetLongUrl)
	r.POST("/shorturl/add",shorturl_handler.AddShortUrl)
	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	r.Run() // listen and serve on 0.0.0.0:8080
}
