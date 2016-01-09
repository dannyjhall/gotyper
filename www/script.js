// Forward url params to ws connection
var params = window.location.href.split("/").pop();
var socket = new WebSocket("wss://gotyper.herokuapp.com/connect" + params);

socket.onmessage = function(msg) {

    // Show cursor and enable key presses when ready
    if(msg.data == "goTyper:Ready") {
        document.getElementById("cursor").style.display = "inline-block";
        document.onkeydown = function(e) {
            if (e.keyCode == 8) { // Backspace
                e.returnValue = false;
            }
            socket.send(e.keyCode);
        }
    } else {
        document.getElementById('terminal').innerHTML += msg.data;
    }
    window.scrollBy(0,50);
};
