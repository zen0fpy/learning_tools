package network

import (
	"net"
	"bufio"
	"bytes"
	"encoding/binary"
	"fmt"
)

type TcpClient struct {
	Tag    string
	Conn   net.Conn
	reader *bufio.Reader
	writer *bufio.Writer
	len    int32
}

func NewTcpClint(conn *net.TCPConn) *TcpClient {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	return &TcpClient{Conn: conn, reader: reader, writer: writer, len: 4}
}

func (c *TcpClient) Read() {
	for {
		// 读取消息的长度
		lengthByte, err := c.reader.Peek(int(c.len))
		if len(lengthByte) == 0 || err != nil {
			fmt.Println(err)
			continue
		}
		lengthBuff := bytes.NewBuffer(lengthByte)
		var length int32
		err = binary.Read(lengthBuff, binary.BigEndian, &length)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if int32(c.reader.Buffered()) < length+c.len {
			fmt.Println(int32(c.reader.Buffered()))
			continue
		}
		// 读取消息真正的内容
		pack := make([]byte, int(c.len+length))
		_, err = c.reader.Read(pack)
		if err != nil {
		}
		//获取到的数据
		fmt.Println(pack[4:])
	}
}

