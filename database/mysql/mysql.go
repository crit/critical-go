package mysql

import (
	"github.com/crit/critical-go/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"sync"
)

var conn *gorm.DB
var lock sync.Mutex

type Client struct {
	dsn  string
	max  int
	idle int
	log  logger.Logger
}

func New(dsn string, max, idle int, log logger.Logger) Client {
	return Client{dsn, max, idle, log}
}

func (c Client) Connection() *gorm.DB {
	if conn != nil {
		return conn
	}

	lock.Lock()
	defer lock.Unlock()

	// when handed the lock, conn may be initialized
	if conn != nil {
		return conn
	}

	db, err := gorm.Open("mysql", c.dsn)

	if err != nil {
		if c.log != nil {
			c.log.Errorf("database.MySQL Connection Error: %s -> %s", c.dsn, err.Error())
		}

		conn = &db

		return conn
	}

	db.DB().SetMaxIdleConns(c.idle)
	db.DB().SetMaxOpenConns(c.max)

	if c.log != nil {
		db.LogMode(true)
		db.SetLogger(c.log)
	}

	conn = &db

	return conn
}
