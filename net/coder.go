/*
@Author: nullzz
@Date: 2021/11/8 8:42 下午
@Version: 1.0
@DEC:
*/
package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"go-net/net/iface"
	"io"
)

const (
	HeadLen       = 8 //8 byte，len+msgId
	MaxPacketSize = 1024
)

type DefaultCode struct{}

func (d *DefaultCode) Decode(session iface.ISession) (iface.IMessage, error) {
	headData := make([]byte, HeadLen) //8 byte,msgId+len
	if _, err := io.ReadFull(session.GetConn(), headData); err != nil {
		fmt.Println("read msg head error ", err)
		return nil, err
	}
	var dataLen uint32 = 0
	var id int32 = 0
	dataBuff := bytes.NewReader(headData)

	if err := binary.Read(dataBuff, binary.LittleEndian, &id); err != nil {
		fmt.Println("err ", err.Error())
		return nil, err
	}

	if err := binary.Read(dataBuff, binary.LittleEndian, &dataLen); err != nil {
		fmt.Println("err ", err.Error())
		return nil, err
	}

	var msg iface.IMessage
	if dataLen > 0 {
		if dataLen > MaxPacketSize {
			return nil, errors.New("too large msg data received")
		}
		body := make([]byte, dataLen)
		if _, err := io.ReadFull(session.GetConn(), body); err != nil {
			fmt.Println("read msg head error ", err)
			return nil, err
		}
		msg = NewMsgPackage(id, body)
	} else {
		msg = NewMsgPackage(id, nil)
	}
	return msg, nil
}

func (d *DefaultCode) Encode(msg iface.IMessage) ([]byte, error) {
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.LittleEndian, msg.GetId())
	binary.Write(bytesBuffer, binary.LittleEndian, msg.GetDataLen())
	if msg.GetDataLen() > 0 {
		binary.Write(bytesBuffer, binary.LittleEndian, msg.GetData())
	}
	//fmt.Println("bytesBuffer.Bytes()=", len(bytesBuffer.Bytes()))
	return bytesBuffer.Bytes(), nil
}
