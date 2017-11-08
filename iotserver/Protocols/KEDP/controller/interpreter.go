package controller

import (
	//"../../../utils"
	"net"
	"encoding/json"
	//"github.com/ncoskay/goserver/models"
)

const  (
	CONN_REQ	=	1
	CONN_RESP	=	2
	PUSH_DATA	=	3
	SAVE_DATA	=	8
	GET_DATA	=	10
	PING		=	13
	PING_RE		=	14
	PUB_DATA	=	15
	PUB_RE		=	16
	SUB_DATA	=	17
	SUB_RE		=	18
	UNSUB_DATA	=	19
	UBSUB_RE	=	20
	REG_REQ		=	21
	REG_RE		=	22
)
/*
原始JSON数据 结构
{"data":[{""}]}


连接请求 数据结构
{""}

 */

type Msg_body struct {
	UUID 	string				`json:"uuid"`
	SN   	string				`json:"sn"`
	DATA	map[string]string	`json:"data"`
	CMD		[]string			`json:"cmd"`
}

type Msg_data struct {
	Key		string	`json:"key"`
	Value	string	`json:"value"`
}

type DevInfo_T struct{
	SN 		string
	Time 	int
}

var DevConnListMap map[net.Conn]DevInfo_T



// 连接
func func_ConnEvent(conn net.Conn,msg *Msg_body){
	if msg.SN==""{

	}
}
// 上传
func func_PushEvent(conn net.Conn,msg *Msg_body){

}
// 保存
func func_SaveEvent(conn net.Conn,msg *Msg_body){

}
// 请求
func func_GetEvent(conn net.Conn,msg *Msg_body){

}
//发布
func func_PubEvent(conn net.Conn,msg *Msg_body){

}
//订阅
func func_SubEvent(conn net.Conn,msg *Msg_body){

}
//心跳
func func_PingEvent(conn net.Conn,msg *Msg_body){

}
//注册
func func_RegEvent(conn net.Conn,msg *Msg_body){

}




func init(){
	DevConnListMap=make(map[net.Conn]DevInfo_T)
}

func Interpreter(conn net.Conn,protype uint8,msg []byte){
	Msg :=&Msg_body{}
	_ =json.Unmarshal(msg,Msg)

	switch protype {
	case CONN_REQ: //连接
		func_ConnEvent(conn,Msg)

	case PUSH_DATA: //上传数据
		func_PushEvent(conn,Msg)

	case SAVE_DATA:
		func_SaveEvent(conn,Msg)

	case GET_DATA:
		func_GetEvent(conn,Msg)

	case PUB_DATA:
		func_PubEvent(conn,Msg)

	case SUB_DATA:
		func_SubEvent(conn,Msg)

	case REG_REQ://设备注册
		func_RegEvent(conn,Msg)

	case PING: //心跳保持
		func_PingEvent(conn,Msg)

	}

	conn.Write(msg);
}