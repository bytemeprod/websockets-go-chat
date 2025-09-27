const host = "ws://localhost:8080/ws";
let conn;

function sendMessage() {
  var message = document.getElementById("message");
  if (message != null) {
    conn.send(message.value);
  }
  return false;
}

window.onload = function () {
  document.getElementById("send_message_form").onsubmit = sendMessage;

  conn = new WebSocket(host);
  conn.onopen = function (evt) {
    console.log("ws connected");
  };
  conn.onmessage = function (evt) {
    let messages_area = document.getElementById("messages_area");
    let el = document.createElement("p");
    el.textContent = evt.data;
    messages_area.append(el);
  };
  conn.onerror = function (evt) {
    console.log("ws error occured");
  };
  conn.onclose = function (evt) {
    console.log("ws disconnected");
  };
};
