package handler

import (
	"fmt"
	"go-public/common"
	"net"
	"strconv"
)

func PublicPort(localPort, remotePort int) {
	conn, err := net.Dial("tcp", common.ClientConfig.Host+":"+strconv.Itoa(common.ClientConfig.Port))
	if err != nil {
		fmt.Println("Failed to connect to server:", err.Error())
		return
	}
	fmt.Println("Connected to server with remote port", remotePort)
	err = sendHelloPacket(conn, remotePort)
	if err != nil {
		fmt.Println("Failed to send hello:", err.Error())
		return
	}
	fmt.Println("Hello packet sent")
	buf := make([]byte, ConnPacketSize)
	for {
		_, err := conn.Read(buf)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		// Should be a connection packet
		localConn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(localPort))
		if err != nil {
			fmt.Println("Failed to connect to local server:", err.Error())
			continue
		}
		remoteConn, err := net.Dial("tcp", common.ClientConfig.Host+":"+strconv.Itoa(common.ClientConfig.Port))
		if err != nil {
			fmt.Println("Failed to connect to remote server:", err.Error())
			continue
		}
		go func() {
			_, err := remoteConn.Write(buf)
			if err != nil {
				fmt.Println("Failed to send connection packet:", err.Error())
			}
		}()
		go forward(localConn, remoteConn)
		go forward(remoteConn, localConn)
	}
}
