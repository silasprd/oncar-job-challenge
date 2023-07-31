package controller_test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"regexp"
	"strconv"
	"testing"

	"oncar-job-challenge/core/controller"
	model "oncar-job-challenge/core/model"
	"oncar-job-challenge/core/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestCarController(t *testing.T) {
	// Configura o mock do banco de dados
	db, mockDB, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	// Inicializa o gorm passando o driver do mock
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	assert.Nil(t, err)

	// Instância do service usando o banco de dados mockado
	mockCarService := service.NewCarService(gormDB)

	// Instância do controller usando o service mockado
	mockCarController := controller.NewCarController(mockCarService)

	// Modelo de carro válido para adicionar
	carToAdd := model.Car{
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2022,
		Price: 70000,
	}

	// Modelo de lista de carros válida
	listCars := []model.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Year: 2022, Price: 70000},
		{ID: 2, Brand: "Honda", Model: "Civic", Year: 2019, Price: 90000},
		{ID: 3, Brand: "Chevrolet", Model: "Onix", Year: 2013, Price: 35000},
	}

	// Modelo de carro válido para consulta
	car := model.Car{
		ID:    1,
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2022,
		Price: 70000,
	}

	// Modelo de carro válido para deleção
	carToDelete := model.Car{
		ID:    1,
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2022,
		Price: 70000,
	}

	t.Run("TestAddCar", func(t *testing.T) { testAddCar(t, mockDB, *mockCarController, carToAdd) })
	t.Run("TestGetAllCars", func(t *testing.T) { testGetAllCars(t, mockDB, *mockCarController, listCars) })
	t.Run("TestGetCar", func(t *testing.T) { testGetCar(t, mockDB, mockCarController, car) })
	t.Run("TestDeleteCar", func(t *testing.T) { testDeleteCar(t, mockDB, mockCarController, carToDelete) })

}

func testAddCar(t *testing.T, mockDB sqlmock.Sqlmock, mockController controller.CarController, car model.Car) {

	// Modifica o objeto em JSON para o corpo da requisição
	carJSON, err := json.Marshal(car)
	assert.Nil(t, err)

	// Cria uma requisição HTTP com o carro no corpo
	req, err := http.NewRequest("POST", "/api/cars", bytes.NewBuffer(carJSON))
	assert.Nil(t, err)

	// Define um Content-Type do corpo da requisição como application/json
	req.Header.Set("Content-Type", "application/json")

	// Grava a resposta HTTP em um ResponseRecorder
	rr := httptest.NewRecorder()

	// Iniciar a transação no banco
	mockDB.ExpectBegin()

	// Expectativa de execução da inserção
	mockDB.ExpectExec("INSERT INTO `cars`").
		WithArgs(car.Brand, car.Model, car.Year, car.Price).
		WillReturnResult(sqlmock.NewResult(1, 1)) // Resultado do Exec
	mockDB.ExpectCommit()

	// Chama o método AddCarHandler do controller
	mockController.AddCarHandler(rr, req)

	// Verifica o status code da resposta
	assert.Equal(t, http.StatusCreated, rr.Code, "Não foi possível adicionar o carro")

	assert.Nil(t, mockDB.ExpectationsWereMet(), "Expectativas mockadas não foram atendidas.")
}

func testGetAllCars(t *testing.T, mockDB sqlmock.Sqlmock, mockController controller.CarController, listCars []model.Car) {

	// Cria uma requisição HTTP com o carro no corpo
	req, err := http.NewRequest("GET", "/api/cars", nil)
	assert.Nil(t, err)

	// Grava a resposta HTTP em um ResponseRecorder
	rr := httptest.NewRecorder()

	// Configura o banco de dados para a consulta
	rows := sqlmock.NewRows([]string{"id", "brand", "model", "year", "price"})
	for _, car := range listCars {
		rows.AddRow(car.ID, car.Brand, car.Model, car.Year, car.Price)
	}

	// O que o teste espera que aconteça
	mockDB.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `cars`")).WillReturnRows(rows)

	rr = httptest.NewRecorder()

	// Chama o método GetAllCarsHandler do controller
	mockController.ListCarsHandler(rr, req)

	// Verifica o status code da resposta
	assert.Equal(t, http.StatusOK, rr.Code, "Não foi possível obter a lista de carros")

	// Verifica o conteúdo da resposta
	responseList := []model.Car{}
	err = json.NewDecoder(rr.Body).Decode(&responseList)
	assert.Nil(t, err, "Falha ao decodificar a resposta JSON")

	// Verifica se a lista de carros está preenchida corretamente
	assert.NotEmpty(t, responseList, "Lista de carros está vazia")
	assert.Equal(t, listCars, responseList, "Lista de carros retornada incorreta")

	// Verifica as expectativas do mock do banco de dados
	assert.Nil(t, mockDB.ExpectationsWereMet(), "Expectativas mockadas não foram atendidas.")

}

func testGetCar(t *testing.T, mockDB sqlmock.Sqlmock, mockController *controller.CarController, car model.Car) {

	// Configura uma requisição HTTP com o ID do carro no caminho
	router := mux.NewRouter()
	router.HandleFunc("/api/cars/{id}", mockController.GetCarHandler).Methods("GET")
	url := "/api/cars/" + strconv.Itoa(int(car.ID))
	req, err := http.NewRequest("GET", url, nil)
	assert.Nil(t, err)

	// Grava a resposta HTTP em um ResponseRecorder
	rr := httptest.NewRecorder()

	// Configura o banco de dados para a consulta
	columns := []string{"id", "brand", "model", "year", "price"}
	rows := sqlmock.NewRows(columns).AddRow(car.ID, car.Brand, car.Model, car.Year, car.Price)
	query := "SELECT * FROM `cars` WHERE `cars`.`id` = ? ORDER BY `cars`.`id` LIMIT 1"
	mockDB.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(car.ID).WillReturnRows(rows)

	// Chama a função configurada no router
	router.ServeHTTP(rr, req)

	// Verifica o status code da resposta (esperado: 200 OK)
	assert.Equal(t, http.StatusOK, rr.Code, "Carro retornado incorretamente!")

	// Verifica o conteúdo da resposta (esperado: carro serializado)
	var responseCar model.Car
	err = json.NewDecoder(rr.Body).Decode(&responseCar)
	assert.Nil(t, err, "Erro ao decodificar a resposta JSON.")

	// Verifica se o carro retornado não é nulo
	assert.NotNil(t, responseCar, "Carro retornado é nulo.")

	// Verifica o conteúdo do carro retornado
	assert.Equal(t, car.ID, responseCar.ID, "ID do carro retornado incorreto.")

	// Verifica as expectativas do mockadas
	assert.Nil(t, mockDB.ExpectationsWereMet(), "Expectativas mockadas não foram atendidas.")

}

func testDeleteCar(t *testing.T, mockDB sqlmock.Sqlmock, mockController *controller.CarController, carToDelete model.Car) {

	router := mux.NewRouter()
	router.HandleFunc("/api/cars/{id}", mockController.DeleteCarHandler).Methods("DELETE")
	url := "/api/cars/" + strconv.Itoa(int(carToDelete.ID))
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("Erro ao criar a requisição HTTP: %v", err)
	}

	// Grava a resposta HTTP em um ResponseRecorder
	rr := httptest.NewRecorder()

	// Configura o banco de dados para a deleção do carro
	query := "DELETE FROM `cars` WHERE `cars`.`id` = ?"
	mockDB.ExpectBegin()
	mockDB.ExpectExec(regexp.QuoteMeta(query)).WithArgs(carToDelete.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mockDB.ExpectCommit()

	router.ServeHTTP(rr, req)

	// Verifica o retorno da função delete
	assert.Equal(t, http.StatusOK, rr.Code, "Carro não foi deletado com sucesso!")

	// Verifica se a propriedade ID do carro existe
	assert.NotEmpty(t, carToDelete.ID, "ID do carro não pode ser vazio!")

	// Verifica as expectativas do mockadas
	assert.Nil(t, mockDB.ExpectationsWereMet(), "Expectativas mockadas não foram atendidas.")

}
