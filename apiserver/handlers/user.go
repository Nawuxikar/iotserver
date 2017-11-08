package handlers

import (
	"github.com/ncoskay/goserver/models"
	"net/http"
	"github.com/gin-gonic/gin"
	//"github.com/ncoskay/goserver/utils"
	//"github.com/go-blog/model"
	"strconv"
)

func UserAddNewDevice(c *gin.Context){
	//var user models.User
	//err := c.Bind(&user)
	//if err !=nil{
	//	utils.LogDebug(err.Error())
	//}aaaaaaa
	var err error
	var re bool
	var temp int
	device := &models.Device{}
	device.Name = c.PostForm("name")
	typeid :=c.PostForm("type")
	if typeid !=""{
		temp,err =strconv.Atoi(typeid)
		device.SysTypeId=int64(temp)
		if err!=nil{
			c.JSON(http.StatusOK,gin.H{
				"error":err.Error(),
				"success":false,
				"sn":"",
			})
			return
		}
	}

	uuid,_ := c.Get("uuid")
	device.Userid=uuid.(string)
	sn :=""
	sn,re,err=device.AddNewDevice()
	if !re{
		c.JSON(http.StatusOK,gin.H{
			"error":err.Error(),
			"success":re,
			"sn":"",
		})
	}

	c.JSON(http.StatusOK,gin.H{
		//"device":device,
		"sn":sn,
		"error":err,
		"success":re,
	})
}

func UserInsertOneData(c *gin.Context){
	sn := c.PostForm("sn")
	k := c.PostForm("key")
	v := c.PostForm("value")
	re,err := models.SaveData(sn,k,v)
	if !re{
		c.JSON(http.StatusOK,gin.H{
			"success":false,
			"error":err.Error(),
		})
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"error":nil,
	})
}



