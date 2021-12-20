/*
@Author: nullzz
@Date: 2021/11/4 8:39 下午
@Version: 1.0
@DEC:
*/
package net

import (
	"fmt"
	"go-net/net/iface"
	"reflect"
)

type MsgHandlerManager struct {
	handlers        map[int32]iface.IHandler
	workerQueue     []chan iface.IRequest //service pool
	workerLen       int                   //queue len
	workerPoolCount int                   //workerPool count
	close           chan interface{}
}

func NewMsgHandlerManager(workerLen, workerPoolCount int) *MsgHandlerManager {
	msgHandlerMgr := &MsgHandlerManager{
		handlers:        make(map[int32]iface.IHandler),
		workerPoolCount: workerPoolCount,
		workerLen:       workerLen,
		workerQueue:     make([]chan iface.IRequest, workerPoolCount),
	}
	return msgHandlerMgr
}

func (m *MsgHandlerManager) Register(id int32, handler iface.IHandler) {
	m.handlers[id] = handler
	t := reflect.TypeOf(handler)
	fmt.Println("MsgHandlerManager Register id=", id, "handler=", t.String())
}

func (m *MsgHandlerManager) DoHandler(session iface.ISession, msg iface.IMessage) error {
	req := NewRequest(session, msg)
	index := session.GetId() % m.workerPoolCount
	m.workerQueue[index] <- req
	return nil
}

func (m *MsgHandlerManager) startWorker(workerID int, taskQueue chan iface.IRequest) {
	fmt.Println("MsgHandlerWorker ID = ", workerID, " is started.")
	defer fmt.Println("MsgHandlerManager stop workerID=", workerID)
	for {
		select {
		case <-m.close:
			return
		case request, ok := <-taskQueue:
			if ok {
				msg := request.GetMessage()
				h, ok := m.handlers[msg.GetId()]
				if !ok {
					fmt.Errorf("handler nil apiId=%d", msg.GetId())
				} else {
					h.Handler(request.GetSession(), msg.GetData())
					//todo slow log
				}
			}
		}
	}
}

func (m *MsgHandlerManager) Run() {
	for i := 0; i < m.workerPoolCount; i++ {
		m.workerQueue[i] = make(chan iface.IRequest, m.workerLen)
		go m.startWorker(i, m.workerQueue[i])
	}
}

func (m *MsgHandlerManager) Stop() {
	m.close <- 1
}
