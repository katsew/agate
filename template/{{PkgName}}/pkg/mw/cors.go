package mw

import (
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	m "github.com/labstack/echo/middleware"
)

func FastRetCors(config m.CORSConfig) echo.MiddlewareFunc {

	if config.Skipper == nil {
		config.Skipper = m.DefaultSkipper
	}
	if len(config.AllowOrigins) == 0 {
		config.AllowOrigins = m.DefaultCORSConfig.AllowOrigins
	}
	if len(config.AllowMethods) == 0 {
		config.AllowMethods = m.DefaultCORSConfig.AllowMethods
	}

	allowMethods := strings.Join(config.AllowMethods, ",")
	allowHeaders := strings.Join(config.AllowHeaders, ",")
	exposeHeaders := strings.Join(config.ExposeHeaders, ",")
	maxAge := strconv.Itoa(config.MaxAge)

	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config.Skipper(c) {
				return next(c)
			}

			req := c.Request()
			res := c.Response()
			origin := req.Header.Get(echo.HeaderOrigin)
			allowOrigin := ""
			method := req.Method
			allowMethod := ""

			// Check allowed origins
			for _, o := range config.AllowOrigins {
				if o == "*" || o == origin {
					allowOrigin = o
					break
				}
			}
			if req.Method == echo.OPTIONS {
				// Preflight request
				res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
				res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestMethod)
				res.Header().Add(echo.HeaderVary, echo.HeaderAccessControlRequestHeaders)
				res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
				res.Header().Set(echo.HeaderAccessControlAllowMethods, allowMethods)
				if config.AllowCredentials {
					res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
				}
				if allowHeaders != "" {
					res.Header().Set(echo.HeaderAccessControlAllowHeaders, allowHeaders)
				} else {
					h := req.Header.Get(echo.HeaderAccessControlRequestHeaders)
					if h != "" {
						res.Header().Set(echo.HeaderAccessControlAllowHeaders, h)
					}
				}
				if config.MaxAge > 0 {
					res.Header().Set(echo.HeaderAccessControlMaxAge, maxAge)
				}
				return c.NoContent(http.StatusNoContent)
			}

			// If request origin is not allowed, return http error.
			if allowOrigin == "" {
				c.Logger().Warn("Origin is not allowed...")
				res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
				res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
				if config.AllowCredentials {
					res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
				}
				if exposeHeaders != "" {
					res.Header().Set(echo.HeaderAccessControlExposeHeaders, exposeHeaders)
				}
				return echo.NewHTTPError(http.StatusBadRequest)
			}

			// Check allowed methods
			for _, m := range config.AllowMethods {
				if m == method {
					allowMethod = m
					break
				}
			}
			// If request method is not allowed, return http error.
			if allowMethod == "" {
				res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
				res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
				if config.AllowCredentials {
					res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
				}
				if exposeHeaders != "" {
					res.Header().Set(echo.HeaderAccessControlExposeHeaders, exposeHeaders)
				}
				return echo.NewHTTPError(http.StatusMethodNotAllowed)
			}

			res.Header().Add(echo.HeaderVary, echo.HeaderOrigin)
			res.Header().Set(echo.HeaderAccessControlAllowOrigin, allowOrigin)
			if config.AllowCredentials {
				res.Header().Set(echo.HeaderAccessControlAllowCredentials, "true")
			}
			if exposeHeaders != "" {
				res.Header().Set(echo.HeaderAccessControlExposeHeaders, exposeHeaders)
			}
			return next(c)

		}
	}
}
