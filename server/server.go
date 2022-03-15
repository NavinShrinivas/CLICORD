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
//Coding Style :
/*
*All functions follow PascalCase : Why? GOLANG complies with this
*All variables follow snake_case : Cus these kind of naming makes sense
*All types follow PascalCase too : Look more uniform
*To indent if on vim do : gg=G, if not leave it to the ones with vim
*
 */
package main

import (
	"flag"
	"github.com/gorilla/mux"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"sync"
)

//----------globals----------

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

//In theory uprage is an HTTP header that asks the client to comply with
//HTTP standards, but it can ofc also be used to open connections.

var addr = flag.String("addr", "0.0.0.0:80", "http service address")

//for port fwding : https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers

var wait_group sync.WaitGroup

type LobbyData struct {
	active_conn  map[*websocket.Conn]int
	realtime_msg chan string
}

//Keeping the the ds as array makes it linear time for adding and deleting
//connections, which imo happens all the time
//hence a HashMap can do this in constant time, and also removes all restrictions
//on number of maintained connections
//Concurrency channels
//to be an atomic process
//https://medium.com/@thejasbabu/concurrency-in-go-e4a61ec96491

var lobby_data = make(map[string]*LobbyData) //for n number of lobbies

//----------End of globals----------

func main() {
	//https://pkg.go.dev/net/http#ServeMux.HandleFunc
	//HandleFunc provides the write and reader for HTTP
	log.Print("Starting server...")
	r := mux.NewRouter()
	r.HandleFunc("/ping", PingServer)
	r.HandleFunc("/lobby/{lobby_id}", LobbyServer)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe(*addr, nil))

}
