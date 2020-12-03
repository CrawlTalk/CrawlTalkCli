package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"sync"
)

const (
	version = "0.01"
	//	defaultServer = "127.0.0.1"
	// defaultServer = "35.228.157.42"
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
	lastMessageTime  int
	flagNoColor      = flag.Bool("no-color", false, "Disable color output")
	mu               sync.Mutex
)

func main() {
	logfile, err := os.OpenFile("CrawlTalkCli.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	log.SetOutput(logfile)

	parseFlags()
	printBanner()
	requestServerUrl()
	fmt.Printf("Connecting to %s\n", color.BlueString(serverUrl.String()))
	if connectToServer() {
		fmt.Println(color.YellowString("Connected successfully!"))

		defer serverConnection.Close()

		for {
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
	}
}

func parseFlags() {
	flag.Parse()
	if *flagNoColor {
		color.NoColor = true // disables colorized output
	}
}
