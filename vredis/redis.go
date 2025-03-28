package vredis

import (
	"context"

	"github.com/redis/go-redis/v9"

	"github.com/sunshinexcode/gotk/vmap"
	"github.com/sunshinexcode/gotk/vreflect"
)

type Redis struct {
	Client  redis.UniversalClient
	Options *redis.UniversalOptions
}

var (
	defaultOptions = map[string]any{
		"Addrs":    []string{"localhost:6379"},
		"DB":       0,
		"Password": "",
		"PoolSize": 100,
	}
)

// New create new redis
func New(options map[string]any) (redisS *Redis, err error) {
	redisS = &Redis{}
	err = redisS.SetConfig(options)

	return
}

// C get client
func (redisS *Redis) C() redis.UniversalClient {
	return redisS.Client
}

func (redisS *Redis) Ping() (string, error) {
	return redisS.C().Ping(context.TODO()).Result()
}

// SetConfig set config
func (redisS *Redis) SetConfig(options map[string]any) (err error) {
	redisS.Options = &redis.UniversalOptions{}
	if err = vreflect.SetAttrs(redisS.Options, vmap.Merge(defaultOptions, options)); err != nil {
		return
	}

	redisS.Client = redis.NewUniversalClient(redisS.Options)
	return
}
