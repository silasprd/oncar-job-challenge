package routes

import (
	"oncar-job-challenge/api/controller"
	"oncar-job-challenge/api/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Configurar as rotas
func ConfigureRoutes(dbConn *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	carService := service.NewCarService(dbConn)

	carController := controller.NewCarController(carService)

	// Defina suas rotas aqui
	router.HandleFunc("/cars", carController.ListCarsHandler).Methods("GET")
	router.HandleFunc("/cars/{id}", carController.GetCarHandler).Methods("GET")
	router.HandleFunc("/cars", carController.AddCarHandler).Methods("POST")
	router.HandleFunc("/cars/{id}", carController.DeleteCarHandler).Methods("DELETE")

	return router
}
