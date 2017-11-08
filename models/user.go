package models

import (
	"time"
	"github.com/ncoskay/goserver/utils"
	"regexp"
	"github.com/satori/go.uuid"
	"fmt"
	"errors"
)

const datasheet_sql  = "(`id`  bigint NOT NULL AUTO_INCREMENT ," +
						"`deviceid`  int NULL ," +
						"`name`  varchar(64) NULL,"+
						"`type`  varchar(255) NULL ," +
						"`key`  varchar(255) NULL," +
						"`value`  varchar(255) NULL," +
						"`time`  datetime NULL ,PRIMARY KEY (`id`)) DEFAULT CHARSET utf8"

type Datastream struct {
	Id				int64
	Deviceid		int64
	Name 			string
	Type 			string
	Key             string
	Value			string
	Time			time.Time     `xorm:"created"`
}


type User struct {
	Id         int64       //`xorm:"int(11) pk not null autoincr"`
	Uuid       string    `xorm:"varchar(64) unique"`
	Apikey		string   `xorm:"varchar(32) unique"`

	Account    string    `xorm:"varchar(32) unique" form:"account" json:"account"`
	Email       string	 `xorm:"varchar(64) " form:"email"   json:"email"`
	Telephone   string	 `xorm:"varchar(14) " form:"telephone" json:"telephone"`

	Password   	string    `xorm:"varchar(128)" form:"password" json:"password"`
	AuthGroup 	string    `xorm:"default 'ordinary'"`   //权限组   ordinary  admin

	Register   time.Time `xorm:"created"`
	Lastlogin  time.Time `xorm:"updated"`

	Name        string  `xorm:"varchar(128)" form:"name" json:"name"`
	Gender      string 	`xorm:"varchar(128)" form:"gender" json:"gender"`
	Age         int

	Status 		bool
	Iconaddress string   //图标地址
	//Devicelist []int64
	Datashream_table string   //用户数据流表名
}

// 检查用户名是否存在
// 参数
//返回 bool  True 存在    Flase 不存在或为空
func (self *User)CheckUserAccountExist() bool{
	var re bool
	var err error
	re=false
	if self.Account!=""{
		re,err =Orm_sys.Get(&User{Account:self.Account})
		if err!=nil{
			utils.LogErr(err.Error())
		}
	}
	return re
}

// 检查邮箱是否存在
// 参数
//返回 bool  True 存在    Flase 不存在或为空
func (self *User)CheckUserEmailExist() bool{
	var re bool
	var err error
	re=false
	if self.Email!=""{
		re,err =Orm_sys.Get(&User{Email:self.Email})
		if err!=nil{
			utils.LogErr(err.Error())
		}
	}
	return re
}

// 检查手机号码是否存在
// 参数
//返回 bool  True 存在    Flase 不存在或为空
func (self *User)CheckUserTelephoneExist() bool{
	var re bool
	var err error
	re=false
	if self.Telephone!=""{
		re,err =Orm_sys.Get(&User{Telephone:self.Telephone})
		if err!=nil{
			utils.LogErr(err.Error())
		}
	}
	return re
}

// 检查APIKEY是否存在
// 参数
//返回 bool  True 存在    Flase 不存在或为空
func (self *User)CheckUserApikeyExist() (bool,*User){
	var re bool
	var err error
	user:=&User{}
	user.Apikey=self.Apikey
	re=false
	if self.Apikey!=""{
		re,err =Orm_sys.Get(user)
		if err!=nil{
			utils.LogErr(err.Error())
		}
	}
	return re,user
}

//检查用户名，邮箱，手机号 是否合法
//参数 用户名 至少5位 不超过40位 大小写字母加数字加下划线
//返回 True 合法  Flase 不合法
func (self *User)CheckUserIsLegal() (bool,bool,bool){
	var ac,em,te bool
	ac=false
	em=false
	te=false
	if self.Account !=""{
		r,_:= regexp.Compile("^[a-zA-Z0-9_]{5,40}$")
		ac=r.MatchString(self.Account)
	}
	if self.Email !=""{
		r,_:= regexp.Compile("^([A-Za-z0-9_.-])+@([A-Za-z0-9_.-])+.([A-Za-z]{2,4})$")
		em=r.MatchString(self.Email)
	}
	if self.Telephone !=""{
		r,_:= regexp.Compile("^(13[0-9]|14[57]|15[0-35-9]|18[0,5-9]|(17[0-9]))[\\d]{8}$")
		te=r.MatchString(self.Telephone)
	}
	return ac,em,te
}

//检查用户名，邮箱，手机号 是否存在
//参数 无
//返回 True 可以使用  Flase不能使用（已存在或者为空）
func (self *User)CheckUserInfoExist() (bool,bool,bool){
	var ac,em,te bool
	var err error
	ac=false
	em=false
	te=false
	if self.Account!=""{
		ac,err =Orm_sys.Get(&User{Account:self.Account})
		if err!=nil{
			utils.LogErr(err.Error())
		}
	}
	if self.Email!=""{
		em,err =Orm_sys.Get(&User{Email:self.Email})
		if err!=nil{
			utils.LogErr(err.Error())
		}
	}
	if self.Telephone!=""{
		te,err =Orm_sys.Get(&User{Telephone:self.Telephone})
		if err!=nil{
			utils.LogErr(err.Error())
		}
	}
	return !ac,!em,!te
}

//添加新用户
//参数 无
//返回 True 成功  Flase 不成功
func (self *User) AddNewUser() (bool,error){
	ac,em,te := self.CheckUserIsLegal()
	fmt.Println(ac,em,te)
	if ac {
		if self.CheckUserAccountExist() {
			return false, errors.New("Account Is Exist")
		}
		if self.Password == "" {
			return false, errors.New("Password Can Not Be Empty")
		}
		if self.Email !=""{
			if !em {
				return false, errors.New("Email Is Not Legitimate")
			}
		}
		if self.Telephone !=""{
			if !te {
				return false, errors.New("Telephone Is Not Legitimate")
			}
		}
		if self.CheckUserEmailExist() {
			return false, errors.New("Email Is Exist")
		}

		if self.CheckUserTelephoneExist() {
			return false, errors.New("Telepone Is Exist")
		}
		self.Uuid = uuid.NewV4().String()
		self.Password = utils.StrToMD5v1(self.Password)
		self.Apikey = utils.StrToMD5v2(uuid.NewV1().String()[:20])
		self.Datashream_table = fmt.Sprintf("datastream_%s", uuid.NewV1())
		sql := fmt.Sprintf("CREATE TABLE `%s` %s", self.Datashream_table, datasheet_sql)
		_, err := Orm_data.Query(sql)
		if err != nil {
			self.Datashream_table = fmt.Sprintf("datastream_%s", uuid.NewV1())
			sql := fmt.Sprintf("CREATE TABLE `%s` %s", self.Datashream_table, datasheet_sql)
			_, err = Orm_data.Query(sql)
			if err != nil {
				utils.LogErr(err.Error())
				return false, err
			}
		}
		_, err = Orm_sys.Insert(self)
		if err != nil {
			_, err = Orm_sys.Insert(self)
			if err != nil {
				utils.LogErr(err.Error())
				return false, err
			}
		}
		return true, nil
	}
	if em || te {
		if self.Account !=""{
			return false,errors.New("Account Is Not Legitimate")
		}
		if self.Email !=""{
			if !em {
				return false, errors.New("Email Is Not Legitimate")
			}
		}
		if self.Telephone !=""{
			if !te {
				return false, errors.New("Telephone Is Not Legitimate")
			}
		}
		if self.CheckUserEmailExist(){
			return false, errors.New("Email Is Exist")
		}
		if self.CheckUserTelephoneExist() {
			return false, errors.New("Telepone Is Exist")
		}
		fmt.Println("123123")
		self.Uuid = uuid.NewV4().String()
		self.Account=fmt.Sprintf("User_%s%s",utils.StrToMD5v2(self.Uuid[:15])[:6],uuid.NewV1().String()[:6])
		self.Password = utils.StrToMD5v1(self.Password)
		self.Apikey = utils.StrToMD5v2(uuid.NewV1().String()[:20])
		self.Datashream_table = fmt.Sprintf("datastream_%s", uuid.NewV1())
		sql := fmt.Sprintf("CREATE TABLE `%s` %s", self.Datashream_table, datasheet_sql)
		_, err := Orm_data.Query(sql)
		if err != nil {
			self.Datashream_table = fmt.Sprintf("datastream_%s", uuid.NewV1())
			sql := fmt.Sprintf("CREATE TABLE `%s` %s", self.Datashream_table, datasheet_sql)
			_, err = Orm_data.Query(sql)
			if err != nil {
				utils.LogErr(err.Error())
				return false, err
			}
		}
		_, err = Orm_sys.Insert(self)
		if err != nil {
			_, err = Orm_sys.Insert(self)
			if err != nil {
				utils.LogErr(err.Error())
				return false, err
			}
		}
		return true, nil
	}
	return false,errors.New("String Null Or String Exist")
}


//验证用户登录消息是否正确
//参数 (account | email | telephone)  password
 //返回 1  user
 //     2  account | email | telephone
 //     3  password
 //		3  err
func (self *User) UserLoginCheck() (*User,bool,bool,error){
	user,_,err:= self.GetUser()
	if err !=nil{
		return user,false,false,err
	}
	if user.Password == utils.StrToMD5v1(self.Password){
		return user,true,true,nil
	}
	return &User{},true,false,errors.New("Password error")
}


func (self *User) GetUser() (*User,bool,error){
	var err error
	var re bool
	re=false
	user := &User{}
	if self.Account !=""{
		re,err=Orm_sys.Where("account=?",self.Account).Get(user)
		if err == nil{
			if !re{
				return user,re,errors.New("User Account Not Exist")
			}
			return user,re,err
		}
		return user,re,err
	}
	if self.Email !=""{
		re,err=Orm_sys.Where("email=?",self.Email).Get(user)
		if err == nil{
			if !re{
				return user,re,errors.New("User Email Not Exist")
			}
			return user,re,err
		}
		return user,re,err
	}
	if self.Telephone !=""{
		re,err=Orm_sys.Where("telephone=?",self.Telephone).Get(user)
		if err == nil{
			if !re{
				return user,re,errors.New("User Telephone Not Exist")
			}
			return user,re,err
		}
		return user,re,err
	}
	return user,re,errors.New("User Not Exist")
}



func AddNewUser(account string,email string,telephone string,password string) (bool,error){
	user := &User{Account:account,Email:email,Telephone:telephone,Password:password}
	return user.AddNewUser()
}

func CheckUserIsExist(account string,email string,telephone string) (bool,bool,bool){
	user := &User{Account:account,Email:email,Telephone:telephone}
	return user.CheckUserInfoExist()
}

func CheckUserLogin(account string,email string,telephone string,password string)(string,bool,bool,error){
	user := &User{Account:account,Email:email,Telephone:telephone,Password:password}
	userback,re1,re2,err :=user.UserLoginCheck()
	return userback.Uuid,re1,re2,err
}