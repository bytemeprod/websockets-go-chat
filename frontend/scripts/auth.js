const loginForm = document.getElementById("login-form");

document.getElementById("login-btn").addEventListener("click", () => {
  const username = document.getElementById("login-username").value.trim();
  if (username) {
    alert(`Logging in...`);
    // request for backend
  } else {
    alert("Please, type your username");
  }
});
