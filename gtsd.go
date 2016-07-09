package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gts/gts_config"
	"os"
)

const VERSION = "0.0.1"

func main() {
	var configFile string
	var debugMode, daemonMode, checkConf, showHelp, showVersion bool
	flag.StringVar(&configFile, "c", "/etc/gts/gts.conf", "Configuration filename.")
	flag.BoolVar(&debugMode, "D", false, "Debug mode.")
	flag.BoolVar(&daemonMode, "d", false, "Daemon mode.")
	flag.BoolVar(&checkConf, "C", false, "Check Configuration only.")
	flag.BoolVar(&showHelp, "h", false, "Show help.")
	flag.BoolVar(&showVersion, "v", false, "Show version.")
	flag.Parse()

	os.Stderr.Write([]byte(fmt.Sprintf("GTS %v - Another Traffic Server.\n", VERSION)))
	os.Stderr.Write([]byte("Copyright (C) 2016 Mingcai SHEN.\n"))
	if showVersion {
		os.Exit(0)
	}

	if showHelp {
		flag.Usage()
		os.Exit(0)
	}

	g_config, err := gts_config.LoadConfig(configFile, nil)
	if nil != err {
		os.Stderr.Write([]byte(fmt.Sprintf("Configuration(%s) Invalid <%s> ! \n", configFile, err)))
		flag.Usage()
		os.Exit(1)
	}

	if checkConf {

		fmt.Printf("Configuration(%s) valid!\n", configFile)
		fmt.Println("Configurations: ")
		out, err := json.MarshalIndent(g_config, "", " ")
		if nil != err {
			fmt.Errorf("Failed!\n")
		} else {
			fmt.Println(string(out))

		}
		os.Exit(0)
	}

}
