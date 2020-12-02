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
	fmt.Printf("Enter server address %s: ", color.GreenString("[%s]", defaultServer))
	fmt.Scanln(&server)
	if server == "" {
		server = defaultServer
	}
	fmt.Printf("Enter server port %s: ", color.GreenString("[%s]", defaultPort))
	fmt.Scanln(&port)
	if port == "" {
		port = defaultPort
	}
	fmt.Printf("Enter server scheme %s: ", color.GreenString("[%s]", defaultScheme))
	fmt.Scanln(&protocol)
	if protocol == "" {
		protocol = defaultScheme
	}
	if protocol != "ws" && protocol != "wss" {
		fmt.Println(color.RedString("Bad scheme, default using: %s", defaultScheme))
		protocol = defaultScheme
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
			return false
		}
	}
	return false
}

func EnterToFlow(flowId int) {
	var lastMessageTime = 1
	fmt.Println(color.CyanString("Opened flow ID %d", flowId))
	for {
		lastMessageTime = requestMessagesList(flowId, lastMessageTime) + 1
		if awaitUserCommandOrExit("message_list") {
			break
		}
	}
}
