package handler

import (
	"fmt"
	"go-public/common"
	"io"
	"net"
	"os"
)

func ServeForever() {
	fmt.Println("Go Public server started.")
	addr := net.TCPAddr{
		IP:   nil,
		Port: common.ServerConfig.Port,
	}
	listener, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	fmt.Printf("Server listening on port %d. \nWaiting for connections...\n", common.ServerConfig.Port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		fmt.Println("Client connection established:", conn.RemoteAddr().String())
		go handleClientConnection(conn)
	}
}

func handleClientConnection(conn net.Conn) {
	defer conn.Close()
	okay, port := recvHello(conn)
	if !okay {
		return
	}
	addr := net.TCPAddr{
		IP:   nil,
		Port: port,
	}
	listener, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		println(err.Error())
		return
	}
	for {
		otherConn, err := listener.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		fmt.Printf("[%d] connection %p established\n", port, otherConn)
		go func() {
			_, err = io.Copy(conn, otherConn)
			if err != nil {
				println(err.Error())
			}
			defer otherConn.Close()
		}()
		go func() {
			_, err = io.Copy(otherConn, conn)
			if err != nil {
				println(err.Error())
			}
			defer otherConn.Close()
		}()
	}
}
