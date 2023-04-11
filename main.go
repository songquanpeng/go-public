package main

import (
	"go-public/client"
	"go-public/common"
	"go-public/server"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "init" {
		common.InitConfigFile()
		os.Exit(0)
	}
	if len(os.Args) == 1 {
		// Server mode
		server.ServeForever()
		os.Exit(0)
	}
	// Client mode
	common.LoadConfigFile()
	port, err := strconv.Atoi(os.Args[1])
	if err != nil {
		println("Usage: go-public <port>")
		os.Exit(1)
	}
	client.PublicPort(port)
}
