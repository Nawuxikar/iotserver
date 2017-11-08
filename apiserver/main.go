package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ncoskay/goserver/apiserver/routers"
)
func main() {
	router := gin.Default()
	routers.Url(router)
	router.Run(":8081")

}
