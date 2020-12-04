package main

import (
	"fmt"
	"net/url"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func printBanner() {
	fmt.Println(color.YellowString("CrawlTalk Client Version: %s", version))
}

func requestServerUrl() {
	var server string
	var port string
	var protocol string
	if *flagServer == "" {
		fmt.Printf("Enter server address %s: ", color.GreenString("[%s]", defaultServer))
		fmt.Scanln(&server)
		if server == "" {
			server = defaultServer
		}
	} else {
		server = *flagServer
	}
	if *flagPort == 0 {
		fmt.Printf("Enter server port %s: ", color.GreenString("[%s]", defaultPort))
		fmt.Scanln(&port)
		if port == "" {
			port = defaultPort
		}
	} else {
		port = strconv.Itoa(*flagPort)
	}
	if *flagSchema != "ws" && *flagServer != "wss" {
		fmt.Printf("Enter server scheme  (ws, wss) %s: ", color.GreenString("[%s]", defaultScheme))
		fmt.Scanln(&protocol)
		if protocol == "" {
			protocol = defaultScheme
		}
		if protocol != "ws" && protocol != "wss" {
			fmt.Println(color.RedString("Bad scheme, default using: %s", defaultScheme))
			protocol = defaultScheme
		}
	} else {
		protocol = *flagSchema
	}
	serverUrl = url.URL{Scheme: "ws", Host: server + ":" + port, Path: defaultPath}
}

func requestLoginAndConnect() {
	var register string
	var login string
	var password string
	var username string
	var email string

	fmt.Printf("Do you want to register new account (Y/N) %s: ", color.GreenString("[%s]", defaultRegister))
	fmt.Scanln(&register)
	register = strings.ToUpper(register)
	if register != "Y" {
		register = defaultRegister
	}
	fmt.Printf("Enter login name %s: ", color.GreenString("[%s]", defaultLogin))
	fmt.Scanln(&login)
	if login == "" {
		login = defaultLogin
	}
	fmt.Printf("Enter password %s: ", color.GreenString("[%s]", defaultPassword))
	fmt.Scanln(&password)
	if password == "" {
		password = defaultPassword
	}
	if register == "N" {
		uuid, authId = authUser(login, password)
	} else {
		fmt.Printf("Enter username %s: ", color.GreenString("[%s]", defaultUsername))
		fmt.Scanln(&username)
		if username == "" {
			username = defaultUsername
		}
		fmt.Printf("Enter email %s: ", color.GreenString("[%s]", defaultEmail))
		fmt.Scanln(&email)
		if email == "" {
			email = defaultEmail
		}
		uuid, authId = registerUser(login, password, username, email)
	}
}

func awaitUserCommandOrExit(mode string) bool {
	var command string
	if mode == "flow_list" {
		fmt.Printf("%s", color.HiGreenString("Enter flow ID or system command: "))
		fmt.Scanln(&command)
		if command == "" {
			return false
		} else {
			lowCaseCommand := strings.ToLower(command)
			if strings.HasPrefix(lowCaseCommand, "/exit") {
				return true
			}
			if strings.HasPrefix(lowCaseCommand, "/create") {
				return true
			}
			if strings.HasPrefix(lowCaseCommand, "/help") {
				PrintHelpInteractive()
				return false
			}
			if flowId, err := strconv.Atoi(lowCaseCommand); err != nil {
				fmt.Println(color.RedString("Unknown command"))
			} else {
				EnterToFlow(flowId)
			}
			return false
		}
	}
	if mode == "message_list" {
		fmt.Printf("%s", color.HiGreenString("Type message text or system command: "))
		fmt.Scanln(&command)
		if command == "" {
			return false
		} else {
			lowCaseCommand := strings.ToLower(command)
			if strings.HasPrefix(lowCaseCommand, "/exit") {
				return true
			}
			if strings.HasPrefix(lowCaseCommand, "/inception") {
				lastMessageTime = 1
				return false
			}
			if strings.HasPrefix(lowCaseCommand, "/help") {
				PrintHelpInteractive()
				return false
			}
			return false
		}
	}
	return false
}

func EnterToFlow(flowId int) {
	lastMessageTime = 1
	fmt.Println(color.CyanString("Opened flow ID %d", flowId))
	for {
		lastMessageTime = requestMessagesList(flowId, lastMessageTime)
		if awaitUserCommandOrExit("message_list") {
			break
		}
	}
}

func PrintHelpInteractive() {
	fmt.Println(color.HiBlueString("Available commands in interactive mode:"))
	fmt.Println(color.HiBlueString("/help - this help page"))
	fmt.Println()
	fmt.Println(color.HiBlueString("Flow list mode:"))
	fmt.Println(color.HiBlueString("/exit - exit from program to command prompt"))
	fmt.Println()
	fmt.Println(color.HiBlueString("In flow mode:"))
	fmt.Println(color.HiBlueString("/exit - exit to flow list"))
	fmt.Println(color.HiBlueString("/inception - show all flow messages from beginning"))
}
