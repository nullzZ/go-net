/*
@Author: nullzz
@Date: 2021/11/4 4:48 下午
@Version: 1.0
@DEC:
*/
package iface

type IServer interface {
	Start()
	Stop()
	Serve()
}
