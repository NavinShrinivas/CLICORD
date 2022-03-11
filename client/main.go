package main

import (
	"log"
	"net/url"
    "flag"
	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")
func main(){
    // about URL schemes : https://en.wikipedia.org/wiki/List_of_URI_scheme
    url_obj := url.URL{Scheme: "ws", Host: *addr, Path: "/ping"} ;
    //ws is for websockets scheme, this allows the protocol to be upgraded
    log.Print("Connection to : ",url_obj,"...");
    conn, _, err := websocket.DefaultDialer.Dial(url_obj.String(), nil);
    if err!=nil{
        log.Print("Error from trying to connnect : ",err);
    }
    e := conn.WriteMessage(websocket.TextMessage,[]byte("Ping"));
    if e!=nil{
        log.Print("Write Err : ",err);
    }
    _, message, err := conn.ReadMessage()
    if err != nil {
        log.Println("read:", err)
        return
    }
    log.Printf("got back : %s", message)
}
