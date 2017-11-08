package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index_html(c *gin.Context){

	c.HTML(http.StatusOK,"index.html",gin.H{
		"username":"kay",
	})
}
