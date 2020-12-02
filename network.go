package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	"github.com/fatih/color"
	"github.com/gorilla/websocket"
)

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

	log.Println("Sending:", string(body))
	err := serverConnection.WriteMessage(
		websocket.TextMessage,
		[]byte(string(body)),
	)
	if err != nil {
		log.Println("write:", err)
		return 0, ""
	}

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}
	fmt.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
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

	log.Println("Sending:", string(body))
	err := serverConnection.WriteMessage(
		websocket.TextMessage,
		[]byte(string(body)),
	)
	if err != nil {
		log.Println("write:", err)
		return 0, ""
	}

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}
	fmt.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
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

	log.Println("Sending:", string(body))
	err := serverConnection.WriteMessage(
		websocket.TextMessage,
		[]byte(string(body)),
	)
	if err != nil {
		log.Println("write:", err)
		return
	}

	var response MoreliaT
	if err := serverConnection.ReadJSON(&response); err != nil {
		log.Println(err)
	}
	fmt.Println(response.Errors.Code, response.Errors.Status, response.Errors.Detail)
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

func connectToServer() {
	fmt.Printf("Connecting to %s\n", color.BlueString(serverUrl.String()))

	var err error

	serverConnection, _, err = websocket.DefaultDialer.Dial(serverUrl.String(), nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	//defer serverConnection.Close()
	fmt.Println(color.YellowString("Connected successfully!"))
}
