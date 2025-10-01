const loginForm = document.getElementById("login-form");

let host = "http://localhost:8080";
let loginPath = "/login"
let chatPath = "/chat.html"

document.getElementById("login-btn").addEventListener("click", async () => {
  let username = document.getElementById("login-username").value.trim();

  if(!username) {
    alert("Please, type your username");
    return;
  }

  let data = {
      username: username,
  };

  let response = await fetch(host + loginPath, {
    method: "POST",
    body: JSON.stringify(data),
  });

  let status = response.status;
  switch (status) {
    case 500:
      alert("Internal Server Error");
      break;
    case 400:
      alert("Bad Request");
      break;
    case 409:
      alert("Username already exists");
      break;
    case 200:
      let resp = await response.json();
      let token = resp["token"];
      sessionStorage.setItem("jwt-token", token);
      sessionStorage.setItem("username", username);
      window.location.href = host + chatPath;
      break;
    default:
      alert("Unknown Error");
      break;
  }
});
