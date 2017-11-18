{{- $gopath := env "GOPATH" -}}
{{- $pwd := env "PWD" -}}
{{- $relPath := trimPrefix $pwd $gopath -}}
{{- $path := trimPrefix $relPath "/src/" -}}
{{- $pkgpath := printf "%s/%s/pkg" $path PkgName -}}
package conf

import (
	"github.com/labstack/echo"
	c "{{$pkgpath}}/ctrl"
)

type Route interface {
	GetMethod() string
	GetPath() string
	GetHandler() echo.HandlerFunc
	GetMiddlewares() (mws []echo.MiddlewareFunc, ok bool)
}

var Routes []Route

func init() {
	Routes = []Route{
		&c.GetHealthz{},
		&c.GetReadiness{},
	}
}

