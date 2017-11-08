package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/ncoskay/goserver/models"
	"net/http"
	"github.com/ncoskay/goserver/utils"

)

//查询用户名是否存在  GET
func SysUserExist (c *gin.Context){
	var user models.User
	err := c.Bind(&user)
	if err !=nil{
		utils.LogDebug(err.Error())
	}
	ac,em,te := user.CheckUserInfoExist()
	c.JSON(http.StatusOK,gin.H{
		"account":ac,
		"email":em,
		"telephone":te,
	})
}

//新建用户  POST
func SysAddNewUser	 (c *gin.Context){

	var user models.User
	var err error
	err = c.Bind(&user)

	if err !=nil {
		utils.LogDebug(err.Error())
	}

	re,errr:= user.AddNewUser()
	erstr :=""
	if errr==nil{
		erstr="Null"
	}else {
		erstr =errr.Error()
	}

	c.JSON(http.StatusOK,gin.H{
		"success":re,
		"error":erstr,
		"account":user.Account,
		"uuid":user.Uuid,
		//"email":user.Email,
		//"telephone":user.Telephone,

	})
}

//用户登录验证  POST
func SysUserLogin(c *gin.Context){
	var user models.User
	var err error
	err = c.Bind(&user)
	if err !=nil {
		utils.LogDebug(err.Error())
	}
	userback,ur,pr,err :=user.UserLoginCheck()
	if !ur || !pr{
		c.JSON(http.StatusOK,gin.H{
			"success":false,
			"error":err.Error(),
			"uuid":"",
		})
		return
	}
	c.JSON(http.StatusOK,gin.H{
		"success":true,
		"error":nil,
		"uuid":userback.Uuid,
	})
}