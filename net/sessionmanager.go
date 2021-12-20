/*
@Author: nullzz
@Date: 2021/11/4 8:40 下午
@Version: 1.0
@DEC:
*/
package net

import (
	"errors"
	"fmt"
	"go-net/net/iface"
	"sync"
)

var (
	ErrorConnIdSame = errors.New("ConnIdSameErr")
	ErrorConnIdNil  = errors.New("ConnIdNilErr")
)

type SessionManager struct {
	sessions map[int]iface.ISession
	connLock sync.RWMutex
}

func NewConnectionManager() *SessionManager {
	ctxMgr := &SessionManager{
		sessions: make(map[int]iface.ISession),
	}
	return ctxMgr
}

func (c *SessionManager) GetSession(id int) (iface.ISession, error) {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	conn, ok := c.sessions[id]
	if !ok {
		return nil, ErrorConnIdNil
	}
	return conn, nil
}

func (c *SessionManager) AddSession(session iface.ISession) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()
	id := session.GetId()
	if _, h := c.sessions[id]; h {
		fmt.Errorf("AddConn same id connId=%d \n", id)
		return ErrorConnIdSame
	}
	c.sessions[id] = session
	fmt.Println("AddConn add connId=", id)
	return nil
}

func (c *SessionManager) DelSession(id int) error {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	delete(c.sessions, id)
	fmt.Println("DelConn del connId=", id)
	return nil
}

func (c *SessionManager) Clear() {
	c.connLock.Lock()
	defer c.connLock.Unlock()

	for connID, session := range c.sessions {
		session.Stop()
		delete(c.sessions, connID)
	}
	fmt.Println("Clear All Connections successfully: conn num = ", c.Len())
}

func (c *SessionManager) Len() int {
	c.connLock.RLock()
	defer c.connLock.RUnlock()
	return len(c.sessions)
}
