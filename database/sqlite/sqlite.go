package sqlite

import (
	"github.com/crit/critical-go/logger"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"sync"
)

var conn *gorm.DB
var lock sync.Mutex

type Client struct {
	dsn string
	log logger.Logger
}

func New(dsn string, log logger.Logger) Client {
	return Client{dsn, log}
}

func (c Client) Connection() *gorm.DB {
	if conn != nil {
		return conn
	}

	lock.Lock()
	defer lock.Unlock()

	db, err := gorm.Open("sqlite3", c.dsn)

	if err != nil {
		if c.log != nil {
			c.log.Errorf("database.MySQL Connection Error: %s -> %s", c.dsn, err.Error())
		}

		conn = &db

		return conn
	}

	if c.log != nil {
		db.LogMode(true)
		db.SetLogger(c.log)
	}

	conn = &db

	return conn
}
