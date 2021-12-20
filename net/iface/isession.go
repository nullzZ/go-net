/*
@Author: nullzz
@Date: 2021/11/6 11:55 上午
@Version: 1.0
@DEC:
*/
package iface

import "net"

type ISession interface {
	GetId() int
	GetConn() net.Conn
	SetVal(key string, val interface{})
	GetVal(key string) (interface{}, bool)
	Run()
	Stop()
	SendMsg(data IMessage)
	SendMsgBuff(data IMessage)
}
