package models

import (
	"database/sql"
	"testing"
)

type MockDB struct {}

func (mdb *MockDB) Exec(query string, args ...interface{}) (sql.Result, error){
	return nil, nil
}

func TestInit(t *testing.T) {
	err := Init("../test.Db")
	if err != nil {
		t.Fatal("can't connect to server")
	}
	err = Db.Ping()
	if err != nil {
		t.Fatal("can't ping server")
	}
}
