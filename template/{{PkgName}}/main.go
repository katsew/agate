{{- $gopath := env "GOPATH" -}}
{{- $pwd := env "PWD" -}}
{{- $relPath := trimPrefix $pwd $gopath -}}
{{- $path := trimPrefix $relPath "/src/" -}}
{{- $pkgpath := printf "%s/%s" $path PkgName -}}
package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"{{$pkgpath}}/pkg/conf"
	"{{$pkgpath}}/pkg/mw"
)

func main() {

	e := echo.New()

	e.Pre(
		middleware.RecoverWithConfig(middleware.DefaultRecoverConfig),
		middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
			c.Logger().Info("RequestBody: ", string(reqBody))
		}),
		mw.FastRetCors(middleware.CORSConfig{
			AllowOrigins: conf.AllowOrigins[conf.Env],
			AllowMethods: conf.AllowMethods[conf.Env],
			AllowHeaders: conf.AllowHeaders[conf.Env],
		}),
		mw.MaintenanceHandler(),
	)

	for _, route := range conf.Routes {
		if mws, ok := route.GetMiddlewares(); ok {
			e.Add(route.GetMethod(), route.GetPath(), route.GetHandler(), mws...)
		} else {
			if mws != nil || len(mws) > 0 {
				e.Logger.Warn("Route have some middleware but not registered. Did you forget to set ok on GetMiddlewares func?")
			}
			e.Add(route.GetMethod(), route.GetPath(), route.GetHandler())
		}
	}

	go func() {
		if err := e.Start(":8080"); err != nil {
			e.Logger.Info("Shutting down the server")
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Print("Shutdown server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := e.Shutdown(ctx); err != nil {
		e.Logger.Fatal(err)
	}

}
