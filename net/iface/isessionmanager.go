/*
@Author: nullzz
@Date: 2021/11/4 7:59 下午
@Version: 1.0
@DEC:
*/
package iface

type ISessionManager interface {
	GetSession(id int) (ISession, error)
	AddSession(session ISession) error
	DelSession(id int) error
	Clear()
	Len() int
}
