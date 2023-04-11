package handler

import (
	"fmt"
	"go-public/common"
	"io"
	"net"
	"strconv"
)

func PublicPort(localPort, remotePort int) {
	localConn, err := net.Dial("tcp", "localhost:"+strconv.Itoa(localPort))
	if err != nil {
		fmt.Println("Failed to connect to local port:", err.Error())
		return
	}
	fmt.Println("Connected to local port", localPort)
	conn, err := net.Dial("tcp", common.ClientConfig.Host+":"+strconv.Itoa(common.ClientConfig.Port))
	if err != nil {
		fmt.Println("Failed to connect to server:", err.Error())
		return
	}
	fmt.Println("Connected to server with remote port", remotePort)
	err = sendHello(conn, remotePort)
	if err != nil {
		fmt.Println("Failed to send hello:", err.Error())
		return
	}
	go func() {
		_, err = io.Copy(localConn, conn)
		if err != nil {
			println(err.Error())
		}
		defer localConn.Close()
	}()
	func() {
		_, err = io.Copy(conn, localConn)
		if err != nil {
			println(err.Error())
		}
		defer localConn.Close()
	}()
	//go forward(localConn, conn)
	//forward(conn, localConn)
	fmt.Println("Connection broken")
}
