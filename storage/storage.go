package storage

import (
	"github.com/crit/critical-go/storage/folder"
	"github.com/crit/critical-go/storage/local"
	"github.com/crit/critical-go/storage/redis"
	"github.com/crit/critical-go/storage/s3"
)

// Storage is a simple proxy for many different back-end systems. Currently supporting
// Local (in-memory), Folder (system folder), S3 (Amazon's S3 Storage), and Redis.
type Storage interface {
	// Get returns the stored data by key.
	Get(key string) []byte

	// Put overwrites (or creates) the key with the passed in data.
	Put(key string, data []byte) error

	// Delete removes data by key. Does nothing if key does not exist.
	Delete(key string)

	// Flush removes all keys. Does nothing on S3 or Folder.
	Flush()
}

// Local returns an instance of Storage backed only by the
// running binary's memory.
func Local() Storage {
	return local.New()
}

// Folder returns an instance of Storage backed by the local
// file system. Each key is a file at the provided path.
func Folder(path string) Storage {
	if path == "" {
		return local.New()
	}

	return folder.New(path)
}

// S3 returns an instance of Storage backed by a specified S3 bucket. Each key is
// a file. `content` should be the something like "application/json; charset=utf-8" for whatever
// http representation of the data that is being stored.
func S3(secret, access, bucket, region, content string) Storage {
	return s3.New(secret, access, bucket, region, content)
}

// Redis returns an instance of Storage backed by a Redis service.
func Redis(host, port string) Storage {
	if host == "" || port == "" {
		return local.New()
	}

	return redis.New(host, port)
}
