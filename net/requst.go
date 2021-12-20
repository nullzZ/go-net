/*
@Author: nullzz
@Date: 2021/11/6 3:41 下午
@Version: 1.0
@DEC:
*/
package net

import "go-net/net/iface"

type Request struct {
	Msg     iface.IMessage
	Session iface.ISession
}

func NewRequest(session iface.ISession, msg iface.IMessage) *Request {
	return &Request{
		Msg:     msg,
		Session: session,
	}

}

func (r *Request) GetMessage() iface.IMessage {
	return r.Msg
}

func (r *Request) GetSession() iface.ISession {
	return r.Session
}
