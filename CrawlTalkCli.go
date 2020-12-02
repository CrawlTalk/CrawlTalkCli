package main

import (
	"flag"
	"log"
	"net/url"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

const (
	version = "0.01"
	//	defaultServer = "127.0.0.1"
	//	defaultServer = "35.228.157.42"
	defaultServer = "192.168.60.173"
	defaultPort   = "8000"
	defaultScheme = "ws"
	defaultPath   = "/ws"

	defaultRegister = "N"
	defaultLogin    = "admin"
	defaultPassword = "admin"
	defaultUsername = "Administrator"
	defaultEmail    = "admin@morelia.com"
)

var (
	serverUrl        url.URL
	uuid             int64 = 0
	authId                 = ""
	serverConnection *websocket.Conn
	flagNoColor      = flag.Bool("no-color", false, "Disable color output")
)

func main() {
	log.SetFlags(0)
	parseFlags()
	printBanner()
	requestServerUrl()
	connectToServer()
	defer serverConnection.Close()
	requestLoginAndConnect()
	if authId != "" {
		for {
			requestFlowList()
			if awaitUserCommandOrExit("flow_list") {
				break
			}
		}
	}
}

func parseFlags() {
	flag.Parse()
	if *flagNoColor {
		color.NoColor = true // disables colorized output
	}
}
