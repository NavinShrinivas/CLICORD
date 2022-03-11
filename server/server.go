package main

import "github.com/gorilla/websocket"
import "fmt"
import "net/http"
import "flag"
import "log"

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

//In theory uprage is an HTTP header that asks the client to comply with
//HTTP standards, but it can ofc also be used to open connections.


func Echo_server(w http.ResponseWriter, r *http.Request){
    //a websocket connection is done using "conn" which is recvd from HTTP
    //so we are using port 80 for all our trasanctions
    //To make HTTP connections we need an "Upgrader"
    conn,err := upgrader.Upgrade(w,r,nil);
    //https://pkg.go.dev/github.com/gorilla/websocket#Upgrader.Upgrade
    //last nil for telling the header on the request packet is nil 

    if err != nil {
        log.Print("Error on Updrage header : ", err)
		return
	}
	defer conn.Close(); //closes if err was nil
    //defer delays excution till the near by return does its job.
    fmt.Print("Socket/TCP/HTTP connection successfull!");
    for {
        mt, message, err := conn.ReadMessage();
        if err != nil {
            log.Println("Error on reading : ", err);
            break
        }
        log.Printf("Recived : %s", message)
        err = conn.WriteMessage(mt, message)
        if err != nil {
            log.Println("Error on writing back:", err)
            break
        }
    }
}


var addr = flag.String("addr", "localhost:8080", "http service address")

func main(){

    //https://pkg.go.dev/net/http#ServeMux.HandleFunc
    //HandleFunc provides the write and reader for HTTP
    log.Print("Starting server...");
    http.HandleFunc("/ping" , Echo_server);
    log.Fatal(http.ListenAndServe(*addr, nil))
}
