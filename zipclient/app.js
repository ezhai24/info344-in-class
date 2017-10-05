var cityForm = document.getElementById('city-form');
cityForm.addEventListener("submit", e => {
  e.preventDefault();
 
  var city = document.getElementById('city').value;
  var states = new Set();
  var codes = new Set();
  fetch('http://localhost:4001/zips/' + city)
  .then((response) => {
    return response.json()
  })
  .then((data) => {
    data.forEach((zip) => {
      codes.add(zip.Code);
      states.add(zip.State);
    })
  })

  var statesSelect = document.getElementById('states');
  states.forEach((state) => {
    var option = document.createElement("option");
    option.setAttribute("value", state);
    statesSelect = statesSelect.appendChild(option)
  })
})
