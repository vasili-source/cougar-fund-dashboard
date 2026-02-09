fetch('/cougar-fund-dashboard/data/health.json', { cache: 'no-store' })
  .then(r => r.json())
  .then(d => {
    document.body.insertAdjacentHTML(
      'afterbegin',
      '<div style="background:#d1fae5;padding:8px;font-weight:600">' +
      'DATA HEALTH: ' + d.value +
      '</div>'
    );
  })
  .catch(() => {
    document.body.insertAdjacentHTML(
      'afterbegin',
      '<div style="background:#fee2e2;padding:8px;font-weight:600">' +
      'DATA HEALTH: FAILED' +
      '</div>'
    );
  });
