package ctrl

import (
	"net/http"

	"github.com/labstack/echo"
)

type GetHealthz struct{}

func (gh *GetHealthz) GetMethod() string {
	return echo.GET
}

func (gh *GetHealthz) GetPath() string {
	return "/healthz"
}

func (gh *GetHealthz) GetHandler() echo.HandlerFunc {
	return func(c echo.Context) error {
		c.JSON(http.StatusOK, map[string]string{"status": "200", "health": "ok", "msg": "everything is going well."})
		return nil
	}
}

func (gh *GetHealthz) GetMiddlewares() ([]echo.MiddlewareFunc, bool) {
	return nil, false
}
