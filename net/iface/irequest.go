/*
@Author: nullzz
@Date: 2021/11/6 3:43 下午
@Version: 1.0
@DEC:
*/
package iface

type IRequest interface {
	GetMessage() IMessage
	GetSession() ISession
}
