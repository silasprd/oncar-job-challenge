package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	model "oncar-job-challenge/core/model"
	"oncar-job-challenge/core/service"

	"github.com/gorilla/mux"
)

type ContactController struct {
	contactService *service.ContactService
}

func NewContactController(contactService *service.ContactService) *ContactController {
	return &ContactController{contactService: contactService}
}

func (c *ContactController) AddContactHandler(w http.ResponseWriter, r *http.Request) {

	var contact model.Contact
	err := json.NewDecoder(r.Body).Decode(&contact)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.contactService.AddContact(contact)
	if err != nil {
		http.Error(w, "Erro ao adicionar o contato", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (c *ContactController) ListContactsHandler(w http.ResponseWriter, r *http.Request) {

	contacts, err := c.contactService.GetAllContact()
	if err != nil {
		http.Error(w, "Erro ao obter a lista de carros", http.StatusInternalServerError)
		return
	}

	contactsJSON, err := json.Marshal(contacts)
	if err != nil {
		http.Error(w, "Erro ao serializar a lista de contatos", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(contactsJSON)

}

func (c *ContactController) GetContactHandler(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID do carro inválido", http.StatusBadRequest)
		return
	}

	contact, err := c.contactService.GetContactByID(uint(id))
	if err != nil {
		http.Error(w, "Carro não encontrado", http.StatusNotFound)
		return
	}

	contactJSON, err := json.Marshal(contact)
	if err != nil {
		http.Error(w, "Erro ao serializar o carro", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(contactJSON)

}
