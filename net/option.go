/*
@Author: nullzz
@Date: 2021/11/4 4:58 下午
@Version: 1.0
@DEC:
*/
package net

import "go-net/net/iface"

type Option func(s *Server)

func WithIp(ip string) Option {
	return func(s *Server) {
		s.Config.IP = ip
	}
}

func WithPort(port int) Option {
	return func(s *Server) {
		s.Config.Port = port
	}
}

func WithNetwork(network string) Option {
	return func(s *Server) {
		s.Config.Network = network
	}
}

func WithCoder(coder iface.ICoder) Option {
	return func(s *Server) {
		s.coder = coder
	}
}

// max connection
func WithMaxConn(maxConn int) Option {
	return func(s *Server) {
		s.Config.MaxConn = maxConn
	}
}

func WithWorkerLen(workerLen int) Option {
	return func(s *Server) {
		s.Config.WorkerLen = workerLen
	}
}

func WithWorkerPoolCount(workerCount int) Option {
	return func(s *Server) {
		s.Config.WorkerPoolCount = workerCount
	}
}

func WithSendMsgLen(sendMsgLen int) Option {
	return func(s *Server) {
		s.Config.SendMsgLen = sendMsgLen
	}
}

func WithSessionListener(listener iface.ISessionListener) Option {
	return func(s *Server) {
		s.sessionListener = listener
	}
}
