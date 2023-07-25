package main

import (
	"fmt"
	"net/http"

	"oncar-job-challenge/api/controller"
	"oncar-job-challenge/api/service"
	"oncar-job-challenge/db"

	"github.com/gorilla/mux"
)

func main() {
	dbConn, err := db.ConnectDB()
	if err != nil {
		fmt.Println("Erro ao conectar ao banco de dados:", err)
		return
	}
	defer func() {
		sqlDB, err := dbConn.DB()
		if err != nil {
			fmt.Println("Erro ao obter a conexão do banco de dados:", err)
			return
		}
		sqlDB.Close()
	}()

	carService := service.NewCarService(dbConn)

	carController := controller.NewCarController(carService)

	router := mux.NewRouter()

	router.HandleFunc("/carros", carController.ListCarsHandler).Methods("GET")
	router.HandleFunc("/carros", carController.AddCarHandler).Methods("POST")
	router.HandleFunc("/carros/{id}", carController.GetCarHandler).Methods("GET")

	fmt.Println("Servidor rodando na porta 8090")
	http.ListenAndServe(":8090", router)
}
