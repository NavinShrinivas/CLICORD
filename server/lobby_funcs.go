package main 
import "github.com/gorilla/websocket"
import "net/http"
import "fmt"

func Lobby_main(w http.ResponseWriter, r *http.Request){
    //first need to use upgrader to move to ws connection and also handle connections 
    init_funcs(w,r);
}


func init_funcs(w http.ResponseWriter, r *http.Request){
    conn,err := upgrader.Upgrade(w,r,nil);

    if err!=nil{
        fmt.Print("Upgrade err : %s", err);
    }
    defer conn.Close();
    append_conn(conn);
}

