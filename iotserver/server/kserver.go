package server

import (
	"net"
	"time"

	"github.com/ncoskay/goserver/utils"
)

type Server_t struct {
	Conn_fd 	net.Conn
	Dtime		chan int

}


func Server_start(host string, bufsize int, profunc func(net.Conn,[]byte, *int)) {
	netlisten, err := net.Listen("tcp", host)
	utils.LogCheckErr(err)
	defer netlisten.Close()
	utils.LogInfo(host + "  Waiting for conn")
	for {
		conn, err := netlisten.Accept()
		if err != nil {
			continue
		}
		//utils.LogInfo(conn.RemoteAddr().String())
		utils.LogInfo(conn.RemoteAddr().String() + " connection sucess!")
		go handleConn(conn, bufsize, profunc)
	}

}


/*   使用socket自带 SetDeadline函数完成 超时断开     */
func handleConn(conn net.Conn, bufsize int, profunc func(net.Conn,[]byte, *int)) {
	buffer := make([]byte, bufsize)
	timechan := make(chan int)
	quit := make(chan int)
	var t int
	t = 10
	var tpr *int
	tpr = &t
	conn.SetDeadline(time.Now().Add(time.Duration(t) * time.Second))
	//心跳控制
	go HeartBeating(conn, timechan,quit)
	//发送
	sendchan := make(chan []byte)
	go HanderSend(conn, sendchan)

	for {
		n, err := conn.Read(buffer)
		if err != nil {
			utils.LogWarn(" Connection error: " + err.Error())
			conn.Close()
			close(timechan)
			quit<-0
			//close(sendchan)
			return
		}
		data := (buffer[:n])
		//utils.LogDebug(string(data))
		profunc(conn,data, tpr)
		go GravelChannel(*tpr, timechan)
	}
	defer conn.Close()
}
func HeartBeating(conn net.Conn, timechan chan int,quit chan int) {
	for {
		select {
		case t := <-timechan:
			conn.SetDeadline(time.Now().Add(time.Duration(t) * time.Second))
		case <-quit:
			return
		}
	}
}
func GravelChannel(n int, mess chan int) {
	mess <- n
}

func HanderSend(conn net.Conn, dataChan chan []byte) {
	for {
		select {
		case data := <-dataChan:
			conn.Write(data)
		}
	}
}

///* 使用定时器计时心跳 完成 超时断开 */
//func handleConn(conn net.Conn,bufsize int){
//	buffer := make([]byte,bufsize)
//	timechan :=make(chan int)
//	tout := time.NewTimer(time.Second*30)
//	go HeartBeating(conn,tout,timechan)
//	for {
//		n,err := conn.Read(buffer)
//		if err !=nil{
//			utils.LogErr(conn.RemoteAddr().String(), " connection error: ", err)
//			return
//		}
//		data :=(buffer[:n])
//		utils.LogDebug(string(data))
//		go GravelChannel(10,timechan)
//	}
//	defer conn.Close()
//}
//func HeartBeating(conn net.Conn,tout *time.Timer,timechan chan int){
//	for{
//		select{
//		case t := <- timechan:
//			//Log("time=",t)
//			tout.Reset(time.Second*time.Duration(t))
//		case <-tout.C:
//			utils.LogInfo("sock close!!!!!")
//			conn.Close()
//		}
//	}
//}
//func GravelChannel(n int,mess chan int){
//	mess <- n
//}
