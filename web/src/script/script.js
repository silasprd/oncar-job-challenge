let selectedCarId = null;

let apiUrl = 'http://localhost:8080'
let webUrl = 'http://localhost:3000/'

async function fetchAndRenderCars() {
  try {
    const response = await fetch(`${apiUrl}/cars`);
    let data = await response.json();
    console.log(data)
    const carList = document.getElementById('carList');

    filteredData(data).forEach(car => {
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

      const deleteButton = document.createElement('button')
      deleteButton.classList.add("delete-button")
      deleteButton.innerText = 'X'

      listItem.id = "card"

      listItem.addEventListener("click", function (event) {
        if (event.target == deleteButton) return
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
      listItem.appendChild(deleteButton)
      deleteButton.addEventListener("click", (event) => deleteCar(event, car.ID))
    });
    var closeButton = document.querySelector(".closed");
    closeButton.addEventListener("click", closeModalContact);
  } catch (error) {
    console.error('Erro ao buscar e renderizar a lista de carros:', error);
  }
}

function filteredData(data) {

  const filterType = document.getElementById('filterType').value;
  const filterValue = document.getElementById('filterValue').value.toLowerCase();

  const filteredData = data.filter(car => {
    switch (filterType) {
      case 'brand':
        return car.Brand.toLowerCase() === filterValue;
      case 'model':
        return car.Model.toLowerCase().includes(filterValue);
      case 'year':
        return car.Year.toString() === filterValue;
      default:
        return true;
    }
  });

  if (filterValue) {
    carList.innerHTML = '';
    data = filteredData
  } else {
    carList.innerHTML = ''
  }

  return data

}

function addCar(event) {

  // event.preventDefault()

  const brandInput = document.getElementById("brand");
  const modelInput = document.getElementById("model");
  const yearInput = document.getElementById("year");
  const priceInput = document.getElementById("price");

  const priceString = priceInput.value.substring(3)
  const parseStringWithoutPoint = priceString.replace(/[^\d]/g, '')
  const priceFloat = parseFloat(parseStringWithoutPoint)

  const car = {
    brand: brandInput.value,
    model: modelInput.value,
    year: parseInt(yearInput.value),
    price: priceFloat,
  }

  console.log(car)

  fetch(`${apiUrl}/cars`, {
    method: 'POST',
    header: {
      'Content-Type': 'application/json'
    },
    body: JSON.stringify(car)
  }).then((response) => {
    if (response.status === 201) {
      window.alert("Carro criado co sucesso!")
      closeModalCar();
    } else if (!response.ok) {
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
    if (response.status === 201) {
      window.alert("Dados de contato enviados com sucesso!");
      closeModalContact();
    } else if (!response.ok) {
      console.log('Erro na requisição');
    } else {
      return response.json()
    }
  })

}

function deleteCar(event, carId) {
  if (window.confirm("Você deseja deletar este carro ?")) {
    fetch(`${apiUrl}/cars/${carId}`, { method: 'DELETE' }).
      then((response) => {
        if (response.status === 200) {
          window.alert("Carro deletado com sucesso!")
          console.log(response.status)
          window.location.reload()
        }
      })
  }
}

function getCarById() {
  try {
    const urlParams = new URLSearchParams(window.location.search);
    const carIdParam = urlParams.get("carId");
    const response = fetch(`${apiUrl}/cars/${carIdParam}`)
    const data = response.json()

    var getCar = document.getElementById("get-car")
    getCar.innerText = `${data.Brand} ${data.Model} ${data.Year}`

    var getPriceField = document.getElementById("get-car-price")
    getPriceField.innerText = `R$ ${data.Price}`

  } catch (error) {
    console.log(error)
  }
}

function formValidate(inputField, submitButton) {
  var count = 0
  for (let index = 0; index < inputField.length; index++) {
    if (inputField[index].value === "") {
      count++
    }
  }
  if (count == 0) {
    submitButton.disabled = false
  } else {
    submitButton.disabled = true
  }
}

function applyMask(input, mask) {
  const maskOptions = {
    mask: mask
  }
  const masked = IMask(input, maskOptions)

  return masked
}

function openModalContact() {
  var modal = document.getElementById("form-contact");
  modal.style.display = "block";
  getCarById()
  var nameInput = document.getElementById("name")
  var emailInput = document.getElementById("email")
  var phoneInput = document.getElementById("phone")
  applyMask(phoneInput, '(00) 0 0000-0000')

  inputFields = [nameInput, emailInput, phoneInput]

  var submitButton = document.querySelector(".submit-button")
  submitButton.disabled = true

  for (var index = 0; index < inputFields.length; index++) {
    inputFields[index].addEventListener("input", () => { formValidate(inputFields, submitButton) })
  }
}

function openModalCar() {
  var modal = document.getElementById("form-car");
  modal.style.display = "block";

  var brandInput = document.getElementById("brand")
  var modelInput = document.getElementById("model")
  var yearInput = document.getElementById("year")

  var priceInput = document.getElementById("price")
  applyMask(priceInput, [
    { mask: '' },
    {
      mask: 'R$ num',
      lazy: false,
      blocks: {
        num: {
          mask: Number,
          thousandsSeparator: '.',
          padFractionalZeros: true,
          mapToRadix: ['.'],
        }
      }
    }
  ])

  inputFields = [brandInput, modelInput, yearInput]

  var submitButton = document.getElementById("submit-car-button")
  submitButton.disabled = true

  console.log("asjdahkjh")

  for (var index = 0; index < inputFields.length; index++) {
    inputFields[index].addEventListener("input", () => { formValidate(inputFields, submitButton) })
  }
}

function closeModalContact() {
  var modal = document.getElementById("form-contact");
  modal.style.display = "none";
  history.pushState({}, "", webUrl);
}

function closeModalCar() {
  var modal = document.getElementById("form-car");
  modal.style.display = "none";
  history.pushState({}, "", webUrl);
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
  if (modal.style.display == "block") {
    modal.style.display = "none";
  }
});

var addCarButton = document.querySelector(".add-car-button")
addCarButton.addEventListener("click", function () {
  openModalCar()
});

var filterButton = document.getElementById("filter-button")
filterButton.addEventListener("click", function () {
  fetchAndRenderCars()
})