document.addEventListener("DOMContentLoaded", () => {
  const search = document.getElementById("search");
  const results = document.getElementById("results");
  const output = document.getElementById("output");
  const status = document.getElementById("status");

  let stocks = [];

  fetch("/cougar-fund-dashboard/data/stocks_index.json")
    .then(r => r.json())
    .then(d => {
      stocks = d;
      status.textContent = "Loaded " + stocks.length + " stocks";
    });

  search.addEventListener("input", () => {
    const q = search.value.toLowerCase().trim();
    results.innerHTML = "";
    if (!q) return;

    stocks
      .filter(s =>
        (s.ticker && s.ticker.toLowerCase().includes(q)) ||
        (s.title && s.title.toLowerCase().includes(q))
      )
      .slice(0, 10)
      .forEach(s => {
        const li = document.createElement("li");
        li.textContent = s.ticker + " — " + s.title;
        li.onclick = () => loadValuation(s.ticker);
        results.appendChild(li);
      });
  });

  function loadValuation(ticker) {
    const path = "/cougar-fund-dashboard/data/intrinsic/" + ticker + ".json";

    fetch(path, { method: "HEAD" })
      .then(r => {
        if (!r.ok) {
          output.textContent =
            "Intrinsic value is still being generated for this stock.";
          return;
        }

        fetch(path)
          .then(r => r.json())
          .then(v => {
            output.textContent =
              "Ticker: " + v.ticker +
              "\\nIntrinsic Value: $" + Number(v.value).toFixed(2);
          });
      });
  }
});
