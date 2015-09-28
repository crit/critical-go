package database

import (
	"github.com/crit/critical-go/logger"
	"strconv"
)

type Config struct {
	Driver string
	DSN    string
	Idle   string
	Max    string
	Logger logger.Logger
}

func (c Config) MaxConnections() int {
	i, err := strconv.Atoi(c.Max)

	if err != nil {
		return 1
	}

	return i
}

func (c Config) IdleConnections() int {
	i, err := strconv.Atoi(c.Idle)

	if err != nil {
		return 1
	}

	return i
}

func (c Config) LogMode() bool {
	return c.Logger != nil
}
