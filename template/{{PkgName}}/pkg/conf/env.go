package conf

import (
	environ "github.com/katsew/go-getenv"
)

var Env string

const (
	EnvKey         = "ENVIRONMENT"
	EnvLocalhost   = "localhost"
	EnvDevelopment = "development"
	EnvStaging     = "staging"
	EnvProduction  = "production"
)

func init() {
	Env = environ.GetEnv(EnvKey, EnvLocalhost).String()
}