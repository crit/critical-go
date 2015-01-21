package storer

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"log"
)

var DB gorm.DB

var storer StoreManager = StoreManager{}

const (
	MYSQL = "mysql"
)

type StoreManager struct {
	DSN     string
	Logging bool
	Start   bool
	Error   string
}

type Options struct {
	Engine  string
	DSN     string
	Logging bool
}

func InitStore(options Options) {
	var err error

	DB, err = gorm.Open(options.Engine, options.Engine)

	if err != nil {
		log.Println(err.Error())
		storer.Error = err.Error()
		return
	}

	DB.LogMode(options.Logging)
	DB.DB().SetMaxIdleConns(10)
	DB.DB().SetMaxOpenConns(100)

	if err := DB.DB().Ping(); err != nil {
		log.Println(err.Error())
		storer.Error = err.Error()
		return
	}

	storer.Start = true
}

func Error() string {
	return storer.Error
}
