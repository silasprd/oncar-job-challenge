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
	dbConn, err := db.OpenConnection()
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

	err = db.AutoMigrateTables(dbConn)
	if err != nil {
		fmt.Println("Erro ao realizar auto migrate das tabelas:", err)
		return
	}

	carService := service.NewCarService(dbConn)

	carController := controller.NewCarController(carService)

	router := mux.NewRouter()

	router.HandleFunc("/carros", carController.ListCarsHandler).Methods("GET")
	router.HandleFunc("/carros", carController.AddCarHandler).Methods("POST")
	router.HandleFunc("/carros/{id}", carController.GetCarHandler).Methods("GET")
	router.HandleFunc("/carros/{id}", carController.DeleteCarHandler).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", router)
}
