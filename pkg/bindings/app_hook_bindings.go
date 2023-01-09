package bindings

import (
	"net/http"

	"github.com/labstack/echo/v5"
	"github.com/natrongmbh/natrium/pkg/config"
	"github.com/natrongmbh/natrium/pkg/k8s"
	"github.com/natrongmbh/natrium/pkg/server"
	"github.com/natrongmbh/natrium/pkg/util"
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

		_, err = e.Router.AddRoute(echo.Route{
			Method: http.MethodGet,
			Path:   "/api/cluster/version",
			Handler: func(c echo.Context) error {
				version, err := k8s.GetClusterVersion()
				if err != nil {
					util.Logger.Error(err)
				}
				return c.JSON(200, version)
			},
		})
		if err != nil {
			return err
		}

		return nil
	})
	return nil

}
