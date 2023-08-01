let selectedCarId = null;

let apiUrl = 'http://localhost:8080'
let webUrl = 'http://localhost:3000/'

async function fetchAndRenderCars() {
  try {
    const response = await fetch(`${apiUrl}/cars`);
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

function addCar() {
  const brandInput = document.getElementById("brand");
  const modelInput = document.getElementById("model");
  const yearInput = document.getElementById("year");
  const priceInput = document.getElementById("price");

  const car = {
    brand: brandInput.value,
    model: modelInput.value,
    year: parseInt(yearInput.value),
    price: parseInt(priceInput.value),
  }

  console.log(car)

  fetch(`${apiUrl}/cars`, {
    method: 'POST',
    header: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(car)
  }).then((response) => {
    if(response.status === 201){
      window.alert("Carro criado co sucesso!")
    } else if(!response.ok){
      console.log("Erro na requisição")
    } else {
      return response.json()
    }
  })
}

function addContact(event) {

  event.preventDefault()

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
  fetch(`${apiUrl}/contacts`, {
      method: 'POST',
      headers: {
          'Content-Type': 'application/json'
      },
      body: JSON.stringify(contact)
  }).then((response) => {
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

async function getCarById(){
  try {
    const urlParams = new URLSearchParams(window.location.search);
    const carIdParam = urlParams.get("carId");
    const response = await fetch(`${apiUrl}/cars/${carIdParam}`)
    const data = await response.json()

    var getCar = document.getElementById("get-car")
    getCar.innerText = `${data.Brand} ${data.Model} ${data.Year}`

    var getPriceField = document.getElementById("get-car-price")
    getPriceField.innerText = `R$ ${data.Price}`

  } catch (error) {
    console.log(error)
  }
}

function openModalContact() {
  var modal = document.getElementById("form-contact");
  modal.style.display = "block";
  getCarById()
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

var submitButton = document.querySelector(".submit-button")
submitButton.addEventListener("click", addContact)

var submitCarButton = document.getElementById("submit-car-button")
submitCarButton.addEventListener("click", addCar)

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
});

var addCarButton = document.querySelector(".add-car-button")
addCarButton.addEventListener("click", function () {
  var formCar = document.getElementById("form-car")
  formCar.style.display = "block"
})

