package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/gorilla/websocket"
	"log"
	"strconv"
	"strings"
)

func connectToServer() bool {
	var err error
	log.Printf("Connecting to: %s\n", serverUrl.String())
	serverConnection, _, err = websocket.DefaultDialer.Dial(serverUrl.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
		return false
	}
	log.Printf("Connection established.\n")
	return true
}

func sendJson(body []byte) {
	log.Println("Sending:", string(body))
	err := serverConnection.WriteMessage(
		websocket.TextMessage,
		body)
	if err != nil {
		log.Println("write:", err)
	}
}

func authUser(login string, password string) (int64, string) {
	request := MoreliaT{
		Type: "auth",
		Data: DataT{
			User: []UserT{
				{
					Password: password,
					Login:    login,
				},
			},
		},
	}

	body, _ := json.Marshal(request)

	sendJson(body)

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}

	if response.Errors == nil {
		connectToServer()
		return authUser(login, password)
	}

	log.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
	PrintErrorCode("Auth user", response.Errors.Code, response.Errors.Status)
	if response.Errors.Code == 200 {
		return response.Data.User[0].UUID, response.Data.User[0].AuthId
	} else {
		return 0, ""
	}
}

func registerUser(login string, password string, username string, email string) (int64, string) {
	request := MoreliaT{
		Type: "register_user",
		Data: DataT{
			User: []UserT{
				{
					Password: password,
					Login:    login,
					Username: username,
					Email:    email,
				},
			},
		},
	}

	body, _ := json.Marshal(request)

	sendJson(body)

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}

	if response.Errors == nil {
		connectToServer()
		return registerUser(login, password, username, email)
	}

	log.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
	PrintErrorCode("Register user", response.Errors.Code, response.Errors.Status)
	if response.Errors.Code == 201 {
		return response.Data.User[0].UUID, response.Data.User[0].AuthId
	} else {
		return 0, ""
	}
}

func requestFlowList() {
	request := MoreliaT{
		Type: "all_flow",
		Data: DataT{
			User: []UserT{
				{
					UUID:   uuid,
					AuthId: authId,
				},
			},
		},
	}

	body, _ := json.Marshal(request)

	sendJson(body)

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}

	if response.Errors == nil {
		connectToServer()
		requestFlowList()
		return
	}

	log.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
	if response.Errors.Code == 200 {
		fmt.Println(color.CyanString("Flow list:"))
		for _, flow := range response.Data.Flow {
			fmt.Printf("ID: %s Title %s (%s)\n",
				color.CyanString(strconv.Itoa(flow.ID)),
				color.CyanString(flow.Title),
				color.CyanString(flow.Type))
		}
		return
	} else {
		return
	}

}

func requestMessagesList(flowId int, lastMessageTime int) int {
	request := MoreliaT{
		Type: "all_messages",
		Data: DataT{
			Time: lastMessageTime,
			User: []UserT{
				{
					UUID:   uuid,
					AuthId: authId,
				},
			},
			Flow: []FlowT{
				{
					ID: flowId,
				},
			},
		},
	}

	body, _ := json.Marshal(request)

	sendJson(body)

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}

	if response.Errors == nil {
		connectToServer()
		return requestMessagesList(flowId, lastMessageTime)
	}

	log.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
	if response.Errors.Code == 200 {
		for _, message := range response.Data.Message {
			if message.Time > lastMessageTime {
				lastMessageTime = message.Time
			}
			fmt.Printf("ID: %s From: %s Text: %s\n",
				color.CyanString(strconv.Itoa(message.ID)),
				color.CyanString(strconv.Itoa(message.FromUserUUID)),
				color.CyanString(message.Text))
		}
		return lastMessageTime + 1
	} else {
		return lastMessageTime
	}
}

func addFlow(flowName string, flowType string, userId int64) {
	request := MoreliaT{
		Type: "add_flow",
		Data: DataT{
			User: []UserT{
				{
					UUID:   uuid,
					AuthId: authId,
				},
			},
			Flow: []FlowT{
				{
					Type:  flowType,
					Title: flowName,
				},
			},
		},
	}

	if strings.ToLower(flowType) == "chat" {
		print("1")
		request.Data.User = append(request.Data.User, UserT{UUID: userId})
	}

	body, _ := json.Marshal(request)

	sendJson(body)

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}

	if response.Errors == nil {
		connectToServer()
		addFlow(flowName, flowType, userId)
		return
	}

	log.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
	PrintErrorCode("Add flow", response.Errors.Code, response.Errors.Status)
}

func sendMessage(flowId int, text string) {
	request := MoreliaT{
		Type: "send_message",
		Data: DataT{
			User: []UserT{
				{
					UUID:   uuid,
					AuthId: authId,
				},
			},
			Flow: []FlowT{
				{
					ID: flowId,
				},
			},
			Message: []MessageT{
				{
					Text: text,
				},
			},
		},
	}

	body, _ := json.Marshal(request)

	sendJson(body)

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}

	if response.Errors == nil {
		connectToServer()
		sendMessage(flowId, text)
		return
	}
	log.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
}
