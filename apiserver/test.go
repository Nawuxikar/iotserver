package main

import (


	//"github.com/satori/go.uuid"
	///"github.com/ncoskay/goserver/models"

	"github.com/ncoskay/goserver/models"
	//"fmt"
	"fmt"
)

func main() {
	//user := models.User{Telephone:"18923212932",Password:"443622796"}
	//user := models.User{Telephone:"18923212932",Password:"4436227961"}
	//fmt.Println(user.AddNewUser())
	//fmt.Println(user.GetUser())
	//fmt.Println(user.UserLoginCheck())
	//fmt.Println(uuid.NewV1().String())

	//user:=models.User{Apikey:"4142498b938580716eb08458b8f8eb2b"}
	//fmt.Println(user.CheckUserApikeyExist())


	//device := &models.Device{Sn:"aee75740-c22d-41af-968c-a0e484e03a91"}
	//re,err:=device.ChangeDevType("te1mp","int")
	//fmt.Println(re)

	//fmt.Println(models.CheckType("string"))

	m1 :=[]models.DatastremType{{Key:"temp",Types:"int",Name:"温度"},{Key:"val",Types:"int",Name:"米秒"}}
	d := &models.SysTypeDataModel{Name:"我的LED灯",Remarks:"龙联络",Dataformat:m1}
	re,err:= d.Insert()


	//re,err := models.SaveData("e7bb2c04-efa4-4991-8c7b-abaeb9cf9ee5","temp","317")
	//fmt.Println(re)
	//fmt.Println(err)

	//d := &models.Datastream{Deviceid:10,Name:"新设备",Type:"int",Value:"123"}
	//re,err := models.Orm_data.Table("datastream_05bc27e4-b288-11e7-a5bb-b6ae2bc2d744").Insert(d)
	fmt.Println(re)
	fmt.Println(err)
}