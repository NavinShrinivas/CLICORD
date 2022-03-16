/**********************************************************************************
  ____ _     ___ ____ ___  ____  ____
 / ___| |   |_ _/ ___/ _ \|  _ \|  _ \
| |   | |    | | |  | | | | |_) | | | |
| |___| |___ | | |__| |_| |  _ <| |_| |
 \____|_____|___\____\___/|_| \_\____/

Copyright (c) 2021  CLICORD

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
**********************************************************************************/

package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
)

func LobbyMain(lobby_id string) *websocket.Conn {
	//Build URL for connection endpoint  : /lobby
	var endpoint = "/lobby/" + lobby_id

	fmt.Print("Enter user name : ")
	fmt.Scanln(&username_glob)

	url_obj := url.URL{Scheme: "ws", Host: *addr, Path: endpoint}
	log.Print("Connection to : ", url_obj, ".... Trying to connect to ", lobby_id, " \n")
	conn, _, err := websocket.DefaultDialer.Dial(url_obj.String(), nil)
	fmt.Print("Connection Succesfull. Press ENTER to continue")
	fmt.Scanln()
	//fmt.Print("\033c\n")
	if err != nil {
		log.Print("Connection Error : ", err)
	}
	return conn
}

func SendMessages(conn *websocket.Conn) {
	//Basic UI, just printing. Will need improvements very soon.
	for {
		fmt.Print("Enter message to send : ")
		var msg string
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg = scanner.Text()
		msg_built := &MsgStruct{
			Msg_content: Fancyfier(msg),
			Username:    username_glob,
		}
		e, _ := json.Marshal(msg_built)

		//like C, go has problems with spaces in input as well
		err := conn.WriteMessage(websocket.TextMessage, []byte(e))
		if err != nil {
			fmt.Println("Something went wrong on our end. May be you are giving really weird chars?")
			wait_group.Done()
		}
	}
}

func RecieveMessages(conn *websocket.Conn) {
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Print("Read error : ", err)
			wait_group.Done()
			break
		}
		fmt.Print("\033c\n")
		AppendNode(string(message), msg_list)
		PrintNode(msg_list)
		fmt.Print("Enter message to send : ") //this may be confusig, but this is only for
		//clean ui, we arent doing any sending messages here.
	}
}

//If we have better DB we maybe able to scrape these funs off:

func AppendNode(msg string, list *List) {

	//Can be rewritten so much better, But it does get the job done
	var copy_list = list
	if copy_list.head == nil {
		copy_list.head = &Node{
			Key:  msg,
			next: nil,
		}
		return
	}
	copy_nodes := copy_list.head
	for copy_nodes.next != nil {
		copy_nodes = copy_nodes.next
	}
	copy_nodes.next = &Node{
		Key:  msg,
		next: nil,
	}
}

func PrintNode(list *List) {
	copy_nodes := list.head
	for copy_nodes != nil {
		var msgstruct MsgStruct
		var json_str = copy_nodes.Key
		marshal_err := json.Unmarshal([]byte(json_str), &msgstruct)
		if marshal_err != nil {
			log.Println("Something went wrong, Unmarshal err.")
		}

		fmt.Println(msgstruct.Username, ":", msgstruct.Msg_content)
		copy_nodes = copy_nodes.next
	}
}
