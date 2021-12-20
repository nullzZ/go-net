/*
@Author: nullzz
@Date: 2021/11/4 4:53 下午
@Version: 1.0
@DEC:
*/
package net

import (
	"fmt"
	"go-net/net/iface"
	"net"
)

type Server struct {
	Config          *ServerConfig
	sessionMgr      iface.ISessionManager
	MsgHandlerMgr   iface.IMsgHandlerManager
	coder           iface.ICoder
	stopChan        chan interface{}
	sessionListener iface.ISessionListener
}

func NewServer(opts ...Option) *Server {
	server := &Server{
		Config:     DefaultServerConfig(),
		sessionMgr: NewConnectionManager(),
		coder:      &DefaultCode{},
		stopChan:   make(chan interface{}),
	}
	for _, opt := range opts {
		opt(server)
	}
	if server.MsgHandlerMgr == nil {
		server.MsgHandlerMgr = NewMsgHandlerManager(server.Config.WorkerLen, server.Config.WorkerPoolCount)
	}

	return server
}

func (s *Server) Start() {
	fmt.Printf("Server start ip=%s,port=%d \n", s.Config.IP, s.Config.Port)
	// 获取一个TCP的Addr
	addr, err := net.ResolveTCPAddr(s.Config.Network, fmt.Sprintf("%s:%d", s.Config.IP, s.Config.Port))
	if err != nil {
		panic(err)
	}
	//2 监听服务器地址
	listener, err := net.ListenTCP(s.Config.Network, addr)
	if err != nil {
		panic(err)
	}

	// 启动worker处理
	go s.MsgHandlerMgr.Run()

	fmt.Printf("Server start success ip=%s,port=%d \n", s.Config.IP, s.Config.Port)

	cID := 0
	for {
		// 阻塞等待客户端建立连接请求
		conn, err := listener.AcceptTCP()
		if err != nil {
			fmt.Println("Accept err ", err)
			continue
		}

		// 设置服务器最大连接控制,如果超过最大连接，那么则关闭此新的连接
		if s.sessionMgr.Len() >= s.Config.MaxConn {
			fmt.Println("Accept err connect max")
			conn.Close()
			continue
		}

		cID++
		session := NewSession(s.Config, cID, conn, s.sessionMgr, s.MsgHandlerMgr, s.coder, s.sessionListener) //创建会话
		s.sessionMgr.AddSession(session)                                                                      //存储连接会话

		// 启动当前链接的处理业务
		go session.Run()
	}

}

func (s *Server) Stop() {
	s.sessionMgr.Clear()
	s.MsgHandlerMgr.Stop()
	s.stopChan <- 1
}

func (s *Server) Serve() {
	s.Start()
	select {
	case <-s.stopChan:
		return
	}
}
