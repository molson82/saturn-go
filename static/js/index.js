console.log("Index.js loaded");

const baseURL = window.location.origin;
const options = {
  method: "GET",
};
let ws;

function connectWS() {
  if (ws) {
    return false;
  }

  ws = new WebSocket("ws://" + window.location.host + "/api/ws/twitch-status");
  ws.onopen = () => {
    console.log("OPEN");
    sendPing();
  };

  ws.onclose = () => {
    console.log("CLOSE");
    ws = null;
  };

  ws.onmessage = (evt) => {
    console.log("RESPONSE: " + evt.data);
    const obj = JSON.parse(evt.data);
    if (obj.twitchStatus === "online") {
      notifyOnline();
    } else {
      notifyOffline();
    }
  };

  ws.onerror = (evt) => {
    console.log("ERROR: " + evt.data);
  };
}

function sendPing() {
  ws.send("ping ws");

  setTimeout(sendPing, 1000 * 60);
}

function notifyOnline() {
  if (window.location.pathname.includes("/index")) {
    document.querySelector("#dbc5f9ab-ca75-4498-8ed9-741cc4531530").classList.replace("text-black", "flash-icon");
    document.querySelector("#twitch-live-pill").classList.replace("hidden", "block");
  }
}

function notifyOffline() {
  if (window.location.pathname.includes("/index")) {
    document.querySelector("#dbc5f9ab-ca75-4498-8ed9-741cc4531530").classList.replace("flash-icon", "text-black");
    document.querySelector("#twitch-live-pill").classList.replace("block", "hidden");
  }
}

document.addEventListener("DOMContentLoaded", () => {
  connectWS();

  const darkmode = localStorage.getItem("wv-site-darkmode");
  if (darkmode === "true") {
    document.querySelector("html").classList.add("dark");
    document.querySelector("#darkModeToggle").checked = true;
  }

  document.querySelector("#slideMenuArrow").addEventListener("click", () => {
    const slideMenu = document.querySelector("#slideMenu");
    if (slideMenu.classList.contains("slideHover")) {
      slideMenu.classList.remove("slideHover");
    } else {
      slideMenu.classList.add("slideHover");
    }
  });

  document.querySelector("#app").addEventListener("click", () => {
    console.log("click app");
    document.querySelector("#slideMenu").classList.remove("slideHover");
  });

  document.querySelector("#darkModeToggle").addEventListener("click", (e) => {
    if (e.target.checked) {
      document.querySelector("html").classList.add("dark");
      localStorage.setItem("wv-site-darkmode", "true");
    } else {
      document.querySelector("html").classList.remove("dark");
      localStorage.setItem("wv-site-darkmode", "false");
    }
  });
});
