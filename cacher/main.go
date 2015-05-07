package cacher

import (
	"log"
	"strings"
	"time"

	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pmylund/go-cache"
)

var mc *memcache.Client
var lc *cache.Cache

var cacher CacheManager = CacheManager{}

const (
	MEMCACHE   = "mc"
	LOCALCACHE = "lc"
)

type CacheManager struct {
	Hosts  string
	Engine string
}

type Options struct {
	Hosts  string
	Engine string
}

func InitCache(options Options) {
	cacher.Engine = options.Engine
	cacher.Hosts = options.Hosts

	if cacher.Hosts == "" {
		cacher.Engine = LOCALCACHE
		logThis("Overriding Engine with LOCALCACHE since Hosts is empty")
	}

	switch cacher.Engine {
	case MEMCACHE:
		hosts := strings.Split(cacher.Hosts, ",")
		mc = memcache.New(hosts...)
	case LOCALCACHE:
		lc = cache.New(-1, 24*time.Hour)
	}
}

func Set(key string, value []byte) {
	switch cacher.Engine {
	case LOCALCACHE:
		lcSet(key, value)
	case MEMCACHE:
		mcSet(key, value)
	}
}

func Get(key string) []byte {
	switch cacher.Engine {
	case MEMCACHE:
		return mcGet(key)
	case LOCALCACHE:
		return lcGet(key)
	}

	return []byte{}
}

func Delete(key string) {
	switch cacher.Engine {
	case MEMCACHE:
		mc.Delete(key)
	case LOCALCACHE:
		lc.Delete(key)
	}
}

func mcSet(key string, value []byte) {
	if err := mc.Add(&memcache.Item{Key: key, Value: value}); err != nil {
		logThis(err.Error())
	}
}

func mcGet(key string) []byte {
	item, err := mc.Get(key)

	if err != nil {
		logThis(err.Error())
		return []byte{}
	}

	return []byte(item.Value)
}

func lcSet(key string, value []byte) {
	if err := lc.Add(key, value, 24*time.Hour); err != nil {
		logThis(err.Error())
	}
}

func lcGet(key string) []byte {
	item, found := lc.Get(key)

	if !found {
		logThis("LOCALCACHE miss on " + key)
		return []byte{}
	}

	return item.([]byte)
}

func logThis(msg string) {
	log.Printf("[CACHER] %s", msg)
}
