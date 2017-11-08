package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/ncoskay/goserver/apiserver/handlers"
	"github.com/ncoskay/goserver/apiserver/middlewares"
)

func Url(router *gin.Engine) {

	system := router.Group("/sys/")
	{
		system.GET("exist",handlers.SysUserExist)

		system.POST("newuser",handlers.SysAddNewUser)
		system.POST("login",handlers.SysUserLogin)
	}

	user := router.Group("/user/",middlewares.MiddleWare_Auth())
	{
		user.POST("newdevice",handlers.UserAddNewDevice)
		user.POST("savedata",handlers.UserInsertOneData)
	}
}

