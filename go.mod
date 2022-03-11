module CLICORD

go 1.17

replace CLICORD/server => ./server

require CLICORD/server v0.0.0-00010101000000-000000000000

require github.com/gorilla/websocket v1.5.0 // indirect
