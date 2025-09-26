const host = "ws://localhost:8080/ws";

window.onload = function () {
  let conn = new WebSocket(host);
  conn.onopen = function (evt) {
    console.log("ws connected");
  };
  conn.onmessage = function (evt) {
    console.log(`message: ${evt.data}`);
  };
  conn.onerror = function (evt) {
    console.log("ws error occured");
  };
  conn.onclose = function (evt) {
    console.log("ws disconnected");
  };
};
