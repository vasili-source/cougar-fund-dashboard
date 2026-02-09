const BASE = 'https://api.stlouisfed.org/fred/series/observations';
const KEY = localStorage.getItem('FRED_API_KEY') || '';

function load(id, series) {
  fetch(BASE + '?series_id=' + series + '&api_key=' + KEY + '&file_type=json')
    .then(r => r.json())
    .then(d => {
      const v = d.observations[d.observations.length - 1].value;
      document.getElementById(id).textContent = v;
    })
    .catch(() => {
      document.getElementById(id).textContent = 'N/A';
    });
}

load('fedfunds','FEDFUNDS');
load('t10y','DGS10');
load('t2y','DGS2');
load('cpi','CPIAUCSL');
load('unrate','UNRATE');
load('sp500','SP500');

setTimeout(() => {
  const a = parseFloat(document.getElementById('t10y').textContent);
  const b = parseFloat(document.getElementById('t2y').textContent);
  if (!isNaN(a) && !isNaN(b)) {
    document.getElementById('curve').textContent = (a - b).toFixed(2);
  }
}, 1500);
