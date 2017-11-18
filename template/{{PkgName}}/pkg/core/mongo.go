{{- $gopath := env "GOPATH" -}}
{{- $pwd := env "PWD" -}}
{{- $relPath := trimPrefix $pwd $gopath -}}
{{- $path := trimPrefix $relPath "/src/" -}}
{{- $pkgpath := printf "%s/%s/pkg" $path PkgName -}}
package core

import (
	"time"

	environ "github.com/katsew/go-getenv"
	"github.com/labstack/gommon/log"
	"gopkg.in/mgo.v2"

	"{{$pkgpath}}/conf"
)

const (
	MongoUrlEnvKey      = "MONGO_URL"
	MongoUrlDefault     = "127.0.0.1:37017"
	MongoTimeoutEnvKey  = "MONGO_TIMEOUT"
	MongoTimeoutDefault = "10s"
)

var m *mgo.Session

func init() {

	var err error
	configs := []string{environ.GetEnv(MongoUrlEnvKey, MongoUrlDefault).String()}

	// @todo Fixme!
	log.Printf("Config: %s", configs[0])
	timeout, err := time.ParseDuration(environ.GetEnv(MongoTimeoutEnvKey, MongoTimeoutDefault).String())
	if err != nil {
		log.Panicf("Could not parse timeout value from env, Reason: %s", err.Error())
	}
	info := mgo.DialInfo{
		Addrs:   configs,
		Timeout: timeout,
	}

	m, err = mgo.DialWithInfo(&info)
	if err != nil {
		log.Panicf("Could not initialize DBSession, Reason: %s", err.Error())
	}
	if conf.Env != conf.EnvProduction {
		mgo.SetDebug(true)
	}

}

func GetMongoInstance() *mgo.Session {
	return m
}

