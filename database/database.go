package database

import (
	"github.com/jinzhu/gorm"
)

// Database is a proxy for an instantiated gorm object for a specific
// backing database driver.
type Database interface {
	// Connection returns a connected gorm DB object.
	Connection() *gorm.DB
}

func Mock() Database {
	return mock{}
}

func New(cfg Config) Database {
	switch cfg.Driver {
	case "mssql":
		return mssqlManager{cfg: cfg}
	case "mysql":
		return mysqlManager{cfg: cfg}
	default:
		return mock{}
	}
}
