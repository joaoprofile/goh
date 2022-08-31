package cachedb

import (
	"context"
	"log"
	"sync"
	"time"

	"github.com/joaocprofile/goh/core"
	env "github.com/joaocprofile/goh/environment"

	"github.com/go-redis/redis/v8"
)

var lock = &sync.Mutex{}

type Redis struct {
	client *redis.Client
}

func NewConnection() *Redis {
	lock.Lock()
	connection, _ := createConnection()
	defer lock.Unlock()
	return connection
}
func (r *Redis) Close() error {
	if err := r.client.Close(); err != nil {
		return err
	}
	return nil
}

func createConnection() (*Redis, error) {
	client := redis.NewClient(&redis.Options{
		Addr:        env.Get().CacheDB.ConnectionString,
		DB:          0,
		Password:    env.Get().CacheDB.Password,
		DialTimeout: 100 * time.Millisecond,
		ReadTimeout: 100 * time.Millisecond,
	})

	if _, err := client.Ping(context.Background()).Result(); err != nil {
		log.Fatal(core.Red("Error connecting to Redis: " + err.Error()))
		return nil, err
	}

	return &Redis{
		client: client,
	}, nil
}

func (r *Redis) Get(ctx context.Context, key string) ([]byte, error) {
	cmd := r.client.Get(ctx, key)
	cmdb, err := cmd.Bytes()
	if err != nil {
		return nil, err
	}
	return cmdb, nil
}

func (r *Redis) Set(ctx context.Context, key string, obj interface{}) error {
	obj, err := core.ToJson(obj)
	if err != nil {
		return err
	}

	return r.client.Set(
		ctx,
		key,
		obj,
		env.Get().CacheDB.ExpirationCache,
	).Err()
}

func (r *Redis) Del(ctx context.Context, key string) error {
	cmd := r.client.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}
