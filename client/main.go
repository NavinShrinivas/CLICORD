package main

import (
    "flag"
    "os"
    "sync"
)



//----------globals----------
//As NSK would say explain why global? 
// This address is the constatnt address for the server, and use in every function

var addr = flag.String("addr", "localhost:8080", "http service address")

type msg_struct struct{
    msg_content string
    username string
}

var waitgroup sync.WaitGroup; //for async functions, needs to be global

//TUI Sutff : 
type Node struct{
    key string 
    next *Node 
}
type List struct{
    head *Node
}

var msg_list = &List{
    head : nil,
}

//----------End of Globals----------

func main(){
    args := os.Args
    if args[1] == "ping"{
        Ping_server();
    }
    if args[1] == "lobby1"{
        //var lobby = args[2]
        conn:= LobbyMain();
        waitgroup.Add(1);
        go SendMessages(conn);
        go RecieveMessages(conn);
        waitgroup.Wait();
        //https://stackoverflow.com/questions/24425987/why-is-my-goroutine-not-executedS
        //explains go routine really quickly
    }
}
