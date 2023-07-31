package service_test

import (
	"regexp"
	"testing"

	model "oncar-job-challenge/core/model"
	"oncar-job-challenge/core/service"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GormDB interface {
	Preload(column string, conditions ...interface{}) *gorm.DB
	Find(dest interface{}, conds ...interface{}) *gorm.DB
	Error() error
}

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

	// Lista de carros para teste
	listCars := []model.Car{
		{ID: 1, Brand: "Toyota", Model: "Corolla", Year: 2020},
		{ID: 2, Brand: "Honda", Model: "Civic", Year: 2019},
	}

	//Objeto carro para teste Get
	car := model.Car{
		ID:    1,
		Brand: "Toyota",
		Model: "Corolla",
		Year:  2020,
		Price: 70000,
	}

	// Objeto carro para teste AddCar
	carToAdd := model.Car{
		Brand: "Honda",
		Model: "Civic",
		Year:  2019,
	}

	// Carro para teste delete
	carToDelete := model.Car{
		ID:    1,
		Brand: "Fiat",
		Model: "Toro",
		Year:  2017,
		Price: 70000,
	}

	t.Run("TestAddCar", func(t *testing.T) { testAddCar(t, mock, *carService, carToAdd) })
	t.Run("TestGetAllCars", func(t *testing.T) { testGetAllCars(t, mock, *carService, listCars) })
	t.Run("TestGetCar", func(t *testing.T) { testGetCar(t, mock, *carService, car) })
	t.Run("TestDeleteCar", func(t *testing.T) { testDeleteCar(t, mock, *carService, carToDelete) })

}

func testAddCar(t *testing.T, mock sqlmock.Sqlmock, carService service.CarService, car model.Car) {

	// O que o teste espera
	mock.ExpectBegin()
	mock.ExpectExec("INSERT INTO `cars`").WithArgs(car.Brand, car.Model, car.Year, car.Price).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Executar o método AddCar
	err := carService.AddCar(car)
	assert.Nil(t, err, "Falha ao adicionar carro!")

	// Verificar se obteve o retorno esperado
	assert.Nil(t, mock.ExpectationsWereMet(), "Expectativas mockadas não foram atendidas.")

}

func testGetAllCars(t *testing.T, mock sqlmock.Sqlmock, carService service.CarService, listCars []model.Car) {

	// Configura o que o teste espera da consulta a ser realizada
	rows := sqlmock.NewRows([]string{"id", "brand", "model", "year", "price"})
	rows.AddRow(3, "Chevrolet", "Onix", 2014, 38000)
	rows.AddRow(4, "Fiat", "Mobi", 2014, 31000)

	contacts := sqlmock.NewRows([]string{"id", "name", "email", "phone", "car_id"})
	contacts.AddRow(1, "Silas", "silas.prado@gmail.com", "12998776655", 3)
	contacts.AddRow(2, "Silas", "silas.prado@gmail.com", "12998776655", 4)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `cars`")).WillReturnRows(rows)
	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `contacts` WHERE `contacts`.`car_id` IN (?,?)")).WithArgs(3, 4).WillReturnRows(contacts)

	// Executa o método GetAllCars
	resultCars, err := carService.GetAllCars()
	assert.Nil(t, err, "Falha ao obter a lista de carros!")

	// Verifica se retornou o que o mock estava esperando
	assert.Nil(t, mock.ExpectationsWereMet(), "Expectativas mockadas não foram atendidas.")

	// Testa se uma das listas não estão vazias
	assert.NotEmpty(t, listCars, resultCars, "Lista vazia!")

	// Faz uma verificação campo a campo para valores obrigatórios
	for i, car := range listCars {
		assert.NotEmpty(t, car.ID, resultCars[i].ID, "ID do carro é obrigatório.")
		assert.NotEmpty(t, car.Brand, resultCars[i].Brand, "Marca do carro é obrigatória")
		assert.NotEmpty(t, car.Model, resultCars[i].Model, "Modelo do carro é obrigatório")
		assert.NotEmpty(t, car.Year, resultCars[i].Year, "Ano do carro é obrigatório")
	}

}

func testGetCar(t *testing.T, mock sqlmock.Sqlmock, carService service.CarService, car model.Car) {

	// Configura o que o teste espera da consulta a ser realizada
	rows := sqlmock.NewRows([]string{"id", "brand", "model", "year", "price"}).
		AddRow(car.ID, car.Brand, car.Model, car.Year, car.Price)

	contacts := sqlmock.NewRows([]string{"id", "name", "email", "phone", "car_id"}).
		AddRow(1, "Silas", "silas.prado@gmail.com", "12996523398", 1)

	query := "SELECT * FROM `cars` WHERE `cars`.`id` = ? ORDER BY `cars`.`id` LIMIT 1"
	mock.ExpectQuery(regexp.QuoteMeta(query)).WithArgs(car.ID).WillReturnRows(rows)

	mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM `contacts` WHERE `contacts`.`car_id` = ?")).WithArgs(1).WillReturnRows(contacts)

	// Executa o método GetCarByID
	resultCar, err := carService.GetCarByID(car.ID)
	assert.Nil(t, err, "Não foi possível encontrar o carro!")

	// Verifica se retornou o que o mock estava esperando
	assert.Nil(t, mock.ExpectationsWereMet(), "Expectativas mockadas não foram atendidas.")

	// Verifica se retornou o objeto corretamente
	assert.Equal(t, car.ID, resultCar.ID)
	assert.Equal(t, car.Brand, resultCar.Brand)
	assert.Equal(t, car.Model, resultCar.Model)
	assert.Equal(t, car.Year, resultCar.Year)
	assert.Equal(t, car.Price, resultCar.Price)

}

func testDeleteCar(t *testing.T, mock sqlmock.Sqlmock, carService service.CarService, carToDelete model.Car) {

	// Configura a query que o teste espera que seja executada ao chamar a função
	query := "DELETE FROM `cars` WHERE `cars`.`id` = ?"
	mock.ExpectBegin()
	mock.ExpectExec(regexp.QuoteMeta(query)).WithArgs(carToDelete.ID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	// Executa a função DeleteCar
	err := carService.DeleteCar(carToDelete.ID)
	assert.Nil(t, err, "Não foi possível excluir o carro!")

	// Verifica se a propriedade ID do carro existe
	assert.NotEmpty(t, carToDelete.ID, "ID do carro não pode ser vazio!")

	// Verifica se o retorno do mock está correto
	assert.Nil(t, mock.ExpectationsWereMet(), "Expectativas mockadas não foram atendidas")

}
