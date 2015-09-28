package redis

import (
	"time"

	"github.com/garyburd/redigo/redis"
)

type Redis struct {
	pool *redis.Pool
}

func New(host, port string) Redis {
	dsn := host + ":" + port

	return Redis{
		pool: &redis.Pool{
			MaxIdle:     3,
			IdleTimeout: 240 * time.Second,
			Dial: func() (redis.Conn, error) {
				c, err := redis.Dial("tcp", dsn)
				if err != nil {
					return nil, err
				}

				return c, nil
			},
		},
	}
}

func (r Redis) Get(key string) []byte {
	client := r.pool.Get()
	defer client.Close()

	data, _ := redis.Bytes(client.Do("GET", key))

	return data
}

func (r Redis) Put(key string, data []byte) error {
	client := r.pool.Get()
	defer client.Close()

	_, err := client.Do("SET", key, data)

	return err
}

func (r Redis) Delete(key string) {
	client := r.pool.Get()
	defer client.Close()

	client.Do("DEL", key)
}

func (r Redis) Flush() {
	client := r.pool.Get()
	defer client.Close()

	client.Do("FLUSHALL")
}
