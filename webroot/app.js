var tictactoe = new Vue({
    el: '#tictactoe',
    data: {
        gameState: {
        }
    },
    methods: {
        chooseField: function(fieldNum){
            socket.send(fieldNum);
        },
    }
});


var socket = new WebSocket("ws://localhost:8080/ws");
//socket.send('I say hello to you, backend...');

socket.onmessage = function(evt){
    var newData = JSON.parse(evt.data);
    console.log(evt.data);
    //vue.data = evt.data;
    tictactoe.gameState = newData;
};

