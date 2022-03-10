package server

import "gosrc.io/xmpp"
import "os"

import "fmt"

func Serverprint(){
    // config file for connection, yeah i know its part of client, this is just
    // the genesis commit
	config := xmpp.Config{
		TransportConfiguration: xmpp.TransportConfiguration{
			Address: "localhost:5222",
		},
		Jid:          "test@localhost",
		Credential:   xmpp.Password("test"),
		StreamLogger: os.Stdout,
		Insecure:     true,
		// TLSConfig: tls.Config{InsecureSkipVerify: true},
	}
    fmt.Print(config);
}



