console.log("oi do js");

// async function getData() {
//     const url = "localhost/8080";
//     const response = await fetch(url);
//     const json = await response.json();
//     console.log(json);
// }

// getData()

fetch('http://localhost:8080/api/')
  .then(response => {
    if (!response.ok) {
      throw new Error('Erro na requisição: ' + response.statusText);
    }
    return response.json(); // Converte a resposta em JSON
  })
  .then(data => {
    console.log(data); // Manipula os dados recebidos
  })
  .catch(error => {
    console.error('Erro:', error); // Lida com erros
  });

/*
Adicionar headers na requisição
entender melhor como fazer requisicoes js
*/
