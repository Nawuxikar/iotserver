package controller

import (
	"github.com/ncoskay/goserver/utils"
)


type Data_map struct {
	proType   uint8
	check_tag uint8
	msgType   uint8
	qos       uint8
	contime   int
	msglength uint32
	msgbody   []byte
}

func (dm *Data_map) ProType() uint8 {
	return dm.proType
}
func (dm *Data_map) CheckTag() uint8 {
	return dm.check_tag
}
func (dm *Data_map) MsgType() uint8 {
	return dm.msgType
}
func (dm *Data_map) Qos() uint8 {
	return dm.qos
}
func (dm *Data_map) ConnTime() int {
	return dm.contime
}
func (dm *Data_map) MsgLength() uint32 {
	return dm.msglength
}
func (dm *Data_map) MsgBody() []byte {
	return dm.msgbody
}

func Depack(buf []byte) (*Data_map, bool) {
	dmap := new(Data_map)
	var ok bool = true
	//utils.LogDebug(string(buf))
	var msgl int

	leng := len(buf)
	if leng >= 3 {
		//fmt.Println(leng)   //3
		msgl, dmap.msglength = getLegth(buf)
		//fmt.Println(msgl)
		//fmt.Println(dmap.msglength)
		if int32(leng) == int32(dmap.msglength)+int32(msgl)+2 {
			dmap.proType = uint8(buf[0] >> 3)
			dmap.msgType = uint8(buf[0] & 0x01)
			dmap.check_tag = uint8((buf[0] >> 1) & 0x03)
			dmap.qos = uint8(buf[1] >> 6)
			dmap.contime = getcontime(int(buf[1] & 0x07))
			dmap.msgbody = buf[2+msgl : leng]
		} else {
			utils.LogWarn("message data length is not right")
			ok = false
		}
	} else {
		utils.LogWarn("less message data < 3")
		ok = false
	}
	//fmt.Println(dmap)
	return dmap, ok
}
func getcontime(t int) int {
	x := int(0)
	switch t {
	case 0:
		x = 0
		break
	case 1:
		x = 5
		break
	case 2:
		x = 15
		break
	case 3:
		x = 30
		break
	case 4:
		x = 60
		break
	case 5:
		x = 120
		break
	case 6:
		x = 300
		break
	case 7:
		x = 600
		break
	}
	return x
}

func getLegth(buf []byte) (int, uint32) {
	msgl := int(0)
	msgleng := uint32(0)
	if buf[2]>>7 == 0x01 {
		if buf[3]>>7 == 0x01 {
			if buf[4]>>7 == 0x01 {
				msgl = 4
				msgleng = uint32(buf[2]&0x7F + (buf[3]&0x7F)<<7 + (buf[4]&0x7F)<<7 + (buf[5]&0x7F)<<7)
			} else {
				msgl = 3
				msgleng = uint32(buf[2]&0x7F + (buf[3]&0x7F)<<7 + (buf[4]&0x7F)<<7)
			}
		} else {
			msgl = 2
			msgleng = uint32(buf[2]&0x7F + (buf[3]&0x7F)<<7)
		}
	} else {
		msgl = 1
		msgleng = uint32(buf[2] & 0x7F)
	}
	return msgl, msgleng
}
