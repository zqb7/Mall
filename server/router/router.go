package router

import (
	"github.com/zqhhh/gomall/handler/address"
	"github.com/zqhhh/gomall/handler/auth"
	"github.com/zqhhh/gomall/handler/cart"
	"github.com/zqhhh/gomall/handler/catalog"
	"github.com/zqhhh/gomall/handler/category"
	"github.com/zqhhh/gomall/handler/count"
	"github.com/zqhhh/gomall/handler/goods"
	"github.com/zqhhh/gomall/handler/index"
	"github.com/zqhhh/gomall/handler/order"
	"github.com/zqhhh/gomall/handler/region"
	"github.com/zqhhh/gomall/handler/topic"
	"github.com/zqhhh/gomall/router/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	if gin.IsDebugging() {
		gin.DisableConsoleColor()
	}
}

func InitEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.HandleError())
	engine.GET("/index", index.Index)
	engine.GET("/topic", topic.List)
	engine.GET("/catalog", catalog.List)
	engine.GET("/catalog/:id", catalog.Current)

	engine.GET("/category/:id", category.Detail)
	engine.GET("/category/:id/goods", category.Goods)

	engine.GET("/goods/:id", goods.Detail)
	engine.GET("/region", region.Get)
	engine.GET("/regionAll", region.ListAll)

	authGroup := engine.Group("/auth")
	{
		authGroup.POST("/reg", auth.Register)
		authGroup.POST("/login", auth.Login)
	}

	userGroup := engine.Group("/user")
	userGroup.Use(middleware.Auth())
	{
		userGroup.POST("/cart", cart.Add)
		userGroup.GET("/cart", cart.List)
		userGroup.DELETE("/cart/:id", cart.Del)
		userGroup.PUT("/cart/:id/checked", cart.Checked)
		userGroup.PUT("/cart/:id", cart.Update)
		userGroup.DELETE("/cart", cart.Clear)

		userGroup.POST("/address", address.Add)
		userGroup.GET("/address", address.List)
		userGroup.PUT("/address/:addressID", address.Update)

		userGroup.POST("/order", order.Create)
		userGroup.GET("/order/:orderSn", order.Detail)
		userGroup.GET("/order/:orderSn/pay", order.Pay)
		userGroup.GET("/order", order.List)
	}

	countGroup := engine.Group("/count")
	{
		countGroup.GET("/goods", count.Goods)
	}
	return engine
}
