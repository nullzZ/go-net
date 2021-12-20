/*
@Author: nullzz
@Date: 2021/11/4 8:07 下午
@Version: 1.0
@DEC:
*/
package iface

type IMsgHandlerManager interface {
	Run()
	Stop()
	Register(msgId int32, handler IHandler)
	DoHandler(session ISession, req IMessage) error
}
