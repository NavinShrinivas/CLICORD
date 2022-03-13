package main 
import "net/http"
import "fmt"
import "github.com/gorilla/websocket"
import "log"

func Lobby_main(w http.ResponseWriter, r *http.Request){
    //first need to use upgrader to move to ws connection and also handle connections 
    //this also handle all new messages, cus all messages do come to /lobby1 api
    waitgroup.Add(1);
    //end point
    go init_funcs(w,r);
    go HandleMessages();
    waitgroup.Wait();
}


func add_conn(conn *websocket.Conn){
    active_conn[conn] += 1; 
    //each websocket connection is unique
}

func init_funcs(w http.ResponseWriter, r *http.Request){
    conn,err := upgrader.Upgrade(w,r,nil);

    if err!=nil{
        fmt.Print("Upgrade err :", err);
    }
    defer conn.Close();
    add_conn(conn);
    /*
     *for key,value := range active_conn{
     *    fmt.Print(key,value,"\n");
     *}
     */
    for {
        _, message, err := conn.ReadMessage();
        if err != nil {
            log.Println("Read error : ", err);
            break
        }
        realtime_msg <- string(message)
    }
}

func HandleMessages(){
    for{
        new_msg := <-realtime_msg 
        fmt.Print(new_msg,"\n");
        PushMessage(new_msg);
    }
}

func PushMessage(new_msg string){
    for ws := range active_conn{
        ws.WriteMessage(websocket.TextMessage, []byte(new_msg));
    }
}
