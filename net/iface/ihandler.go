/*
@Author: nullzz
@Date: 2021/11/5 3:39 下午
@Version: 1.0
@DEC:
*/
package iface

type IHandler interface {
	Handler(session ISession, msg interface{})
}
