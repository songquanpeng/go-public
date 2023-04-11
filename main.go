package main

import (
	"go-public/common"
	"go-public/handler"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) == 3 && os.Args[1] == "init" {
		if os.Args[2] == "server" {
			common.InitConfigFile(true)
		} else if os.Args[2] == "client" {
			common.InitConfigFile(false)
		} else {
			println("Usage: go-public init <server|client>")
			os.Exit(1)
		}
		os.Exit(0)
	}
	if len(os.Args) == 1 {
		// Server mode
		common.LoadConfigFile(true)
		handler.ServeForever()
		os.Exit(0)
	}
	if len(os.Args) == 3 {
		// Client mode
		common.LoadConfigFile(false)
		localPort, err1 := strconv.Atoi(os.Args[1])
		remotePort, err2 := strconv.Atoi(os.Args[2])
		if err1 != nil || err2 != nil {
			println("Usage: go-public <local_port> <remote_port>")
			os.Exit(1)
		}
		handler.PublicPort(localPort, remotePort)
		os.Exit(0)
	}
	println("Usage: go-public <local_port> <remote_port> or go-public init <server|client>")
	os.Exit(1)
}
