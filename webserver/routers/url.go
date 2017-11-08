package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ncoskay/goserver/webserver/handlers"
)

func Url(router *gin.Engine) {
	router.GET("/index",handlers.Index_html)

	user := router.Group("/user/")
	{
		user.GET("login",handlers.Login_html)
		user.POST("login",handlers.Login)

		user.GET("logout",handlers.Logout)

		user.GET("register",handlers.Register_html)
		user.POST("register",handlers.Register)

	}
}
