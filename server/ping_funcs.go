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

import "github.com/gorilla/websocket"
import "fmt"
import "net/http"
import "log"

func PingServer(w http.ResponseWriter, r *http.Request) {
	//a websocket connection is done using "conn" which is recvd from HTTP
	//so we are using port 8080 for all our init connection
	//To make ws connections we need an "Upgrader"
	conn, err := upgrader.Upgrade(w, r, nil)
	//https://pkg.go.dev/github.com/gorilla/websocket#Upgrader.Upgrade
	//last nil for telling the header on the request packet is nil

	if err != nil {
		log.Print("Error on Updrage header : ", err)
		return
	}
	defer conn.Close() //closes if err was nil
	//defer delays excution till the near by return does its job.
	fmt.Print("Socket/TCP/HTTP connection successfull!")
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			log.Println("Read error : ", err)
			break
		}
		log.Printf("Recived : %s", message)
		err = conn.WriteMessage(websocket.TextMessage, []byte("Pong"))
		if err != nil {
			log.Println("Error on writing back:", err)
			break
		}
	}
}
