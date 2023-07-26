package controller

import (
	"encoding/json"
	"net/http"
	"strconv"

	models "oncar-job-challenge/api/model"
	"oncar-job-challenge/api/service"

	"github.com/gorilla/mux"
)

type CarController struct {
	carService *service.CarService
}

func NewCarController(carService *service.CarService) *CarController {
	return &CarController{carService: carService}
}

func (c *CarController) AddCarHandler(w http.ResponseWriter, r *http.Request) {
	var car models.Car
	err := json.NewDecoder(r.Body).Decode(&car)
	if err != nil {
		http.Error(w, "Dados inválidos", http.StatusBadRequest)
		return
	}

	err = c.carService.AddCar(car)
	if err != nil {
		http.Error(w, "Erro ao adicionar o carro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (c *CarController) ListCarsHandler(w http.ResponseWriter, r *http.Request) {
	cars, err := c.carService.GetAllCars()
	if err != nil {
		http.Error(w, "Erro ao obter a lista de carros", http.StatusInternalServerError)
		return
	}

	carJSON, err := json.Marshal(cars)
	if err != nil {
		http.Error(w, "Erro ao serializar a lista de carros", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(carJSON)
}

func (c *CarController) GetCarHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID do carro inválido", http.StatusBadRequest)
		return
	}

	car, err := c.carService.GetCarByID(uint(id))
	if err != nil {
		http.Error(w, "Carro não encontrado", http.StatusNotFound)
		return
	}

	carJSON, err := json.Marshal(car)
	if err != nil {
		http.Error(w, "Erro ao serializar o carro", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(carJSON)
}

func (c *CarController) DeleteCarHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "ID do carro inválido", http.StatusBadRequest)
		return
	}

	if err := c.carService.DeleteCar(uint(id)); err != nil {
		http.Error(w, "Erro ao deletar o carro", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
