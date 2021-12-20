/*
@Author: nullzz
@Date: 2021/11/18 2:56 下午
@Version: 1.0
@DEC:
*/
package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}
	for {
		bytesBuffer := bytes.NewBuffer([]byte{})
		binary.Write(bytesBuffer, binary.LittleEndian, int32(1))
		binary.Write(bytesBuffer, binary.LittleEndian, int32(0))
		fmt.Println("bytesBuffer.Bytes()=", len(bytesBuffer.Bytes()))
		conn.Write(bytesBuffer.Bytes())
		time.Sleep(2 * time.Second)

	}

	//conn.Close()
	//select {}
}
