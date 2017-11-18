package conf

import (
	"github.com/labstack/echo"
)

type CorsMap map[string][]string

var AllowOrigins CorsMap = CorsMap{
	EnvLocalhost:   {"*"},
	EnvDevelopment: {"*"},
	EnvStaging:     {"*"},
	EnvProduction:  {"*"},
}

var AllowMethods CorsMap = CorsMap{
	EnvLocalhost:   {echo.GET, echo.POST, echo.OPTIONS, echo.HEAD},
	EnvDevelopment: {echo.GET, echo.POST, echo.OPTIONS, echo.HEAD},
	EnvStaging:     {echo.GET, echo.POST, echo.OPTIONS, echo.HEAD},
	EnvProduction:  {echo.GET, echo.POST, echo.OPTIONS, echo.HEAD},
}

var AllowHeaders CorsMap = CorsMap{
	EnvLocalhost:   {"X-Api-Token"},
	EnvDevelopment: {},
	EnvStaging:     {},
	EnvProduction:  {},
}
