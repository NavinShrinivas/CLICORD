/**********************************************************************************
  ____ _     ___ ____ ___  ____  ____
 / ___| |   |_ _/ ___/ _ \|  _ \|  _ \
| |   | |    | | |  | | | | |_) | | | |
| |___| |___ | | |__| |_| |  _ <| |_| |
 \____|_____|___\____\___/|_| \_\____/

Copyright (c) 2021 CLICORD

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
	"github.com/gorilla/websocket"
	"log"
	"net/url"
)

//Unrefactored, feel free to do it complying to instructions
// in main.go
func PingServer() {
	url_obj := url.URL{Scheme: "ws", Host: *addr, Path: "/ping"}
	log.Print("Connection to : ", url_obj, ".... Trying to ping")
	conn, _, err := websocket.DefaultDialer.Dial(url_obj.String(), nil)

	if err != nil {
		log.Print("Error from trying to connnect : ", err)
	}
	e := conn.WriteMessage(websocket.TextMessage, []byte("Ping"))
	if e != nil {
		log.Print("Write Err : ", err)
	}
	_, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read:", err)
		return
	}
	log.Printf("got back : %s", message)
	conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	//Above is to terminate connection grafully instead of just nuking the connection

}
