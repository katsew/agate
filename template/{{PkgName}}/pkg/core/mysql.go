package core

import (
	"fmt"

	environ "github.com/katsew/go-getenv"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gocraft/dbr"
)

const (
	MySQLHostEnvKey      = "MYSQL_HOST"
	MySQLHostDefault     = "127.0.0.1"
	MySQLPortEnvKey      = "MYSQL_PORT"
	MySQLPortDefault     = "4306"
	MySQLUserEnvKey      = "MYSQL_USER"
	MySQLUserDefault     = "root"
	MySQLPassEnvKey      = "MYSQL_PASSWORD"
	MySQLPassDefault     = "rootpassword"
	MySQLDatabaseEnvKey  = "MYSQL_DATABASE"
	MySQLDatabaseDefault = "sample"

	DialectMySQL = "mysql"
)

var dbConn *dbr.Connection

func init() {

	h := environ.GetEnv(MySQLHostEnvKey, MySQLHostDefault).String()
	p := environ.GetEnv(MySQLPortEnvKey, MySQLPortDefault).String()
	u := environ.GetEnv(MySQLUserEnvKey, MySQLUserDefault).String()
	pw := environ.GetEnv(MySQLPassEnvKey, MySQLPassDefault).String()
	db := environ.GetEnv(MySQLDatabaseEnvKey, MySQLDatabaseDefault).String()
	connStr := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", u, pw, h, p, db)

	conn, err := dbr.Open(DialectMySQL, connStr, nil)
	if err != nil {
		panic(err)
	}

	dbConn = conn
}

func GetMySQLInstance() *dbr.Connection {
	return dbConn
}
