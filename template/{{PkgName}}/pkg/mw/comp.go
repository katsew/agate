{{- $gopath := env "GOPATH" -}}
{{- $pwd := env "PWD" -}}
{{- $relPath := trimPrefix $pwd $gopath -}}
{{- $path := trimPrefix $relPath "/src/" -}}
{{- $pkgpath := printf "%s/%s/pkg" $path PkgName -}}
package mw

import (
	"net/http"

	"{{$pkgpath}}/conf"
	"{{$pkgpath}}/resp"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
)

func Compositor() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			if err := next(c); err != nil {
				return errors.Wrap(err, "Failed to handle route")
			}

			var res resp.Response
			if apiRes, ok := c.Get(conf.CtxApiResultKey).(resp.APIResult); ok {
				res.APIResult = apiRes
			} else {
				return errors.New("Failed to handle compositor")
			}
			if serverReq, ok := c.Get(conf.CtxServerRequestKey).(resp.ServerRequest); ok {
				res.ServerRequest = serverReq
			}

			res.SystemInfo = resp.NewSystemInfo()

			if err := c.JSON(http.StatusOK, res); err != nil {
				c.Logger().Error(err)
			}

			return nil
		}
	}
}
