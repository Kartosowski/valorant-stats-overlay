const rankEl = document.getElementById("rank");
const scoreEl = document.getElementById("score");
const changeEl = document.getElementById("change");
const iconEl = document.getElementById("rank-icon");

async function fetchStats() {
  try {
    const res = await fetch("http://localhost:{PORT}/stats"); 
    const data = await res.json();
    const latest = data.data[0];

    rankEl.textContent = latest.currenttierpatched;
    scoreEl.textContent = `${latest.ranking_in_tier}rr`;
    changeEl.textContent = `${latest.mmr_change_to_last_game >= 0 ? "+" : ""}${latest.mmr_change_to_last_game}`;
    iconEl.src = latest.images.small;
  } catch (err) {
    console.error("Fetch error:", err);
  }
}

// fetch instantly
fetchStats();
// update every minute
setInterval(fetchStats, 60000);
