/*
@Author: nullzz
@Date: 2021/11/4 4:55 下午
@Version: 1.0
@DEC:
*/
package net

import "runtime"

type ServerConfig struct {
	IP              string
	Port            int
	Network         string //"tcp", "tcp4", "tcp6"
	MaxConn         int    //max connect
	WorkerLen       int    //worker队列长度
	WorkerPoolCount int    //worker数量
	SendMsgLen      int
}

func DefaultServerConfig() *ServerConfig {
	c := &ServerConfig{
		IP:              "127.0.0.1",
		Port:            8080,
		Network:         "tcp",
		MaxConn:         128,
		WorkerLen:       1024,
		WorkerPoolCount: runtime.NumCPU(),
		SendMsgLen:      64,
	}
	return c
}
