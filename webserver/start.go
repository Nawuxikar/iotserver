package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ncoskay/goserver/webserver/routers"
)

func Start() {
	router := gin.Default()
	routers.Url(router)
	router.Run("0.0.0.0:8080")
}
