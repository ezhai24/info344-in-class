setInterval(() => {
  fetch('http://localhost:4000/memory')
  .then((response) => {
    return response.json()
  })
  .then((data) => {
    document.getElementById('alloc').innerHTML = 'Allocated Memory: ' + data.Alloc
  })
}, 1000);