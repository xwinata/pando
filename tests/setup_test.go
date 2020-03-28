package tests

import (
	"os"
	"pando/db"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jinzhu/gorm"
)

var (
	mock      sqlmock.Sqlmock
	tempDB    *gorm.DB
	testingDB *gorm.DB
)

// Setup mock environtment
func Setup() {
	sqldb, smock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	gormDBmock, err := gorm.Open(os.Getenv("PANDO_DB_ADAPTER"), sqldb)
	if err != nil {
		panic(err)
	}
	gormDBmock.LogMode(true)
	tempDB = db.DB
	testingDB = gormDBmock
	mock = smock
}

// BeginTest mark testing start by changing db to testing db var temporarily
func BeginTest() {
	db.DB = testingDB
}

// FinishTest finishes test by returning db setting and closing testing db
func FinishTest() {
	db.DB = tempDB
	testingDB.Close()
}
