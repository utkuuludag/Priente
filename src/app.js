const socket = new WebSocket('ws://localhost:8080/ws');

socket.addEventListener('open', function (event) {
    console.log('WebSocket is open now.');
});

socket.addEventListener('message', function (event) {
    const messagesDiv = document.getElementById('messages');
    const newMessage = document.createElement('div');
    newMessage.textContent = 'Message from server: ' + event.data;
    messagesDiv.appendChild(newMessage);
});

socket.addEventListener('close', function (event) {
    console.log('WebSocket is closed now.');
});

function sendMessage() {
    const input = document.getElementById('messageInput');
    const message = input.value;
    socket.send(JSON.stringify({ data: message }));
    input.value = '';
}
