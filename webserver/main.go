package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ncoskay/goserver/webserver/routers"
)
func main() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.Static("/static/","static/")
	routers.Url(router)
	router.Run(":8080")

}
