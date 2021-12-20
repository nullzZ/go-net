/*
@Author: nullzz
@Date: 2021/11/22 8:23 下午
@Version: 1.0
@DEC:
*/
package iface

type IMessage interface {
	GetId() int32
	GetDataLen() uint32
	GetData() []byte

	SetId(id int32)
	//SetDataLen(len uint32)
	SetData(data []byte)
}
