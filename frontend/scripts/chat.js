const host = "ws://localhost:8080/ws";
let token = sessionStorage.getItem("jwt-token");
let username = sessionStorage.getItem("username");
let conn;

let messages_area = document.getElementById("messages_area");

function sendMessage() {
  let message = document.getElementById("message");
  if (message != null) {
    conn.send(message.value);
  }
  return false;
}

window.onload = function () {
  document.getElementById("send_message_form").onsubmit = sendMessage;

  conn = new WebSocket(host + "?token=" + encodeURI(token));

  conn.onopen = function (evt) {
    console.log("ws connected");
  };

  conn.onmessage = async function (evt) {
    if (typeof evt.data === "string") {
      let el = document.createElement("p");
      el.textContent = evt.data;

    console.log(`Username: ${username}`)

    if (evt.data.startsWith(username)) {
      el.classList.add("my-message");
    }

      messages_area.append(el);
    } 
    else if (evt.data instanceof Blob) {
      let text = await evt.data.text();
      let obj = JSON.parse(text);
      if (obj.type == "error") {
        console.log(`Error: ${obj.message}`); // implement error handling
      }
    }
  };

  conn.onerror = function (evt) {
    alert(`not authorized`);
    // user not authorized
  };
  conn.onclose = function (evt) {
    console.log(`ws disconnected`);
  };
};
