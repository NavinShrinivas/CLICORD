package main

import "github.com/gorilla/websocket"
import "fmt"
import "net/http"
import "log"

func Ping_server(w http.ResponseWriter, r *http.Request){
    //a websocket connection is done using "conn" which is recvd from HTTP
    //so we are using port 8080 for all our init connection
    //To make ws connections we need an "Upgrader"
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
        _, message, err := conn.ReadMessage();
        if err != nil {
            log.Println("Read error : ", err);
            break
        }
        log.Printf("Recived : %s", message)
        err = conn.WriteMessage(websocket.TextMessage,[]byte("Pong"))
        if err != nil {
            log.Println("Error on writing back:", err)
            break
        }
    }
}


