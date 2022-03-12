package main

import (
    "flag"
    "os"
)



//----------globals----------
//As NSK would say explain why global? 
// This address is the constatnt address for the server, and use in every function

var addr = flag.String("addr", "localhost:8080", "http service address")

type msg_struct struct{
    msg_content string
    username string
}


//----------End of Globals----------




func main(){
    args := os.Args
    if args[1] == "ping"{
        Helloworld();
        Ping_server();
    }
    if args[1] == "connect_lobby"{
        //var lobby = args[2] 

    }
}
