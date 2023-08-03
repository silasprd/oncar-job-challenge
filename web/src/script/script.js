let selectedCarId = null;

let apiUrl = "http://localhost:8080";
let webUrl = "http://localhost:3000/";

async function fetchAndRenderCars() {
  try {
    const response = await fetch(`${apiUrl}/cars`);
    const data = await response.json();
    const carList = document.getElementById("carList");

    data.forEach((car) => {
      const listItem = document.createElement("li");
      listItem.classList.add("card");

      // Armazena o ID do carro no elemento .card como um atributo personalizado
      listItem.setAttribute("data-car-id", car.ID);

      const txtBrand = document.createElement("span");
      txtBrand.classList.add("txt-brand");
      txtBrand.innerText = `${car.Brand}`;

      const txtModel = document.createElement("span");
      txtModel.classList.add("txt-model");
      txtModel.innerText = `Modelo: ${car.Model}`;

      const txtYear = document.createElement("span");
      txtYear.classList.add("txt-year");
      txtYear.innerText = `Ano: ${car.Year}`;

      const txtPrice = document.createElement("span");
      txtPrice.classList.add("txt-price");
      txtPrice.innerText = `R$ ${car.Price}`;

      const deleteButton = document.createElement("button");
      deleteButton.classList.add("delete-button");
      deleteButton.innerText = "X";

      listItem.id = "card";

      listItem.addEventListener("click", function (event) {
        if (event.target == deleteButton) return;
        selectedCarId = car.ID;
        if (selectedCarId !== null) {
          getCarIdAndOpenModal(selectedCarId);
        }
        event.stopPropagation();
      });

      carList.appendChild(listItem);
      listItem.appendChild(txtBrand);
      listItem.appendChild(txtModel);
      listItem.appendChild(txtYear);
      listItem.appendChild(txtPrice);
      listItem.appendChild(deleteButton);
      deleteButton.addEventListener("click", (event) =>
        deleteCar(event, car.ID)
      );
    });
    var closeButton = document.querySelector(".closed");
    closeButton.addEventListener("click", closeModalContact);
  } catch (error) {
    console.error("Erro ao buscar e renderizar a lista de carros:", error);
  }
}

function addCar(event) {
  event.preventDefault();

  const brandInput = document.getElementById("brand");
  const modelInput = document.getElementById("model");
  const yearInput = document.getElementById("year");
  const priceInput = document.getElementById("price");

  const priceString = priceInput.value.substring(3);
  const parseStringWithoutPoint = priceString.replace(/[^\d]/g, "");
  const priceFloat = parseFloat(parseStringWithoutPoint);

  const car = {
    brand: brandInput.value,
    model: modelInput.value,
    year: parseInt(yearInput.value),
    price: priceFloat,
  };

  console.log(car);

  fetch(`${apiUrl}/cars`, {
    method: "POST",
    header: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(car),
  })
    .then((response) => {
      if (response.status === 201) {
        window.alert("Carro criado com sucesso!");
        closeModalCar();
      } else {
        throw new Error(response.status);
      }
    })
    .finally(() => {
      window.location.reload();
    });
}

function addContact(event) {
  event.preventDefault();

  const nameInput = document.getElementById("name");
  const emailInput = document.getElementById("email");
  const phoneInput = document.getElementById("phone");
  const urlParams = new URLSearchParams(window.location.search);
  const carIdParam = urlParams.get("carId");

  const contact = {
    name: nameInput.value,
    email: emailInput.value,
    phone: phoneInput.value,
    carId: parseInt(carIdParam),
  };

  // Fazer requisição para API
  fetch(`${apiUrl}/contacts`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(contact),
  }).then((response) => {
    if (response.status === 201) {
      window.alert("Dados de contato enviados com sucesso!");
      closeModalContact();
    } else {
      throw new Error(response.status);
    }
  });
}

function deleteCar(event, carId) {
  if (confirm("Deseja realmente deletar este carro?")) {
    fetch(`${apiUrl}/cars/${carId}`, { method: "DELETE" }).then((response) => {
      if (response.status === 200) {
        window.location.reload();
        window.alert("Carro deletado com sucesso!");
      } else {
        throw new Error(response.status);
      }
    });
  } else {
    window.location.reload();
  }
}

async function getCarById(carId) {
  try {
    const response = await fetch(`${apiUrl}/cars/${carId}`);
    const data = await response.json();

    var getCar = document.getElementById("get-car");
    getCar.innerText = `${data.Brand} ${data.Model} ${data.Year}`;

    var getPriceField = document.getElementById("get-car-price");
    getPriceField.innerText = `R$ ${data.Price}`;

    await getContactsByCar(data.Contacts)

  } catch (error) {
    throw new Error(error);
  }
}

async function getContactsByCar(listContacts) {

  const contactList = document.getElementById("contact-list");

  if(listContacts.length == 0){
    const anyContacts = document.createElement("span");
    anyContacts.innerText = "Nenhum contato para este veículo"
    contactList.appendChild(anyContacts)
  }

  listContacts.forEach((contact) => {
    const contactItem = document.createElement("li");
    contactItem.classList.add("contact-item");

    const txtName = document.createElement("span");
    txtName.classList.add("txt-contact");
    txtName.innerText = `Nome: ${contact.Name}`;

    const txtEmail = document.createElement("span");
    txtEmail.classList.add("txt-email");
    txtEmail.innerText = `Email: ${contact.Email}`;

    const txtPhone = document.createElement("span");
    txtPhone.classList.add("txt-phone");
    txtPhone.innerText = `Telefone: ${contact.Phone}`;

    contactList.appendChild(contactItem);
    contactItem.appendChild(txtName);
    contactItem.appendChild(txtEmail);
    contactItem.appendChild(txtPhone);
  });
}

function formValidate(inputField, submitButton) {
  var count = 0;
  for (let index = 0; index < inputField.length; index++) {
    if (inputField[index].value === "") {
      count++;
    }
  }
  if (count == 0) {
    submitButton.disabled = false;
  } else {
    submitButton.disabled = true;
  }
}

function applyMask(input, mask) {
  const maskOptions = {
    mask: mask,
  };
  const masked = IMask(input, maskOptions);

  return masked;
}

async function openModalContact() {
  var modal = document.getElementById("form-contact");
  modal.style.display = "block";
  if (selectedCarId !== null) {
    await getCarById(selectedCarId);
  }
  var nameInput = document.getElementById("name");
  var emailInput = document.getElementById("email");
  var phoneInput = document.getElementById("phone");
  applyMask(phoneInput, "(00) 0 0000-0000");

  inputFields = [nameInput, emailInput, phoneInput];

  var submitButton = document.querySelector(".submit-button");
  submitButton.disabled = true;

  for (var index = 0; index < inputFields.length; index++) {
    inputFields[index].addEventListener("input", () => {
      formValidate(inputFields, submitButton);
    });
  }
}

function openModalCar() {
  var modal = document.getElementById("form-car");
  modal.style.display = "block";

  var brandInput = document.getElementById("brand");
  var modelInput = document.getElementById("model");
  var yearInput = document.getElementById("year");

  var priceInput = document.getElementById("price");
  applyMask(priceInput, [
    { mask: "" },
    {
      mask: "R$ num",
      lazy: false,
      blocks: {
        num: {
          mask: Number,
          thousandsSeparator: ".",
          mapToRadix: ["."],
        },
      },
    },
  ]);

  inputFields = [brandInput, modelInput, yearInput];

  var submitButton = document.getElementById("submit-car-button");
  submitButton.disabled = true;

  for (var index = 0; index < inputFields.length; index++) {
    inputFields[index].addEventListener("input", () => {
      formValidate(inputFields, submitButton);
    });
  }
}

async function changeContactList() {
  var formContact = document.querySelector(".content");
  var listContacts = document.querySelector(".car-contacts");

  if (formContact.style.display == "none") {
    formContact.style.display = "block";
  } else {
    formContact.style.display = "none";
  }

  if (listContacts.style.display == "flex") {
    listContacts.style.display = "none";
  } else {
    listContacts.style.display = "flex";
  }
}

function closeModalContact() {
  var modal = document.getElementById("form-contact");
  document.getElementById("contact-form").reset();
  modal.style.display = "none";
  history.pushState({}, "", webUrl);
  window.location.reload();
}

function closeModalCar() {
  var modal = document.getElementById("form-car");
  document.getElementById("contact-form").reset();

  modal.style.display = "none";
  history.pushState({}, "", webUrl);
}

function getCarIdAndOpenModal(carId) {
  const url = `${webUrl}?carId=${carId}`;
  history.pushState({}, "", url);
  openModalContact();
}

document.addEventListener("DOMContentLoaded", fetchAndRenderCars);

var submitButton = document.querySelector(".submit-button");
submitButton.addEventListener("click", addContact);

var submitCarButton = document.getElementById("submit-car-button");
submitCarButton.addEventListener("click", addCar);

var closeButton = document.querySelector(".closed");
closeButton.addEventListener("click", function () {
  closeModalContact();
});

var closeCar = document.querySelector(".closed-car");
closeCar.addEventListener("click", function () {
  closeModalCar();
});

var addCarButton = document.querySelector(".add-car-button");
addCarButton.addEventListener("click", function () {
  openModalCar();
});

var listContacts = document.querySelector(".span-contacts");
listContacts.addEventListener("click", function () {
  changeContactList();
});

var backButton = document.querySelector(".back-button");
backButton.addEventListener("click", function () {
  changeContactList();
});
