package main

import (
	"log"
	"net/url"
    "flag"
	"github.com/gorilla/websocket"
    "os"
)
//globals 
//As NSK would say explain why global? 
// This address is the constatnt address for the server, and use in every function

var addr = flag.String("addr", "localhost:8080", "http service address")

//----------End of Globals----------

func ping_server(){
    // about URL schemes : https://en.wikipedia.org/wiki/List_of_URI_scheme
    url_obj := url.URL{Scheme: "ws", Host: *addr, Path: "/ping"} ;
    //ws is for websockets scheme, this allows the protocol to be upgraded
    log.Print("Connection to : ",url_obj,".... Trying to ping");
    conn, _, err := websocket.DefaultDialer.Dial(url_obj.String(), nil);
    //default dail up to connect to an ws server
    
    //Below all are self explainatory
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

func main(){
    args := os.Args
    if args[1] == "ping"{
        ping_server();
    }
}
