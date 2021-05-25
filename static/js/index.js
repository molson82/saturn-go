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

getTwitchStatus();
