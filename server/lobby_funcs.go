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

import "net/http"
import "fmt"
import "github.com/gorilla/websocket"
import "log"

func LobbyServer(w http.ResponseWriter, r *http.Request){
    //first need to use upgrader to move to ws connection and also handle connections 
    //this also handle all new messages, cus all messages do come to /lobby1 api
    wait_group.Add(1);
    //end point
    go InitFunc(w,r);
    go HandleMessages();
    wait_group.Wait();
}

//Add connection socket to a list
func AddConn(conn *websocket.Conn){
    active_conn[conn] += 1; 
    //each websocket connection is unique
}

//Init connection handler
func InitFunc(w http.ResponseWriter, r *http.Request){
    conn,err := upgrader.Upgrade(w,r,nil);

    if err!=nil{
        fmt.Print("Upgrade err :", err);
    }
    defer conn.Close();
    AddConn(conn);
    /*
     *for key,value := range active_conn{
     *    fmt.Print(key,value,"\n");
     *}
     */
    for {
        _, message, err := conn.ReadMessage();
        if err != nil {
            log.Println("Read error : ", err);
            break
        }
        realtime_msg <- string(message)
    }
}

func HandleMessages(){
    for{
        new_msg := <-realtime_msg 
        fmt.Print(new_msg,"\n");
        PushMessages(new_msg);
    }
}

func PushMessages(new_msg string){
    for ws := range active_conn{
        err := ws.WriteMessage(websocket.TextMessage, []byte(new_msg));
        if err != nil {
            fmt.Println("Internal write error : ", err);
        }
    }
}
