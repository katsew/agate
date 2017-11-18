package core

import (
	"fmt"

	environ "github.com/katsew/go-getenv"
	"github.com/rainycape/memcache"
)

const (
	MemcachedHostKey     = "MEMCACHED_HOST"
	MemcachedHostDefault = "127.0.0.1"
	MemcachedPortKey     = "MEMCACHED_PORT"
	MemcachedPortDefault = "21211"
)

var memcachedClient *memcache.Client

func init() {

	aServer := fmt.Sprintf(
		"%s:%s",
		environ.GetEnv(MemcachedHostKey, MemcachedHostDefault),
		environ.GetEnv(MemcachedPortKey, MemcachedPortDefault),
	)
	client, err := memcache.New(aServer)
	if err != nil {
		panic(err)
	}
	memcachedClient = client
}

func GetMemcachedInstance() *memcache.Client {
	return memcachedClient
}
