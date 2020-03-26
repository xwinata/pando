package db

import (
	"fmt"
	"os"
	"pando/models"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"

	// import db drivers
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

// DB database instance
var DB *gorm.DB

func init() {
	DBinit()
	Migrate()
}

// DBinit initiates database connection
func DBinit() {
	var connectionString string
	dbAdapter := os.Getenv("PANDO_DB_ADAPTER")
	switch dbAdapter {
	default:
		panic(fmt.Sprintf("invalid adapter %v", os.Getenv("PANDO_DB_ADAPTER")))
	case "mysql":
		connectionString = fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
			os.Getenv("PANDO_DB_USERNAME"),
			os.Getenv("PANDO_DB_PASSWORD"),
			os.Getenv("PANDO_DB_HOST"),
			os.Getenv("PANDO_DB_PORT"),
			os.Getenv("PANDO_DB_TABLE"))
		break
	case "postgres":
		connectionString = fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?sslmode=%s",
			os.Getenv("PANDO_DB_USERNAME"),
			os.Getenv("PANDO_DB_PASSWORD"),
			os.Getenv("PANDO_DB_HOST"),
			os.Getenv("PANDO_DB_PORT"),
			os.Getenv("PANDO_DB_TABLE"),
			os.Getenv("PANDO_DB_SSL"))
		break
	}

	db, err := gorm.Open(dbAdapter, connectionString)
	if err != nil {
		panic(err)
	}
	if err = db.DB().Ping(); err != nil {
		panic(err)
	}

	if logmode := os.Getenv("PANDO_DB_LOGMODE"); logmode == "true" || logmode == "false" {
		switch logmode {
		case "false":
			db.LogMode(false)
			break
		case "true":
			db.LogMode(true)
			break
		}
	}

	db.Exec(fmt.Sprintf("SET TIMEZONE TO '%s'", os.Getenv("PANDO_TIMEZONE")))
	if maxLifeTime, err := strconv.ParseInt(os.Getenv("PANDO_DB_MAXLIFETIME"), 10, 64); err != nil {
		db.DB().SetConnMaxLifetime(time.Second * time.Duration(maxLifeTime))
	}
	if maxIdleConnection, err := strconv.Atoi(os.Getenv("PANDO_DB_MAXIDLECONNECTION")); err != nil {
		db.DB().SetMaxIdleConns(maxIdleConnection)
	}
	if maxOpenConnection, err := strconv.Atoi(os.Getenv("PANDO_DB_MAXOPENCONNECTION")); err != nil {
		db.DB().SetMaxOpenConns(maxOpenConnection)
	}

	DB = db
}

// Migrate updates database structures
func Migrate() {
	err := DB.AutoMigrate(
		&models.Client{},
		&models.User{},
		&models.Role{},
		&models.Permission{},
		&models.UserRole{},
		&models.RolePermission{},
	).Error
	if err != nil {
		panic(err)
	}
}

func transaction(fn func(tx *gorm.DB) error) error {
	// Note the use of tx as the database handle once you are within a transaction
	tx := DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return err
	}

	if err := fn(tx); err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

// Create function
func Create(i interface{}) error {
	return transaction(func(tx *gorm.DB) (err error) {
		return tx.Create(i).Error
	})
}
