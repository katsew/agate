{{- $gopath := env "GOPATH" -}}
{{- $pwd := env "PWD" -}}
{{- $relPath := trimPrefix $pwd $gopath -}}
{{- $path := trimPrefix $relPath "/src/" -}}
{{- $pkgpath := printf "%s/%s/pkg" $path PkgName -}}
package model

import (
	"{{$pkgpath}}/core"
	"github.com/gocraft/dbr"
)

type MySQLModel struct {
	master    *dbr.Session
	slave     *dbr.Session
	UseMaster bool
}

func NewMySQLModel() *MySQLModel {
	return &MySQLModel{}
}

func (m *MySQLModel) GetMaster() *dbr.Session {

	if m.master != nil {
		return m.master
	}

	sess := core.GetMySQLInstance().NewSession(nil)
	m.master = sess
	return m.master
}

func (m *MySQLModel) GetSlave() *dbr.Session {

	if m.UseMaster {
		return m.GetMaster()
	}
	if m.slave != nil {
		return m.slave
	}

	sess := core.GetMySQLInstance().NewSession(nil)
	m.slave = sess
	return m.slave
}
