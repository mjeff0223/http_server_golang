fetch('http://localhost:8080/api/v1/animes/1')
  .then(response => response.json())
  .then(data => {
    console.log(data);
    // Do something with the returned data
  })
  .catch(error => console.error(error));