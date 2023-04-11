package handler

import (
	"fmt"
	"go-public/common"
	"net"
)

var tokenSize = 16

var (
	HelloPacket = []byte("\x00")[0]
	ConnPacket  = []byte("\x01")[0]
)

var (
	HelloPacketSize         = 1 + 2 + tokenSize
	ConnPacketSize          = 1 + tokenSize
	ServerMaxRecvPacketSize = HelloPacketSize
)

func sendHelloPacket(conn net.Conn, remotePort int) error {
	port := uint16(remotePort)
	var tokenBytes = common.Token2Bytes(common.ClientConfig.Token)
	var buf = make([]byte, HelloPacketSize)
	buf[0] = HelloPacket
	buf[1] = byte(port >> 8)
	buf[2] = byte(port)
	copy(buf[3:], tokenBytes)
	_, err := conn.Write(buf)
	return err
}

func parseHelloPacket(buf []byte) (ok bool, port int) {
	port = int(buf[1])<<8 + int(buf[2])
	if port == common.ServerConfig.Port {
		fmt.Println("Invalid port:", port)
		return false, 0
	}
	token := common.Bytes2Token(buf[3 : 3+tokenSize])
	if token == common.ServerConfig.Token {
		return true, port
	}
	fmt.Println("Invalid token:", token)
	return false, 0
}

func SendConnPacket(conn net.Conn, token string) error {
	var buf = make([]byte, ConnPacketSize)
	buf[0] = ConnPacket
	copy(buf[1:], common.Token2Bytes(token))
	_, err := conn.Write(buf)
	return err
}
