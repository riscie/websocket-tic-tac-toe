var tictactoe = new Vue({
    el: '#tictactoe',
    data: {
        gameState: {
            numPlayers: '1',
            someSlice: ["asdf", "bsdf"],
        }
    }
});


var socket = new WebSocket("ws://localhost:8080/ws");

socket.onmessage = function(evt){
    var newData = JSON.parse(evt.data);
    console.log(evt.data);
    //vue.data = evt.data;
    tictactoe.gameState = newData;
};

