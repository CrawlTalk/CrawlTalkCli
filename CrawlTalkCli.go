package main

import (
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

const (
	version = "0.1.2"
	//	defaultServer = "127.0.0.1"
	defaultServer = "35.228.157.42"
	// defaultServer = "192.168.60.173"
	defaultPort   = "8000"
	defaultScheme = "ws"
	defaultPath   = "/ws"

	defaultRegister = "N"
	defaultLogin    = "admin"
	defaultPassword = "admin"
	defaultUsername = "Administrator"
	defaultEmail    = "admin@morelia.com"

	defaultFlowType = "group"

	logFileName = "CrawlTalkCli.log"
)

var (
	serverUrl        url.URL
	uuid             int64 = 0
	authId                 = ""
	serverConnection *websocket.Conn
	lastMessageTime  int
	flagNoColor      = flag.Bool("no-color", false, "Disable color output")
	flagServer       = flag.String("server", "", "Default server. If specified client will not request it interactive.")
	flagPort         = flag.Int("port", 0, "Default port. If specified client will not request it interactive.")
	flagSchema       = flag.String("schema", "", "Default schema (ws or wss). If specified client will not request it interactive.")
	flagRegister     = flag.Bool("register", false, "Register new account")
	flagSignIn       = flag.Bool("sign-in", false, "Sign in to existed account")
	flagLogin        = flag.String("login", "", "Default login. If specified client will not request it interactive.")
	flagPassword     = flag.String("password", "", "Default password. If specified client will not request it interactive.")
	flagUserName     = flag.String("username", "", "Default username. For registration only. If specified client will not request it interactive.")
	flagEmail        = flag.String("email", "", "Default email. For registration only. If specified client will not request it interactive.")
)

func main() {
	logfile, err := os.OpenFile(logFileName, os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0644)
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
					if awaitUserCommandOrExit("flow_list", 0) {
						return
					}
				}
			} else {
				if *flagRegister && *flagLogin != "" {
					return
				}
				if *flagSignIn && *flagLogin != "" && *flagPassword != "" {
					return
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
