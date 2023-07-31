let selectedCarId = null;

let apiUrl = 'http://localhost:8080/cars'
let webUrl = 'http://localhost:3000/'

async function fetchAndRenderCars() {
  try {
    const response = await fetch(apiUrl);
    const data = await response.json();
    console.log(data)
    const carList = document.getElementById('carList');

    data.forEach(car => {
      const listItem = document.createElement('li')
      listItem.classList.add("card");

      // Armazena o ID do carro no elemento .card como um atributo personalizado
      listItem.setAttribute("data-car-id", car.ID);

      const txtBrand = document.createElement('span')
      txtBrand.classList.add("txt-brand")
      txtBrand.innerText = `${car.Brand}`

      const txtModel = document.createElement('span')
      txtModel.classList.add("txt-model")
      txtModel.innerText = `Modelo: ${car.Model}`

      const txtYear = document.createElement('span')
      txtYear.classList.add("txt-year")
      txtYear.innerText = `Ano: ${car.Year}`

      const txtPrice = document.createElement('span')
      txtPrice.classList.add("txt-price")
      txtPrice.innerText = `R$ ${car.Price}`

      listItem.id = "card"

      listItem.addEventListener("click", function (event) {
        selectedCarId = car.ID
        console.log(car)
        if (selectedCarId !== null) {
          getCarIdAndOpenModal(selectedCarId);
        }
        event.stopPropagation();
      });

      carList.appendChild(listItem);
      listItem.appendChild(txtBrand)
      listItem.appendChild(txtModel)
      listItem.appendChild(txtYear)
      listItem.appendChild(txtPrice)
    });
    var closeButton = document.querySelector(".closed");
    closeButton.addEventListener("click", closeModal);
  } catch (error) {
    console.error('Erro ao buscar e renderizar a lista de carros:', error);
  }
}

function sendContact() {
  const nameInput = document.getElementById("name");
  const emailInput = document.getElementById("email");
  const phoneInput = document.getElementById("phone");
  const urlParams = new URLSearchParams(window.location.search);
  const carIdParam = urlParams.get("carId");

  const contact = {
    name: nameInput.value,
    email: emailInput.value,
    phone: phoneInput.value,
    carId: parseInt(carIdParam)
  }

  // Fazer requisição para API
  fetch('http://localhost:8080/contacts', {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json'
      },
      body: JSON.stringify(contact)
  }).then(async (response) => {
    console.log(response)
    if(response.status === 201){
      window.alert("Dados de contato enviados com sucesso!");
      closeModal();
    } else if (!response.ok) {
      console.log('Erro na requisição');
    } else {
      return response.json()
    }   
  })

}

// Função para abrir o modal
function openModal() {
  var modal = document.getElementById("form");
  modal.style.display = "block";
}

// Função para fechar o modal
function closeModal() {
  var modal = document.getElementById("form");
  modal.style.display = "none";
  history.pushState({}, "", webUrl);
}

// Função para alterar a URL, abrir o modal
function getCarIdAndOpenModal(carId) {
  const url = `${webUrl}?carId=${carId}`;
  history.pushState({}, "", url);
  openModal();
}

// Chamar a função fetchAndRenderCars() quando o conteúdo da página for carregado
document.addEventListener("DOMContentLoaded", fetchAndRenderCars);

// Lidar com o evento de "popstate" para manter o modal aberto quando a URL é alterada
window.addEventListener("popstate", function (event) {
  if (event.state && event.state.modalOpen) {
    openModal();
  } else {
    closeModal();
  }
});

// Enviar dados de contato
var submitButton = document.querySelector(".submit-button")
submitButton.addEventListener("click", sendContact)

// Fechar o modal
var closeButton = document.querySelector(".closed");
closeButton.addEventListener("click", function () {
  history.pushState({}, "");
  closeModal();
});