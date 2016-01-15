# websocket tic-tac-toe
#### multiplayer tic-tac-toe using golang with gorilla websockets as backend and vuejs as frontend
![game sample](http://langhard.com/github/tic-tac-toe1.gif "game sample")

*(yeah, playing myself in this gif...)*

### get and run
* `go get https://github.com/riscie/websocket-tic-tac-toe`
* `go build`
* run the produced binary
* connect to http://localhost:8080

### gin (allows live-rebuilding on backend changes)
* install gin: https://github.com/codegangsta/gin
* start with gin:  `gin -a 8080 r .\server.go`


### ...this project is 'work in progress'
### Todos:
* finish implementing the game logic
* writing tests
