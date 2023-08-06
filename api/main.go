package main

import (
	"fmt"
	"net/http"
	"oncar-job-challenge/db"
	"oncar-job-challenge/routes"
	"path/filepath"

	"github.com/gorilla/mux"
)

// Middleware de CORS para adicionar os cabeçalhos CORS em todas as requisições da API
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Adicionar os cabeçalhos de CORS
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// Verificando se a requisição for um OPTIONS (verificação de pré-voo)
		if r.Method == "OPTIONS" {
			return
		}

		// Continuar com a execução da rota
		next.ServeHTTP(w, r)
	})
}

func main() {

	// Definindo o caminho do diretório dos arquivos web
	webDir := "../web/src"
	// Executando servidor web
	startWeb(webDir)

	//Executando servidor da api
	startApi()
}

func startWeb(webDir string) {
	absWebDir, err := filepath.Abs(webDir)
	if err != nil {
		fmt.Println("Erro ao obter o caminho absoluto do diretório web:", err)
		return
	}

	// Criar um servidor de arquivos estáticos para o diretório web
	fileServer := http.FileServer(http.Dir(absWebDir))

	// Rota para servir o arquivo index.html
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, filepath.Join(absWebDir, "/pages/index.html"))
	})

	// Rota para servir os arquivos estáticos (CSS, JS, etc.)
	http.Handle("/static/", http.StripPrefix("/static/", fileServer))

	fmt.Println("Acesse em: > http://localhost:3000")
	go func() {
		// Iniciar o servidor na porta 3000
		http.ListenAndServe(":3000", nil)
	}()
}

func startApi() {
	// Configurar o roteador da API usando o pacote gorilla/mux
	apiRouter := mux.NewRouter()

	// Aplicar o middleware de CORS à API
	apiRouter.Use(corsMiddleware)

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

	apiRouter.PathPrefix("/").Handler(router)
	http.ListenAndServe(":8080", apiRouter)
}
