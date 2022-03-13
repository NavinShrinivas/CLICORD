package main

import "github.com/gorilla/websocket"
import "net/http"
import "flag"
import "log"
import "sync"

//----------globals----------

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
//In theory uprage is an HTTP header that asks the client to comply with
//HTTP standards, but it can ofc also be used to open connections.

var addr = flag.String("addr", "127.0.0.1:80", "http service address")

//for port fwding : https://developer.mozilla.org/en-US/docs/Web/API/WebSockets_API/Writing_WebSocket_servers

//array of active websocket connections 

//Before : for no public rollout 2048 connections should be plenty
var active_conn = make(map[*websocket.Conn]int);
//After : Keeping the the ds as array makes it linear time for adding and deleting 
//connections, which imo happens all the time 
//hence a HashMap can do this in constant time, and also removes all restrictions 
//on number of maintained connections

//Concurrency channels, damn threading is so ez in go!! 
var realtime_msg = make(chan string);
//https://medium.com/@thejasbabu/concurrency-in-go-e4a61ec96491

var waitgroup sync.WaitGroup;

//----------End of globals----------


func main(){
    //https://pkg.go.dev/net/http#ServeMux.HandleFunc
    //HandleFunc provides the write and reader for HTTP
    log.Print("Starting server...");
    http.HandleFunc("/ping" , Ping_server);
    http.HandleFunc("/lobby1" , Lobby_main);
    log.Fatal(http.ListenAndServe(*addr, nil))
    
}
