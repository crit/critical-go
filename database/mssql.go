package database

import (
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/jinzhu/gorm"
)

type mssqlManager struct {
	cfg Config
}

func (m mssqlManager) Connection() *gorm.DB {
	db, err := gorm.Open("mysql", m.cfg.DSN)

	if err != nil {
		if m.cfg.LogMode() {
			m.cfg.Logger.Errorf("database.MySQL Connection Error: %s -> %s", m.cfg.DSN, err.Error())
		}

		return &db
	}

	db.DB().SetMaxIdleConns(m.cfg.MaxConnections())
	db.DB().SetMaxIdleConns(m.cfg.IdleConnections())

	if m.cfg.LogMode() {
		db.LogMode(true)
		db.SetLogger(m.cfg.Logger)
	}

	return &db
}
