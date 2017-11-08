package main

import (
	. "github.com/ncoskay/goserver/iotserver/Protocols/KEDP/controller"
	"fmt"
	"encoding/json"
)


func main(){
	var arr  []string
	arr=append(arr,"1")
	arr=append(arr,"234")
	var data map[string]string
	data=make(map[string]string)
	data["x"]="123"
	data["y"]="111"
	d:=&Msg_body{UUID:"aaaaadf",SN:"11111",CMD:arr,DATA:data}
	xx,err:=json.Marshal(d)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(string(xx))
}