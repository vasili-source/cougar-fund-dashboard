const VER = '20260209033737';
const BASE = '/cougar-fund-dashboard/data/';

function load(id){
  fetch(BASE + id + '.' + VER + '.json', { cache: 'no-store' })
    .then(r => r.json())
    .then(d => document.getElementById(id).textContent = d.value)
    .catch(() => document.getElementById(id).textContent = 'ERR');
}

['fedfunds','t10y','t2y','cpi','unrate','sp500'].forEach(load);

setTimeout(()=>{
  const a=parseFloat(document.getElementById('t10y').textContent);
  const b=parseFloat(document.getElementById('t2y').textContent);
  document.getElementById('curve').textContent =
    (!isNaN(a)&&!isNaN(b))?(a-b).toFixed(2):'ERR';
},200);
