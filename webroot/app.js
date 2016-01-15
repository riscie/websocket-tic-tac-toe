Vue.config.debug = true; //TODO: Remove in production

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

socket.onmessage = function(evt){
    var newData = JSON.parse(evt.data);
    console.log(evt.data); //TODO: Remove in production
    tictactoe.gameState = newData;
};

