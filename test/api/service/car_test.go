package service_test

import (
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

	// Cria uma instância do service de carro usando o gorm mockado
	carService := service.NewCarService(gormDB)

	// Cria um objeto carro para teste
	car := model.Car{
		Brand: "Honda",
		Model: "Civic",
		Year:  2006,
	}

	t.Run("TestAddCar", func(t *testing.T) { testAddCar(t, mock, *carService, &car) })

}

func testAddCar(t *testing.T, mock sqlmock.Sqlmock, carService service.CarService, car *model.Car) {
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
