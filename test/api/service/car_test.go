package service_test

import (
	"regexp"
	"testing"

	model "oncar-job-challenge/api/model"
	"oncar-job-challenge/api/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func TestCarService(t *testing.T) {
	// Configura o mock do banco de dados
	db, mock, err := sqlmock.New()
	assert.Nil(t, err)
	defer db.Close()

	// Inicializa o gorm passando o driver do mock
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn:                      db,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{})
	assert.Nil(t, err)

	// Instância do service de carro usando o gorm mockado
	carService := service.NewCarService(gormDB)

	// Objeto carro para teste AddCar
	carToAdd := model.Car{
		Brand: "Honda",
		Model: "Civic",
		Year:  2019,
	}

	//Objeto carro para teste Get
	car := model.Car{
		ID:    1,
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
		Price: 70000,
	}

	// Lista de carros para teste
	validListCars := []model.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Year: 2020},
		{ID: 2, Brand: "Honda", Model: "Civic", Year: 2019},
	}

	t.Run("TestAddCar", func(t *testing.T) { testAddCar(t, mock, *carService, carToAdd) })
	t.Run("TestGetAllCars", func(t *testing.T) { testGetAllCars(t, mock, *carService, validListCars) })
	t.Run("TestGetCar", func(t *testing.T) { testGetCar(t, mock, *carService, car) })

}

func testAddCar(t *testing.T, mock sqlmock.Sqlmock, carService service.CarService, car model.Car) {
	// O que o teste espera
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `cars`").WithArgs(car.Brand, car.Model, car.Year, car.Price).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Executar o método AddCar
	err := carService.AddCar(car)
	assert.Nil(t, err)

	// Verificar se obteve o retorno esperado
	assert.Nil(t, mock.ExpectationsWereMet())
}

func testGetAllCars(t *testing.T, mock sqlmock.Sqlmock, carService service.CarService, listCars []model.Car) {

	// Configura o que o teste espera da consulta a ser realizada
	rows := sqlmock.NewRows([]string{"id", "brand", "model", "year", "price"})
	rows.AddRow(3, "Chevrolet", "Onix", 2014, 38000)
	rows.AddRow(4, "Fiat", "Mobi", 2014, 31000)

	mock.ExpectQuery("SELECT (.+) FROM `cars`").WillReturnRows(rows)

	// Executa o método GetAllCars
	resultCars, err := carService.GetAllCars()
	assert.Nil(t, err)

	// Verifica se retornou o que o mock estava esperando
	assert.Nil(t, mock.ExpectationsWereMet())

	// Testa se uma das listas não estão vazias
	assert.NotEmpty(t, listCars, resultCars)

	// Faz uma verificação campo a campo para valores que não podem ser nulos
	for i, car := range listCars {
		assert.NotEmpty(t, car.ID, resultCars[i].ID)
		assert.NotEmpty(t, car.Brand, resultCars[i].Brand)
		assert.NotEmpty(t, car.Model, resultCars[i].Model)
		assert.NotEmpty(t, car.Year, resultCars[i].Year)
	}
}

func testGetCar(t *testing.T, mock sqlmock.Sqlmock, carService service.CarService, car model.Car) {

	// Configura o que o teste espera da consulta a ser realizada
	rows := sqlmock.NewRows([]string{"id", "brand", "model", "year", "price"}).
		AddRow(1, "Toyota", "Corolla", 2020, 70000)

	query := "SELECT * FROM `cars` WHERE `cars`.`id` = ? ORDER BY `cars`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(car.ID).WillReturnRows(rows)

	// Executa o método GetCarByID
	resultCar, err := carService.GetCarByID(car.ID)
	assert.Nil(t, err)

	// Verifica se retornou o que o mock estava esperando
	assert.Nil(t, mock.ExpectationsWereMet())

	// Verifica se retornou o objeto corretamente
	assert.Equal(t, car.ID, resultCar.ID)
	assert.Equal(t, car.Brand, resultCar.Brand)
	assert.Equal(t, car.Model, resultCar.Model)
	assert.Equal(t, car.Year, resultCar.Year)
	assert.Equal(t, car.Price, resultCar.Price)

}
