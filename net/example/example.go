/*
@Author: nullzz
@Date: 2021/11/9 5:15 下午
@Version: 1.0
@DEC:
*/
package main

import (
	"fmt"
	"go-net/net"
	"go-net/net/iface"
)

func main() {
	//创建一个server句柄
	s := net.NewServer(net.WithPort(8080),
		net.WithSessionListener(&SessionListener{}),
		net.WithSendMsgLen(1024))
	//路由
	s.MsgHandlerMgr.Register(1, &TestHandler{})
	//连接监听
	s.Serve()
}

type SessionListener struct{}

func (SessionListener) CallOnConnStart(s iface.ISession) {
	fmt.Println("CallOnConnStart !!")
}

func (SessionListener) CallOnConnStop(s iface.ISession) {
	fmt.Println("CallOnConnStop !!")
}
