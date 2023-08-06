let selectedCarId = null;

let apiUrl = 'http://localhost:8080'
let webUrl = 'http://localhost:3000/'

let currentPage = 1;
let itemsPerPage = 12;
let maxPage = 0

async function fetchAndRenderCars() {
  try {
    const response = await fetch(`${apiUrl}/cars`);
    let data = await response.json();
    const carList = document.getElementById('carList');

    let dataFiltered = filteredData(data)

    const startIndex = (currentPage - 1) * itemsPerPage;
    const endIndex = startIndex + itemsPerPage;
    const carsPerPage = dataFiltered.slice(startIndex, endIndex);

    maxPage = Math.ceil(data.length / itemsPerPage);

    var closeButton = document.querySelector(".closed");
    closeButton.addEventListener("click", closeModalContact);

    if (dataFiltered == 0) {
      const noData = document.createElement("span")
      noData.innerText = 'Nenhum carro encontrado!'
      noData.classList.add("nodata-span")
      carList.appendChild(noData)
      carList.style.gridTemplateColumns = '1fr'
      return;
    } else {
      carList.style.gridTemplateColumns = '1fr 1fr 1fr 1fr'
    }

    carsPerPage.forEach(car => {
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

  } catch (error) {
    throw new Error(error)
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

async function getCarById() {
  try {
    const urlParams = new URLSearchParams(window.location.search);
    const carIdParam = urlParams.get("carId");
    const response = await fetch(`${apiUrl}/cars/${carIdParam}`)
    const data = await response.json()

    var getCar = document.getElementById("get-car")
    getCar.innerText = `${data.Brand} ${data.Model} ${data.Year}`

    var getPriceField = document.getElementById("get-car-price")
    getPriceField.innerText = `R$ ${data.Price}`
    
    getContactsByCar(data)

  } catch (error) {
    throw new Error(error)
  }
}

function getContactsByCar(data) {

  const contactList = document.getElementById("contact-list")

  if (data.Contacts.length > 0) {
    data.Contacts.forEach((contact) => {
      const listItem = document.createElement('li')
      listItem.classList.add("card");

      const txtName = document.createElement('span')
      txtName.classList.add("txt-name")
      txtName.innerText = `${contact.Name}`

      const txtEmail = document.createElement('span')
      txtEmail.classList.add("txt-email")
      txtEmail.innerText = `${contact.Email}`

      const txtPhone = document.createElement('span')
      txtPhone.classList.add("txt-phone")
      txtPhone.innerText = `${contact.Phone}`

      listItem.appendChild(txtName)
      listItem.appendChild(txtEmail)
      listItem.appendChild(txtPhone)
      contactList.appendChild(listItem)
    })
  } else {
    contactList.style.display = "flex"
    const noData = document.createElement('span')
    noData.innerText = 'Nenhum contato encontrado!'
    noData.classList.add("nodata-span")
    contactList.append(noData)
  }

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
    } else {
      return response.json()
    }
  })

}

function addContact(event) {

  //event.preventDefault()

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
          window.location.reload()
        }
      })
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
  modal.style.display = "flex";
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
  modal.style.display = "flex";

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

  for (var index = 0; index < inputFields.length; index++) {
    inputFields[index].addEventListener("input", () => { formValidate(inputFields, submitButton) })
  }
}

function openContacts() {
  var contactList = document.getElementById("car-contacts")
  var contactForm = document.querySelector(".content")

  contactForm.style.display = 'none'
  contactList.style.display = 'flex'
}

function backContactForm() {
  var contactList = document.getElementById("car-contacts")
  var contactForm = document.querySelector(".content")

  contactForm.style.display = 'block'
  contactList.style.display = 'none'
}

function closeModalContact() {
  var modal = document.getElementById("form-contact");
  modal.style.display = "none";
  history.pushState({}, "", webUrl);
  window.location.reload()
}

function closeModalCar() {
  var modal = document.getElementById("form-car");
  modal.style.display = "none";
  history.pushState({}, "", webUrl);
  window.location.reload()
}

function getCarIdAndOpenModal(carId) {
  const url = `${webUrl}?carId=${carId}`;
  history.pushState({}, "", url);
  openModalContact();
}

function updatePageInfo() {
  const currentPageElement = document.getElementById('currentPage');
  currentPageElement.textContent = currentPage;
}

document.addEventListener("DOMContentLoaded", fetchAndRenderCars);
document.addEventListener("DOMContentLoaded", updatePageInfo);

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
  closeModalCar()
});

var addCarButton = document.querySelector(".add-car-button")
addCarButton.addEventListener("click", function () {
  openModalCar()
});

var filterButton = document.getElementById("filter-button")
filterButton.addEventListener("click", function () {
  fetchAndRenderCars()
})

var spanContacts = document.querySelector(".span-contacts")
spanContacts.addEventListener("click", function () {
  openContacts();
})

document.getElementById("back-button").addEventListener('click', function () {
  backContactForm()
})

document.getElementById('prevPage').addEventListener('click', () => {
  if (currentPage > 1) {
    currentPage--;
    fetchAndRenderCars();
    updatePageInfo();
  }
});

document.getElementById('nextPage').addEventListener('click', () => {
  if (currentPage < maxPage) {
    currentPage++;
    fetchAndRenderCars();
    updatePageInfo();
  }
});