package bindings

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/natrongmbh/natrium/pkg/config"
	"github.com/natrongmbh/natrium/pkg/server"
	"github.com/pocketbase/pocketbase/core"
)

func BindAppHooks() error {
	server.App.OnBeforeServe().Add(func(e *core.ServeEvent) error {
		// add new "GET /api/hello" route
		_, err := e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/version",
			Handler: func(c echo.Context) error {
				return c.JSON(200, config.ServerConfig.Version)
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
	return nil

}
