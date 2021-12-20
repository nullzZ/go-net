/*
@Author: nullzz
@Date: 2021/11/18 5:52 下午
@Version: 1.0
@DEC:
*/
package iface

//Hook函数
type ISessionListener interface {
	CallOnConnStart(s ISession)
	CallOnConnStop(s ISession)
	//SetOnConnStart(func(IConnection)) 		//设置该Server的连接创建时Hook函数
	//SetOnConnStop(func(IConnection)) 		//设置该Server的连接断开时的Hook函数
}
