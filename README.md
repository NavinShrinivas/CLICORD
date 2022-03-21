# CLICORD 

CLICORD is a cli based chat system written from scratch using websockets in go.

## Features

- File tansfer [Any file as long as you give a relative path]
- Infitine Lobbies with Infinite users, only your backend resources are the constraint 
- Emojis :)) 😄
- Self hostable server [If yo want that privacy, but ofc network setup is left to you]
- Multiplatform Client [Tested on Linux and Windows]


## To be implemented 

- User Authentication
- E2E encryption
- File size restriction
- Database to avoid risky socket to socket data tansmission

## Bugs being worked on 

- CLI sucks, transistion to TUI is in the works
- Pushing file uplods to seperate thread
 
## Usage

Needs GO toolchain version over 1.6.

### Ping pong server 

```sh
   in /client : $ go run *.go ping 
   in /server : $ go run *.go
```
### Creating a particular server/lobby for clients

```sh
   in /client: $ go run *.go lobby server_name
   in /server: $ go run *.go
```
