/*
@Author: nullzz
@Date: 2021/11/5 2:59 下午
@Version: 1.0
@DEC:
*/
package iface

type ICoder interface {
	Encode(msg IMessage) ([]byte, error)
	Decode(session ISession) (IMessage, error)
}
