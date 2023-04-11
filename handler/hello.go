package handler

import (
	"fmt"
	"go-public/common"
	"net"
)

func sendHello(conn net.Conn, remotePort int) error {
	port := uint16(remotePort)
	var tokenBytes = []byte(common.ClientConfig.Token)
	var tokenLen = uint8(len(tokenBytes))
	var buf = make([]byte, 2+1+tokenLen)
	buf[0] = byte(port >> 8)
	buf[1] = byte(port)
	buf[2] = tokenLen
	copy(buf[3:], tokenBytes)
	_, err := conn.Write(buf)
	return err
}

func recvHello(conn net.Conn) (ok bool, port int) {
	buf := make([]byte, 256)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return false, 0
	}
	port = int(buf[0])<<8 + int(buf[1])
	if port == common.ServerConfig.Port {
		fmt.Println("Invalid port:", port)
		return false, 0
	}
	tokenLen := int(buf[2])
	token := string(buf[3 : 3+tokenLen])
	if token == common.ServerConfig.Token {
		return true, port
	}
	fmt.Println("Invalid token:", token)
	return false, 0
}
