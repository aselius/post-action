package main

import (
	"os"
	"fmt"
	_ "net/http"
	_ "net/http/httptest"
	"testing"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	_"github.com/stretchr/testify/assert"
)

func setupTestDB() (*gorm.DB) {
	var dbConnInfo = fmt.Sprintf("host=%s port=%d user=%s "+
	"password=%s dbname=%s sslmode=disable",
	"localhost", 5432, "postgres", "postgres", "postgres")

	gormDB, err := gorm.Open("postgres", dbConnInfo)
	if err != nil {
		panic(err)
	}
	
	return gormDB
}

var (
	dB = setupTestDB()
)

func TestMain(m *testing.M) {
	err := dB.Exec("create table if not exists foobar (id serial primary key, name varchar(128) not null, points integer)").Error
	if err != nil {
		panic(err)
	}

	os.Exit(m.Run())
}