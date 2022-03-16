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
	"os"
	"sync"
)

//----------globals----------

// This address is the constatnt address for the server, and use in every function
var addr = flag.String("addr", "catss.me:8080", "http service address")

//messgae structure, needs a lot more config to work.
type MsgStruct struct {
	Msg_content string `json : "Msg_content"`
	Username    string `json : "Username"`
}

//json library expects exported members of struct, hence capital starting

var wait_group sync.WaitGroup //for async functions, needs to be global

//TUI and UI Sutff :

//message managment , will also help with database operations
//impl for these struct are not bound the the struct, this can and has to be done
//evetually
type Node struct {
	Key  string
	next *Node
}
type List struct {
	head *Node
}

var msg_list = &List{
	head: nil,
}

var username_glob string

var glob_emoji_data map[string]string

//----------End of Globals----------

func main() {
	args := os.Args
	if args[1] == "ping" {
		PingServer()
	}
	if args[1] == "lobby" {
		emoji_data, _ := EmojiMapGenerator()
		glob_emoji_data = emoji_data
		var lobby_id = args[2]
		conn := LobbyMain(lobby_id) //connection established in this function
		wait_group.Add(1)
		go SendMessages(conn)
		go RecieveMessages(conn)
		wait_group.Wait() //concept very near to semaphores from os
		//https://stackoverflow.com/questions/24425987/why-is-my-goroutine-not-executedS
		//explains go routine really quickly
	}
}
