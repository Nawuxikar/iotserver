package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)


func Login_html(c *gin.Context){
	//c.String(http.StatusOK,"hello world")
	c.HTML(http.StatusOK,"login.html",gin.H{
		"username":"kay",
	})
}
func Login (c *gin.Context){

}

func Logout (c *gin.Context){

}

func Register_html (c *gin.Context){

}

func Register (c *gin.Context){

}

