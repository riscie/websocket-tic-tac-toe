# websocket tic-tac-toe 
#### golang backend with gorilla websockets, vuejs frontend
![game sample](http://langhard.com/github/tic-tac-toe1.gif "game sample")
*(yeah, playing myself in this gif...)*

### get and run
* `go get https://github.com/riscie/websocket-tic-tac-toe`
* `go build`
* run the produced binary
* connect to http://localhost:8080

### gin (allows live-rebuilding the backend on changes)
* install gin: https://github.com/codegangsta/gin
* start with gin:  `gin --appPort 8080 r .\server.go`
