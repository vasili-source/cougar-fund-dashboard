function load(id) {
  fetch('./data/' + id + '.json', { cache: 'no-store' })
    .then(r => r.json())
    .then(d => document.getElementById(id).textContent = d.value)
    .catch(() => document.getElementById(id).textContent = 'N/A');
}

['fedfunds','t10y','t2y','cpi','unrate','sp500'].forEach(load);

setTimeout(() => {
  const a = parseFloat(document.getElementById('t10y').textContent);
  const b = parseFloat(document.getElementById('t2y').textContent);
  document.getElementById('curve').textContent =
    (!isNaN(a) && !isNaN(b)) ? (a - b).toFixed(2) : 'N/A';
}, 100);
