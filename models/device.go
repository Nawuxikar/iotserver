package models

import (
	"time"
	"errors"
	"github.com/satori/go.uuid"
	"fmt"
)



type Device struct {
	Id               int64
	Sn               string  			`xorm:"varchar(64) unique" form:"sn" json:"sn"`
	Name             string				`xorm:"varchar(64) " form:"name" json:"name"`
	Remarks			 string 			`xorm:"varchar(256) " form:"remarks" json:"remarks"`
	Userid			 string				`xorm:"varchar(64) " form:"uuid" json:"uuid"`
	Create           time.Time 			`xorm:"created"`
	Update           time.Time 			`xorm:"updated"`
	Lastlogin        time.Time 			`xorm:"updated"`

	Status   	     bool				`xorm:"default false"`
	Datashream_table string
	Is_public        bool				`xorm:"default true"`
	Devdata_struct	 []DatastremType
	SysTypeId		 int64				`form:"type" json:"type"`
	CustomTypeId	 int64
	//  {数据名称:}
	//  {"var":"int",""}

}

//系统数据流模型
type SysTypeDataModel struct {
	Id 				int64
	Name 			string   					`xorm:"varchar(64) " form:"name" json:"name"`
	Remarks			string 						`xorm:"varchar(256) " form:"remarks" json:"remarks"`
	Dataformat		[]DatastremType
}

//自定义产品数据模型
type CustomTypeDataModel struct {
	Id 				int64
	Name 			string						`xorm:"varchar(64) " form:"name" json:"name"`
	Remarks			string 						`xorm:"varchar(256) " form:"remarks" json:"remarks"`
	Owner			string						`xorm:"varchar(64) " form:"uuid" json:"uuid"`
	Dataformat		[]DatastremType
	//{"var":"int","button":"bool"}
}

type DatastremType struct {
	Key 		string
	Types		string
	Name		string
}

//
//数据流存储表
//type datastream struct{
//	id  				int
//	deviceid			string
//	dataid				string
//	type				string
//	data				string
//	time				string
//}


func (self *Device)AddNewDevice() (string,bool,error){

	if self.Userid == ""{
		return "",false,errors.New("UserUUID is Not Exist")
	}else{
		user := &User{Uuid:self.Userid}
		re,err :=Orm_sys.Get(user)
		if !re{
			return "",false,err
		}
		self.Datashream_table=user.Datashream_table
	}
	self.Sn = uuid.NewV4().String()

	if self.SysTypeId !=0{
		systype :=&SysTypeDataModel{Id:int64(self.SysTypeId)}
		re,_ := Orm_sys.Get(systype)
		if !re {
			//for k,v := range systype.Dataformat{
				//self.Devdata_struct[k]=v
			//}
			return "",false,errors.New("SysTypeId is Wrong")
		}
		self.Devdata_struct=systype.Dataformat
		self.Remarks=systype.Remarks
		if self.Name ==""{
			self.Name=systype.Name
		}


	}else {
		if self.Name ==""{
			self.Name="新设备"
		}
	}
	fmt.Println(self)
	_,err := Orm_sys.Insert(self)
	if err != nil{
		return "",false,err
	}
	return self.Sn,true,nil
}

//编辑自定义类型存储列表
func (self *SysTypeDataModel)Insert() (bool,error){

	_,err:=Orm_sys.Insert(self)
	if err!=nil {
		return false,err
	}
	return true,nil
}


//添加设备
func AddNewDevice(id string,name string,type_id int)(string,bool,error){

}


//保存数据
//参数 sn 设备序列号  k 数据名称   v 数据值
//返回 true/false  error
func SaveData(sn string,k string,v string) (bool,error){
	var re bool
	var err error
	ds := &Datastream{}
	device := &Device{Sn:sn}
	re,err = Orm_sys.Get(device)
	if !re {
		return false,err
	}

	//tp,ok := device.Devdata_struct[k]
	var n,i int
	var mp DatastremType

	for n,mp =range device.Devdata_struct{
		fmt.Println(n)
		fmt.Println(mp)
		if mp.Key == k{
			ds.Type=mp.Types
			ds.Name=mp.Name
			ds.Key=k
			ds.Value=v
			break
		}
		i=i+1
	}

	if (n+1)==i{
		return false,errors.New("Cannot Find Keys In The Struct")
	}

	ds.Deviceid= device.Id
	_ ,err =Orm_data.Table(device.Datashream_table).Insert(ds)
	if err !=nil{
		return false,err
	}
	return true,nil
}


func checkType(t string)bool{
	tp := []string{"int","bool","float","string"}
	for _,v := range tp{
		if t == v{
			return true
		}
	}
	return false
}