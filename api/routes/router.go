package routes

import (
	"oncar-job-challenge/core/controller"
	"oncar-job-challenge/core/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

// Configurar as rotas
func ConfigureRoutes(dbConn *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	carService := service.NewCarService(dbConn)
	contactService := service.NewContactService(dbConn, carService)

	carController := controller.NewCarController(carService)
	contactController := controller.NewContactController(contactService)

	// Defina suas rotas aqui
	router.HandleFunc("/cars", carController.ListCarsHandler).Methods("GET")
	router.HandleFunc("/cars/{id}", carController.GetCarHandler).Methods("GET")
	router.HandleFunc("/cars", carController.AddCarHandler).Methods("POST")
	router.HandleFunc("/cars/{id}", carController.DeleteCarHandler).Methods("DELETE")
	router.HandleFunc("/contacts", contactController.ListContactsHandler).Methods("GET")
	router.HandleFunc("/contacts/{id}", contactController.GetContactHandler).Methods("GET")
	router.HandleFunc("/contacts", contactController.AddContactHandler).Methods("POST")

	return router
}
