package main

import "github.com/gorilla/websocket"
import "net/http"
import "flag"
import "log"

//----------globals----------

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}
//In theory uprage is an HTTP header that asks the client to comply with
//HTTP standards, but it can ofc also be used to open connections.

var addr = flag.String("addr", "localhost:8080", "http service address")

//----------End of globals----------


func main(){

    //https://pkg.go.dev/net/http#ServeMux.HandleFunc
    //HandleFunc provides the write and reader for HTTP
    log.Print("Starting server...");
    http.HandleFunc("/ping" , Ping_server);
    log.Fatal(http.ListenAndServe(*addr, nil))
}
