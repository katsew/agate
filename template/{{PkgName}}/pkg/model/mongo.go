{{- $gopath := env "GOPATH" -}}
{{- $pwd := env "PWD" -}}
{{- $relPath := trimPrefix $pwd $gopath -}}
{{- $path := trimPrefix $relPath "/src/" -}}
{{- $pkgpath := printf "%s/%s/pkg" $path PkgName -}}
package model

import (
	"{{$pkgpath}}/core"
	"gopkg.in/mgo.v2"
)

type MongoModel struct {
	Session *mgo.Session    `json:"-" bson: "-"`
	DB      *mgo.Database   `json:"-" bson: "-"`
	C       *mgo.Collection `json:"-" bson: "-"`
	dbName  string          `json:"-" bson: "-"`
	colName string          `json:"-" bson: "-"`
}

func NewMongoModel(dbName string, colName string) MongoModel {
	session := core.GetMongoInstance().Copy()
	return MongoModel{
		Session: session,
		DB:      session.DB(dbName),
		C:       session.DB(dbName).C(colName),
		dbName:  dbName,
		colName: colName,
	}
}

func (mm *MongoModel) GetDBName() string {
	return mm.dbName
}

func (mm *MongoModel) GetColName() string {
	return mm.colName
}
