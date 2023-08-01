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
    closeButton.addEventListener("click", closeModalContact);
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
    if(response.status === 201){
      window.alert("Dados de contato enviados com sucesso!");
      closeModalContact();
    } else if (!response.ok) {
      console.log('Erro na requisição');
    } else {
      return response.json()
    }   
  })

}

function openModalContact() {
  var modal = document.getElementById("form-contact");
  modal.style.display = "block";
}

function openModalCar() {
  var modal = document.getElementById("form-car");
  modal.style.display = "block";
}

function closeModalContact() {
  var modal = document.getElementById("form-contact");
  modal.style.display = "none";
  history.pushState({}, "", webUrl);
}

function closeModalCar() {
  var modal = document.getElementById("form-car");
  modal.style.display = "none";
}

function getCarIdAndOpenModal(carId) {
  const url = `${webUrl}?carId=${carId}`;
  history.pushState({}, "", url);
  openModalContact();
}

document.addEventListener("DOMContentLoaded", fetchAndRenderCars);

// Lidar com o evento de "popstate" para manter o modal aberto quando a URL é alterada
window.addEventListener("popstate", function (event) {
  if (event.state && event.state.modalOpen) {
    openModal();
  } else {
    closeModal();
  }
});

var submitButton = document.querySelector(".submit-button")
submitButton.addEventListener("click", sendContact)

var closeButton = document.querySelector(".closed");
closeButton.addEventListener("click", function () {
  history.pushState({}, "");
  closeModalContact();
});

var closeCar = document.querySelector(".closed-car");
closeCar.addEventListener("click", function () {
  var modal = document.getElementById("form-car");
  if(modal.style.display == "block"){
    modal.style.display = "none";
  }

  console.log("asdajsdalksdj")
});

var addCarButton = document.querySelector(".add-car-button")
addCarButton.addEventListener("click", function () {
  var formCar = document.getElementById("form-car")
  formCar.style.display = "block"
})
