console.log("Index.js loaded");

const baseURL = window.location.origin;
const options = {
  method: "GET",
};

function getTwitchStatus() {
  fetch(baseURL + "/api/redis/twitch-status", options)
    .then((response) => response.json())
    .then((response) => {
      console.log("res from fetch: " + response.twitchStatus);
      if (response.twitchStatus === "online") {
        notifyOnline();
      } else {
        notifyOffline();
      }
    });

  setTimeout(getTwitchStatus, 1000 * 60);
}

function notifyOnline() {
  document.querySelector("#dbc5f9ab-ca75-4498-8ed9-741cc4531530").classList.replace("text-black", "flash-icon");
  document.querySelector("#twitch-live-pill").classList.replace("hidden", "block");
}

function notifyOffline() {
  document.querySelector("#dbc5f9ab-ca75-4498-8ed9-741cc4531530").classList.replace("flash-icon", "text-black");
  document.querySelector("#twitch-live-pill").classList.replace("block", "hidden");
}

document.addEventListener("DOMContentLoaded", () => {
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

getTwitchStatus();
