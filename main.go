package main

import (
	"fmt"
	"net/http"

	"oncar-job-challenge/db"
	"oncar-job-challenge/routes"
)

func main() {
	dbConn, err := db.OpenConnection("./.env")
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

	router := routes.ConfigureRoutes(dbConn)
	fmt.Println("Servidor rodando na porta 8080")
	http.ListenAndServe(":8080", router)
}
