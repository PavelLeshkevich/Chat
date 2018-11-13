window.socket = new WebSocket("ws://" + location.host + "/ws");

function sendMessage(msg)
{
    var len = msg.length
    socket.send(msg)
}

function handleSubmit()
{
    var el = document.getElementById("chat-msg")
    sendMessage(el.value)
    el.value = ''

    return false;
}


function setUpSocket(onmessage)
{
    socket.onopen = function() {
        console.log("Connection");
    }

    socket.onclose = function(event) {
        if (event.wasClean) {
            console.log('Connection closed');
        } else {
            console.log('ERROR: Connection reset');
            console.log('Code: ' + event.code + ' reason: ' + event.reason);
        }
    }

    socket.onmessage = onmessage;

    socket.onerror = function(error) {
        console.log("Error " + error.message);
    }
}

function displayMessage(msg)
{
    var container = document.getElementById("container")

    var div = document.createElement("div")
    div.className = 'message'

    var textNode = document.createTextNode(msg.data);

    div.appendChild(textNode)
    container.appendChild(div)
}