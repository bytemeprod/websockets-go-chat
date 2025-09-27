const toLoginLink = document.getElementById("to-login");
const toRegisterLink = document.getElementById("to-register");
const regForm = document.getElementById("registration-form");
const loginForm = document.getElementById("login-form");

toLoginLink.addEventListener("click", () => {
  regForm.style.display = "none";
  loginForm.style.display = "block";
});

toRegisterLink.addEventListener("click", () => {
  loginForm.style.display = "none";
  regForm.style.display = "block";
});

document.getElementById("reg-btn").addEventListener("click", () => {
  const username = document.getElementById("reg-username").value.trim();
  const password = document.getElementById("reg-password").value.trim();
  if (username && password) {
    alert(`Successfull registration\nUsername: ${username}`);
    // implement request to backend
  } else {
    alert("Please, fill in all fields");
  }
});

document.getElementById("login-btn").addEventListener("click", () => {
  const username = document.getElementById("login-username").value.trim();
  const password = document.getElementById("login-password").value.trim();
  if (username && password) {
    alert(`Logging in...`);
    // implement request to backend
  } else {
    alert("Please, fill in all fields");
  }
});
