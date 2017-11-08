package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"github.com/ncoskay/goserver/models"

)

func MiddleWare_Auth() gin.HandlerFunc{
	return func(c *gin.Context){
		var re bool
		user := &models.User{}
		apikey:= c.Request.Header.Get("apikey")
		user.Apikey=apikey
		re,user =user.CheckUserApikeyExist()
		if re {
			//fmt.Println(user)
			//c.Set("apikey",user.Apikey)
			c.Set("uuid",user.Uuid)
			c.Next()
			return
		}
		c.JSON(http.StatusOK,gin.H{
			"error":"Apikey Is Wrong",
			"data":"",
		})
		c.Abort()
		return
	}
}