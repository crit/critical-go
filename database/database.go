package database

import (
	"github.com/crit/critical-go/database/mssql"
	"github.com/crit/critical-go/database/mysql"
	"github.com/crit/critical-go/database/sqlite"
	"github.com/jinzhu/gorm"
)

// Database is a proxy for an instantiated gorm object for a specific
// backing database driver.
type Database interface {
	// Connection returns a connected gorm DB object.
	Connection() *gorm.DB
}

// New returns an unconnected Database object that is setup to use
// a specific database driver.
func New(cfg Config) Database {
	switch cfg.Driver {
	case "mssql":
		return mssql.New(cfg.DSN, cfg.MaxConnections(), cfg.IdleConnections(), cfg.Logger)
	case "mysql":
		return mysql.New(cfg.DSN, cfg.MaxConnections(), cfg.IdleConnections(), cfg.Logger)
	case "sqlite":
		return sqlite.New(cfg.DSN, cfg.Logger)
	default:
		return sqlite.New("/tmp/gorm.db?loc=auto", cfg.Logger)
	}
}
