package main

import (
	"go-public/client"
	"go-public/common"
	"go-public/server"
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
		server.ServeForever()
		os.Exit(0)
	}
	if len(os.Args) == 2 {
		// Client mode
		common.LoadConfigFile(false)
		port, err := strconv.Atoi(os.Args[1])
		if err != nil {
			println("Usage: go-public <port>")
			os.Exit(1)
		}
		client.PublicPort(port)
	}
	println("Usage: go-public <port> or go-public init <server|client>")
	os.Exit(1)
}
