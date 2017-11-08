package main

import (
	"github.com/ncoskay/goserver/iotserver/Protocols/KEDP"
	"github.com/ncoskay/goserver/iotserver/server"
	"github.com/ncoskay/goserver/utils"
	"fmt"
)

func main() {
	server_init("./conf/config.conf")
}

func server_init(configpath string) {
	//var c chan int

	// 初始化 日志  配置文件
	err :=utils.Config_init(configpath)
	if err!=nil{
		fmt.Println("config model error")
		return
	}


	// 建立服务器监听
	host := utils.Config.String("host")
	server.Server_start(host, 4096, KEDP.KEDP_Core)
	//go server.Server_start("localhost:8899",4096,Protocol.KEDP_Core)

	//阻塞
	//c <- 1

}
