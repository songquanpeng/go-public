package common

import (
	"flag"
	"os"
)

var (
	PrintVersion = flag.Bool("version", false, "Print version and exit")
	PrintHelp    = flag.Bool("help", false, "Print help and exit")
	ConfigFile   = flag.String("config", "go-public.yaml", "Config file path")
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
}
