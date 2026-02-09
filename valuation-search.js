document.addEventListener("DOMContentLoaded", () => {
  console.log("valuation-search.js loaded");

  let stocks = [];

  const search = document.getElementById("search");
  const results = document.getElementById("results");
  const output = document.getElementById("output");

  if (!search || !results || !output) {
    alert("Search UI failed to load");
    return;
  }

  fetch("./data/stocks_index.json")
    .then(r => r.json())
    .then(d => {
      stocks = d;
      console.log("Loaded stocks:", stocks.length);
    })
    .catch(err => {
      console.error("Failed to load stock index", err);
      alert("Failed to load stock data");
    });

  search.addEventListener("input", () => {
    const q = search.value.toLowerCase().trim();
    results.innerHTML = "";
    if (!q) return;

    stocks
      .filter(s =>
        s.ticker.toLowerCase().includes(q) ||
        s.name.toLowerCase().includes(q)
      )
      .slice(0, 10)
      .forEach(s => {
        const li = document.createElement("li");
        li.textContent = `${s.ticker} — ${s.name}`;
        li.onclick = () => loadValuation(s.ticker);
        results.appendChild(li);
      });
  });

  function loadValuation(ticker) {
    fetch("data/intrinsic/" + ticker + ".json")
      .then(r => r.json())
      .then(v => {
        output.textContent =
          `Ticker: ${v.ticker}\nIntrinsic Value: $${Number(v.value).toFixed(2)}`;
      })
      .catch(() => {
        output.textContent =
          "Intrinsic value is still being generated for this stock.";
      });
  }
});
