/*
@Author: nullzz
@Date: 2021/11/6 11:55 上午
@Version: 1.0
@DEC:
*/
package net

import (
	"context"
	"fmt"
	"go-net/net/iface"
	"net"
	"sync"
)

type Session struct {
	propertyLock  sync.RWMutex
	sLock         sync.RWMutex
	SessionId     int
	property      map[string]interface{}
	MsgHandlerMgr iface.IMsgHandlerManager
	ConnMgr       iface.ISessionManager
	decode        iface.ICoder
	ctx           context.Context
	cancel        context.CancelFunc
	conn          *net.TCPConn
	msgChan       chan iface.IMessage
	//buffer chan
	msgBuffChan   chan iface.IMessage
	msgBuffCancel chan bool
	cfg           *ServerConfig
	listener      iface.ISessionListener
	Status        int
}

const (
	SessionStatusCon   = 0 //connected
	SessionStatusClose = 1 //close
)

func NewSession(cfg *ServerConfig, id int, conn *net.TCPConn, connMgr iface.ISessionManager,
	msgHMgr iface.IMsgHandlerManager, decode iface.ICoder, listener iface.ISessionListener) *Session {
	session := &Session{
		SessionId:     id,
		conn:          conn,
		MsgHandlerMgr: msgHMgr,
		ConnMgr:       connMgr,
		property:      make(map[string]interface{}),
		decode:        decode,
		cfg:           cfg,
		msgChan:       make(chan iface.IMessage),
		msgBuffChan:   make(chan iface.IMessage, cfg.SendMsgLen),
		msgBuffCancel: make(chan bool),
		Status:        SessionStatusCon,
		listener:      listener,
	}
	return session
}

func (s *Session) Run() {
	s.ctx, s.cancel = context.WithCancel(context.Background())
	go s.StartReader() //读数据
	go s.StartWriter() //发数据
	if s.listener != nil {
		s.listener.CallOnConnStart(s)
	}
}

func (s *Session) Stop() {
	s.sLock.Lock()
	defer s.sLock.Unlock()
	s.Status = SessionStatusClose
	s.msgBuffCancel <- true
	s.cancel()
	close(s.msgChan)
	close(s.msgBuffChan)
	if s.listener != nil {
		s.listener.CallOnConnStop(s)
	}
}

func (s *Session) SetVal(key string, val interface{}) {
	s.propertyLock.Lock()
	defer s.propertyLock.Unlock()
	s.property[key] = val
}

func (s *Session) GetVal(key string) (interface{}, bool) {
	s.propertyLock.RLock()
	defer s.propertyLock.RUnlock()
	v, ok := s.property[key]
	if !ok {
		return nil, false
	}
	return v, true
}

func (s *Session) GetId() int {
	return s.SessionId
}

func (s *Session) GetConn() net.Conn {
	return s.conn
}

func (s *Session) StartReader() {
	fmt.Println("StartReader session=", s.SessionId)
	defer s.Stop()
	defer fmt.Println(s.RemoteAddr().String(), "session=", s.SessionId, " conn Reader exit!")
	for {
		select {
		case <-s.ctx.Done():
			return
		default:
			msg, err := s.decode.Decode(s)
			if err != nil {
				fmt.Errorf("StartReader err=%v", err)
				return
			}
			err = s.MsgHandlerMgr.DoHandler(s, msg)
			if err != nil {
				fmt.Errorf("StartReader err=%v", err)
				return
			}
		}
	}

}

func (s *Session) StartWriter() {
	fmt.Println("StartWriter session=", s.SessionId)
	defer fmt.Println(s.RemoteAddr().String(), "session=", s.SessionId, " conn Writer exit!")
	for {
		select {
		case <-s.msgBuffCancel:
			return
		case data, ok := <-s.msgBuffChan:
			if !ok { //是否关闭
				return
			}
			bb, err := s.decode.Encode(data)
			if err != nil {
				fmt.Println("StartReader err=", err)
				return
			}
			if _, err := s.conn.Write(bb); err != nil {
				fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
				return
			}
		case data, ok := <-s.msgChan:
			if !ok {
				return
			}
			bb, err := s.decode.Encode(data)
			if err != nil {
				fmt.Println("StartReader err=", err)
				return
			}
			if _, err := s.conn.Write(bb); err != nil {
				fmt.Println("Send Buff Data error:, ", err, " Conn Writer exit")
				return
			}
		}
	}
}

func (s *Session) AddListener(listener iface.ISessionListener) {
	s.listener = listener
}

func (s *Session) RemoteAddr() net.Addr {
	return s.conn.RemoteAddr()
}

func (s *Session) SendMsg(data iface.IMessage) {
	s.sLock.RLock()
	if s.Status == SessionStatusClose { //session close
		return
	}
	s.sLock.RUnlock()

	s.msgChan <- data
}

func (s *Session) SendMsgBuff(data iface.IMessage) {
	s.sLock.RLock()
	if s.Status == SessionStatusClose {
		return
	}
	s.sLock.RUnlock()

	s.msgBuffChan <- data
}
