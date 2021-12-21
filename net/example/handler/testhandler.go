/*
@Author: nullzz
@Date: 2021/11/18 3:08 下午
@Version: 1.0
@DEC:
*/
package handler

import (
	"fmt"
	"go-net/net"
	"go-net/net/iface"
)

type TestHandler struct{}

func (TestHandler) Handler(session iface.ISession, msg interface{}) {
	fmt.Println("@@@@@@@@@@@@@@")
	session.SendMsgBuff(net.NewMsgPackage(1, nil))
}
