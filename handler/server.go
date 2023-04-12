package handler

import (
	"fmt"
	"go-public/common"
	"net"
	"os"
)

func ServeForever() {
	fmt.Printf("Go Public server started at port %d.\n", common.ServerConfig.Port)
	addr := net.TCPAddr{
		IP:   nil,
		Port: common.ServerConfig.Port,
	}
	listener, err := net.ListenTCP("tcp", &addr)
	if err != nil {
		println(err.Error())
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Printf("Waiting for connections...\n")
	for {
		conn, err := listener.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		ip := conn.RemoteAddr().(*net.TCPAddr).IP.String()
		if !isInWhitelist(ip) {
			conn.Close()
			continue
		}
		go handleClientConnection(conn)
	}
}

func handleClientConnection(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, ServerMaxRecvPacketSize)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	switch buf[0] {
	case HelloPacket:
		handleHelloPacket(conn, buf)
	case ConnPacket:
		handleConnPacket(conn, buf)
	}
}

func handleHelloPacket(conn net.Conn, buf []byte) {
	okay, port := parseHelloPacket(buf)
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
	defer listener.Close()
	fmt.Printf("New client connected, forwarding port %d.\n", port)
	for {
		userConn, err := listener.Accept()
		if err != nil {
			println(err.Error())
			continue
		}
		//fmt.Printf("[%d] Connection %p established.\n", port, userConn)
		token := common.GenerateToken()
		store.add(token, &userConn)
		err = SendConnPacket(conn, token)
		if err != nil {
			fmt.Printf("[%d] Failed to send connection packet: %s\n", port, err.Error())
			return
		}
	}
}

func handleConnPacket(conn net.Conn, buf []byte) {
	uuid := common.Bytes2Token(buf[1 : 1+tokenSize])
	userConn := store.get(uuid)
	store.remove(uuid)
	if userConn == nil {
		fmt.Println("Invalid UUID:", uuid)
		return
	}
	fmt.Printf("Connection established: %s <-> %s\n", conn.RemoteAddr().String(), (*userConn).RemoteAddr().String())
	go forward(*userConn, conn)
	forward(conn, *userConn)
}
