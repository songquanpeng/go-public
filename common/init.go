package common

import (
	"flag"
	"os"
)

var (
	PrintVersion = flag.Bool("version", false, "Print version and exit")
	PrintHelp    = flag.Bool("help", false, "Print help and exit")
	ConfigPath   = flag.String("config", "", "Config file path")
)

func PrintUsage() {
	println("Go Public " + Version + " - A simple port forwarding tool.")
	println("Copyright (C) 2023 JustSong. All rights reserved.")
	println("GitHub repository: https://github.com/songquanpeng/go-public")
	println("Usage: go-public [--config <config file path>] [--version] [--help]")
	println("       go-public init <server|client>")
	println("       go-public <local_port> <remote_port>")
}

func init() {
	flag.Parse()
	if *PrintVersion {
		println(Version)
		os.Exit(0)
	}
	if *PrintHelp {
		PrintUsage()
		os.Exit(0)
	}
}
