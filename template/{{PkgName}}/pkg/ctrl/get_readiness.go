package ctrl

import (
	"net/http"

	"github.com/labstack/echo"
)

type GetReadiness struct{}

func (gr *GetReadiness) GetMethod() string {
	return echo.GET
}

func (gr *GetReadiness) GetPath() string {
	return "/readiness"
}

//@todo Implement
func (gr *GetReadiness) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {

		c.Logger().Warn("@todo Implement correctly...")
		c.JSON(http.StatusOK, map[string]string{"status": "200", "health": "ok", "msg": "everything is going well."})
		return nil

	}
}

func (gr *GetReadiness) GetMiddlewares() ([]echo.MiddlewareFunc, bool) {
	return nil, false
}
