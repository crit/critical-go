package database

import (
	"github.com/jinzhu/gorm"
)

type mock struct{}

func (mock) Connection() *gorm.DB {
	return &gorm.DB{}
}
