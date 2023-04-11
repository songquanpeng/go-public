package handler

import (
	"fmt"
	"go-public/common"
	"net"
	"strconv"
)

func PublicPort(localPort, remotePort int) {
	fmt.Println("Go Public client started.")
	conn, err := net.Dial("tcp", common.ClientConfig.Host+":"+strconv.Itoa(common.ClientConfig.Port))
	if err != nil {
		fmt.Println("Failed to connect to server:", err.Error())
		return
	}
	fmt.Printf("Connected to %s:%d\n", common.ClientConfig.Host, common.ClientConfig.Port)
	err = sendHelloPacket(conn, remotePort)
	if err != nil {
		fmt.Println("Failed to send hello:", err.Error())
		return
	}
	fmt.Printf("Ready to public local port %d to %s:%d\n", localPort, common.ClientConfig.Host, remotePort)
	for {
		buf := make([]byte, ConnPacketSize) // do not move buf outside this for loop, it brings a bug
		n, err := conn.Read(buf)
		if err != nil {
			fmt.Println("Failed to read connection packet:", err.Error())
			fmt.Println("Abort.")
			return
		}
		if n != ConnPacketSize {
			fmt.Println("Warning: mismatched packet size.")
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
		fmt.Printf("Connection established: %s <-> %s\n", localConn.LocalAddr().String(), remoteConn.LocalAddr().String())
		go func() {
			_, err := remoteConn.Write(buf)
			if err != nil {
				fmt.Println("Failed to send connection packet:", err.Error())
				return
			}
			go forward(localConn, remoteConn)
			go forward(remoteConn, localConn)
		}()
	}
}
