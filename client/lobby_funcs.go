package main

import (
	"log"
	"net/url"
	"github.com/gorilla/websocket"
    "fmt"
    "bufio"
    "os"
)

func LobbyMain() *websocket.Conn{
    url_obj := url.URL{Scheme: "ws", Host: *addr, Path: "/lobby1"} ;
    log.Print("Connection to : ",url_obj,".... Trying to connect to lobby1 \n");
    conn, _, err := websocket.DefaultDialer.Dial(url_obj.String(), nil);
    
    if err!=nil{
        log.Print("Error from trying to connnect : ",err);
    }
    return conn
}

func SendMessages(conn *websocket.Conn){
    // this has to be rewritten as all this is front and ive written basiccc 
    //just to test
    for{
        fmt.Print("Enter message to send : ")
        var msg string
        scanner := bufio.NewScanner(os.Stdin)
        scanner.Scan() 
        msg = scanner.Text()
        //like C, go has problems with space
        conn.WriteMessage(websocket.TextMessage , []byte(msg));
    }
}

func RecieveMessages(conn *websocket.Conn){
    for{
        _,message,err := conn.ReadMessage();
        if err != nil{
            log.Print("Read error : ",err);
            waitgroup.Done();
            break;
        }
        fmt.Print("\033c\n")
        AppendNode(string(message),msg_list);
        PrintNode(msg_list);
        fmt.Print("Enter message to send : "); //this may be confusig, but this is only for 
        //clean ui, we arent doing any sending messages here.
    }
}



//If we have better TUI we can scrape these parts off : 

func AppendNode(msg string, list *List){
    var copy_list = list;
    if copy_list.head == nil{
        copy_list.head = &Node{
            key: msg,
            next: nil,
        }
        return
    }
    copy_nodes := copy_list.head;
    for copy_nodes.next != nil{
        copy_nodes = copy_nodes.next;
    }
    copy_nodes.next = &Node{
        key : msg,
        next: nil,
    }
}

func PrintNode(list *List){
    copy_nodes := list.head;
    for copy_nodes != nil{
        fmt.Println(copy_nodes.key);
        copy_nodes = copy_nodes.next;
    }
}

