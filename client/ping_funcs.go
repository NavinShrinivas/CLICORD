package main

import (
	"log"
	"net/url"
	"github.com/gorilla/websocket"
)


func Ping_server(){
    url_obj := url.URL{Scheme: "ws", Host: *addr, Path: "/ping"} ;
    log.Print("Connection to : ",url_obj,".... Trying to ping");
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
    conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""));
    //Above is to terminate connection grafully instead of just nuking the connection

}

