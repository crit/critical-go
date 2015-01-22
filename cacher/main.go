package cacher

import (
	"github.com/bradfitz/gomemcache/memcache"
	"github.com/pmylund/go-cache"
	"log"
	"strings"
	"time"
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
		log.Println("-- Overriding Engine with LOCALCACHE since Hosts is empty --")
	}

	switch cacher.Engine {
	case MEMCACHE:
		hosts := strings.Split(cacher.Hosts, ",")
		mc = memcache.New(hosts...)
	case LOCALCACHE:
		lc = cache.New(-1, 24*time.Hour)
	}
}

func Set(key, value string) {
	switch cacher.Engine {
	case LOCALCACHE:
		lcSet(key, value)
	case MEMCACHE:
		mcSet(key, value)
	}
}

func Get(key string) string {
	switch cacher.Engine {
	case MEMCACHE:
		return mcGet(key)
	case LOCALCACHE:
		return lcGet(key)
	}

	return ""
}

func Delete(key string) {
	switch cacher.Engine {
	case MEMCACHE:
		mc.Delete(key)
	case LOCALCACHE:
		lc.Delete(key)
	}
}

func mcSet(key, value string) {
	if err := mc.Add(&memcache.Item{Key: key, Value: []byte(value)}); err != nil {
		log.Println(err.Error())
	}
}

func mcGet(key string) string {
	item, err := mc.Get(key)

	if err != nil {
		log.Println(err.Error())
		return ""
	}

	return string(item.Value)
}

func lcSet(key, value string) {
	if err := lc.Add(key, value, 24*time.Hour); err != nil {
		log.Println(err.Error())
	}
}

func lcGet(key string) string {
	item, found := lc.Get(key)

	if !found {
		log.Println("LC Cache miss")
		return ""
	}

	return item.(string)
}
