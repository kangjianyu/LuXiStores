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
	//黑名单
	r.POST("/blacklist",user_handler.BlackListAdd)
	r.DELETE("/blacklist",user_handler.BlackListDel)
	r.GET("/blacklist/check",user_handler.BlackListCheck)
	//用户资料
	r.GET("/user_profile",user_handler.GetProfile)
	r.PUT("/user_profile",user_handler.UpdateProfile)
	r.POST("/user_profile",user_handler.AddProfile)
	//会员
	r.GET("/user_profile/super",user_handler.GetUserSuperInfo)
	r.POST("user_profile/super",user_handler.AddUserSuperInfo)
	//用户登录
	r.POST("/account/login", user_handler.Login)
	//用户注册
	r.POST("/account/signup", user_handler.UserSignUp)
	//登出
	r.GET("/account/exit", user_handler.Logout)
	//验证码
	r.GET("/account/captcha", user_handler.GenerateCaptcha)
	//忘记密码
	r.POST("/account/forget",user_handler.ForgetPassword)
	r.POST("/account/update",user_handler.UpdatePassword)
	//商品分类
	r.GET("/goods/category", category_handler.CategoryForNext)
	r.PUT("/goods/category", category_handler.CategoryUpdate)
	r.POST("/goods/category", category_handler.AddCategory)
	//商品信息
	r.GET("/goods/list",goods_handler.GetGoodsInfo)
	r.GET("/goods/info",goods_handler.GetGoodsInfoDetail)
	r.POST("/goods/info",goods_handler.AddGoodsInfo)
	r.PUT("/goods/info",goods_handler.UpdateGoodsInfo)
	r.DELETE("/goods/info",goods_handler.DelGoodsInfo)
	r.PUT("/goods/status",goods_handler.UpdateGoodsStatus)
	//收货地址
	r.GET("/goods/receiver/list", receiver_handler.GetGoodsReceiverAddressList)
	r.DELETE("/goods/receiver/info", receiver_handler.DelGoodsReceiverAddress)
	r.POST("/goods/receiver/info", receiver_handler.AddGoodsReceiverAddress)
	r.PUT("/goods/receiver/info", receiver_handler.UpdateGoodsReceiverAddress)
	r.PUT("/goods/receiver/default", receiver_handler.SetGoodsReceiverAddress)
	//购物车
	r.GET("/goods/cart/list",cart_handler.GetGoodsCartList)

	r.DELETE("/goods/cart/",cart_handler.DelGoodsCartList)
	r.POST("/goods/cart/",cart_handler.AddGoodsCartList)
	r.PUT("/goods/cart/",cart_handler.UpdateGoodsCartList)
	//收藏
	r.GET("/goods/collection/list",goods_handler.GetGoodsCollectionList)
	r.POST("/goods/collection",goods_handler.AddGoodsCollection)
	r.DELETE("/goods/collection",goods_handler.DelGoodsCollection)
	//订单
	r.POST("/order",order_handler.AddOrder)
	r.GET("/order/check",order_handler.CheckAliPay)
	//评论
	r.POST("/order/comment",order_handler.AddOrderComment)
	r.DELETE("/order/comment",order_handler.DelOrderComment)
	//交易
	r.POST("/trade/add",order_handler.AddTrade)
	//短链接
	r.GET("/m78/:ww",shorturl_handler.GetLongUrl)
	r.POST("/shorturl/add",shorturl_handler.AddShortUrl)


	//r.GET("/user/:name/*action", func(c *gin.Context) {
	//	name := c.Param("name")
	//	action := c.Param("action")
	//	message := name + " is " + action
	//	c.String(http.StatusOK, message)
	//})
	//kafka.SaramaProducer()

	r.Run("127.0.0.1:8001") // listen and serve on 0.0.0.0:8080

}
