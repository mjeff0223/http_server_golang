const getDataBtn = document.getElementById('getDataBtn');
const dataElement = document.getElementById('data');

getDataBtn.addEventListener('click', () => {
 fetch('http://localhost:8080/api/v1/animes/')
  .then(response => response.json())
  .then(data => {
    console.log(data);
    // Do something with the returned data
  })
  .catch(error => console.error(error));
})