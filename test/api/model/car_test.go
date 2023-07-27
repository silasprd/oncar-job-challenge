package model_test

import (
	model "oncar-job-challenge/api/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarModel(t *testing.T) {

	validCar := model.Car{
		ID:    1,
		Brand: "Honda",
		Model: "Civic",
		Year:  2006,
	}

	invalidCar := model.Car{
		ID:    0,
		Brand: "",
		Model: "",
		Year:  0,
	}

	t.Run("TestValidCarFields", func(t *testing.T) { testValidCarFields(t, validCar) })
	t.Run("TestInvalidCarFields", func(t *testing.T) { testInvalidCarFields(t, invalidCar) })
	t.Run("TestValidModel", func(t *testing.T) { testValidModel(t, validCar) })

}

func testValidCarFields(t *testing.T, validCar model.Car) {

	assert.NotZero(t, validCar.ID, "ID deve ser diferente de 0.")
	assert.NotEmpty(t, validCar.Brand, "Marca do carro não pode ser vazio.")
	assert.NotEmpty(t, validCar.Model, "Modelo do carro não pode ser vazio.")
	assert.NotEmpty(t, validCar.Brand, "Ano do carro não pode ser vazia")
	assert.NotZero(t, validCar.ID, "Ano deve ser diferente de 0.")

}

func testInvalidCarFields(t *testing.T, invalidCar model.Car) {

	assert.Zero(t, invalidCar.ID, "ID deve ser diferente de 0.")
	assert.Empty(t, invalidCar.Brand, "Marca do carro não pode ser vazio.")
	assert.Empty(t, invalidCar.Model, "Modelo do carro não pode ser vazio.")
	assert.Zero(t, invalidCar.Year, "Ano deve ser diferente de 0.")

}

func testValidModel(t *testing.T, validCar model.Car) {

	assert.IsType(t, uint(0), validCar.ID, "ID deve ser do tipo uint")
	assert.IsType(t, "string", validCar.Brand, "ID deve ser do tipo string")
	assert.IsType(t, "string", validCar.Model, "ID deve ser do tipo string")
	assert.IsType(t, int(0), validCar.Year, "ID deve ser do tipo int")
	assert.IsType(t, float64(0), validCar.Price, "ID deve ser do tipo float64")

}
