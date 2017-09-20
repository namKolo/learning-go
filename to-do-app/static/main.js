function connect() {
  var ws = new WebSocket('ws://localhost:3000/ws/all');
  var list = $('#todo-list');

  ws.onmessage = function(e) {
    var data = JSON.parse(e.data);
    console.log(data);
  };
}
