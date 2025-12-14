package database

import (
	"example/shop-progect/config"
	"net"
	"strconv"

	"github.com/boj/redistore"
)

func NewRedisStore(sessionKey []byte) *redistore.RediStore {
	addr := net.JoinHostPort(config.Cfg.RedisHost, strconv.Itoa(config.Cfg.RedisPort))

	store, err := redistore.NewRediStore(10, "tcp", addr, config.Cfg.RedisUser, config.Cfg.RedisPassword, sessionKey)

	if err != nil {
		panic(err)
	}

	store.SetMaxAge(86400 * 7)
	store.SetSerializer(redistore.JSONSerializer{})

	return store
}
