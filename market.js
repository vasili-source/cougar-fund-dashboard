function load(id) {
  fetch('./data/' + id + '.json')
    .then(r => r.json())
    .then(d => document.getElementById(id).textContent = d.value)
    .catch(() => document.getElementById(id).textContent = 'N/A');
}

['fedfunds','t10y','t2y','cpi','unrate','sp500'].forEach(load);

setTimeout(() => {
  const a = parseFloat(document.getElementById('t10y').textContent);
  const b = parseFloat(document.getElementById('t2y').textContent);
  if (!isNaN(a) && !isNaN(b)) {
    document.getElementById('curve').textContent = (a - b).toFixed(2);
  } else {
    document.getElementById('curve').textContent = 'N/A';
  }
}, 300);
