// creating the websocket connection. Test for dev environment
var socket;

if (window.location.host === "127.0.0.1:8080") {
  socket = new WebSocket("ws://127.0.0.1:8080/ws");
} else {
  socket = new WebSocket("ws://tic-tac-toe.langhard.com/ws");
}
// when an update is received via ws connection, we update the model
socket.onmessage = function(evt){
    var newData = JSON.parse(evt.data);
    console.log(evt.data); //TODO: Remove in production
    tictactoe.gameState = newData;
};


// vuejs debug mode
Vue.config.debug = true; //TODO: Remove in production


// transistions
Vue.transition('board', {
    enterClass: 'bounceInDown',
    leaveClass: 'bounceOutDown'
});

// creating the vue instance here
// trying to have all my logic in the backend, only updating the view
// on model changes and passing moves back to the backend
var tictactoe = new Vue({
    el: '#tictactoe',
    data: {
        gameState: {
            started: false,
            fields: [],
        },
        //Special Move coding scheme
        RESTART: 10,
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
        makeMove: function(fieldNum){
            socket.send(fieldNum);
        },
    }
});
