package test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	db "oncar-job-challenge/db"
)

func TestConnectDB(t *testing.T) {
	dbConn, err := db.ConnectDB()

	assert.Nil(t, err)
	assert.NotNil(t, dbConn)

	err = dbConn.Exec("SELECT 1").Error
	assert.Nil(t, err)

	// err = dbConn.DB().Close()
	// assert.Nil(t, err)
}
