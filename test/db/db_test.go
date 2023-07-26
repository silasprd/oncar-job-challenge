package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	model "oncar-job-challenge/api/model"
	db "oncar-job-challenge/db"

	_ "github.com/joho/godotenv/autoload"
)

func TestConnectionDB(t *testing.T) {

	//Abrir e testar conexão com o banco
	dbConn, err := db.OpenConnection("../../.env")
	assert.Nil(t, err)
	assert.NotNil(t, dbConn)

	//Executar subtestes
	t.Run("TestOpenConnectionActive", func(t *testing.T) { testOpenConnectionActive(t, dbConn) })
	t.Run("TestAutoMigrateTables", func(t *testing.T) { testAutoMigrateTables(t, dbConn) })
	t.Run("TestClosedConnection", func(t *testing.T) { testClosedConnection(t, dbConn) })

}

func testOpenConnectionActive(t *testing.T, dbConn *gorm.DB) {

	//Consulta ao banco para verificar se a conexão está ativa e funcionando
	err := dbConn.Exec("SELECT 1").Error
	assert.Nil(t, err)

}

func testAutoMigrateTables(t *testing.T, dbConn *gorm.DB) {

	//Testar auto migrate
	err := db.AutoMigrateTables(dbConn)
	assert.Nil(t, err)

	//Verificar se a tabela de carros foi migrada corretamente
	assert.True(t, dbConn.Migrator().HasTable(&model.Car{}))

}

func testClosedConnection(t *testing.T, dbConn *gorm.DB) {

	//Deferir o fechamento da conexão, para garantir que ocorra após o término da função
	defer func() {
		sqlDB, err := dbConn.DB()
		if err != nil {
			t.Errorf("Erro ao obter a conexão do banco de dados: %v", err)
			return
		}
		err = sqlDB.Close()
		assert.Nil(t, err)
	}()

	//Fechar conexão
	sqlDB, err := dbConn.DB()
	assert.Nil(t, err)
	err = sqlDB.Close()
	assert.Nil(t, err)

	//Verificar se a conexão foi fechada corretamente
	assert.True(t, sqlDB.Stats().OpenConnections == 0, "A conexão com o banco de dados não foi fechada corretamente")

}
