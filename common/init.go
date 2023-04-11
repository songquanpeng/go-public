package common

import (
	"flag"
	"os"
)

var (
	PrintVersion = flag.Bool("version", false, "Print version and exit")
	PrintHelp    = flag.Bool("help", false, "Print help and exit")
	ConfigFile   = flag.String("config", "go-template.yaml", "Config file path")
)

func init() {
	flag.Parse()
	if *PrintVersion {
		println(Version)
		os.Exit(0)
	}
	if *PrintHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}
	if len(os.Args) > 1 && os.Args[1] == "init" {
		initConfigFile()
		os.Exit(0)
	}
	loadConfigFile()
}
