module CLICORD

go 1.17

replace CLICORD/server => ./server

require CLICORD/server v0.0.0-00010101000000-000000000000

require (
	github.com/google/uuid v1.1.1 // indirect
	golang.org/x/xerrors v0.0.0-20190717185122-a985d3407aa7 // indirect
	gosrc.io/xmpp v0.5.1 // indirect
	nhooyr.io/websocket v1.6.5 // indirect
)
