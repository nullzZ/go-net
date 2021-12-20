/*
@Author: nullzz
@Date: 2021/11/22 8:13 下午
@Version: 1.0
@DEC:
*/
package net

type Message struct {
	DataLen uint32
	Id      int32
	Data    []byte
}

func NewMsgPackage(id int32, data []byte) *Message {
	return &Message{
		DataLen: uint32(len(data)),
		Id:      id,
		Data:    data,
	}
}

func (m *Message) GetId() int32 {
	return m.Id
}

func (m *Message) GetDataLen() uint32 {
	return m.DataLen
}

func (m *Message) GetData() []byte {
	return m.Data
}

func (m *Message) SetId(id int32) {
	m.Id = id
}

func (m *Message) SetData(data []byte) {
	m.Data = data
}
