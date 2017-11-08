package models

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"github.com/ncoskay/goserver/utils"
)

var Orm_sys 	*xorm.Engine
var Orm_data 	*xorm.Engine

func init() {
	SetEngine()
}

func SetEngine() {
	server := "120.55.64.211"
	username := "kay"
	password := "443622796"
	sys_dbName := "orm_sys"
	data_dbName := "orm_data"
	var err error
	Orm_sys, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":3306)/"+sys_dbName+"?charset=utf8")
	if err != nil {
		utils.LogErr(err.Error())
		return
	}
	Orm_data, err = xorm.NewEngine("mysql", username+":"+password+"@tcp("+server+":3306)/"+data_dbName+"?charset=utf8")

	if err != nil {
		utils.LogErr(err.Error())
		return
	}

	Orm_sys.ShowSQL(true)
	err = Orm_sys.Sync(new(User), new(Device),new(SysTypeDataModel),new(CustomTypeDataModel))
	if err != nil {
		utils.LogErr(err.Error())
		return
	}
	//return orm
}


