package KEDP

import (


	"github.com/ncoskay/goserver/iotserver/Protocols/KEDP/controller"
	"net"
)

func KEDP_Core(conn net.Conn,buf []byte, t *int) {
	dmap, ok := controller.Depack(buf)
	if !ok {
		*t = 0
		return
	}
	//fmt.Println(dmap)
	controller.Interpreter(conn,dmap.ProType(), dmap.MsgBody())
	*t = dmap.ConnTime()
	//conn.Write([]byte("123"))
}
