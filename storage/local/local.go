package local

import (
	"github.com/pmylund/go-cache"
	"time"
)

var db *cache.Cache

type Local struct{}

func New() Local {
	if db == nil {
		db = cache.New(24*time.Hour, 30*time.Second)
	}

	return Local{}
}

func (l Local) Get(key string) []byte {
	v, _ := db.Get(key)

	return v.([]byte)
}

func (l Local) Put(key string, data []byte) error {
	db.Set(key, data, 24*time.Hour)
	return nil
}

func (l Local) Delete(key string) {
	db.Delete(key)
}

func (l Local) Flush() {
	db.Flush()
}
