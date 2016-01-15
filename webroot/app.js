Vue.config.debug = true;

var tictactoe = new Vue({
    el: '#tictactoe',
    data: {
        gameState: {
            started: false,
            fields: [],
        },
    },
    computed: {
        row1: function() {
            return this.gameState.fields.slice(0,3);
        },
        row2: function() {
            return this.gameState.fields.slice(3,6);
        },
        row3: function() {
            return this.gameState.fields.slice(6,9);
        },
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

