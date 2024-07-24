package factory

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"go.uber.org/fx"
)

func startServer(server *http.Server) {
	if err := server.ListenAndServe(); err != nil {
		if err == http.ErrServerClosed {
			log.Logger.Println("Server closed under request")
		} else {
			log.Logger.Err(fmt.Errorf("Server closed unexpected: %s", err))
		}
	}
}

func NewServer(lc fx.Lifecycle, cfg *ServerConfig) *gin.Engine {
	app := gin.New()

	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", cfg.Port),
		Handler: app,
	}

	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			go startServer(server)
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return server.Close()
		},
	})

	return app
}

func NewBaseRoute(app *gin.Engine, cfg *ServerConfig) *gin.RouterGroup {
	group := app.Group(cfg.BaseRoute)
	group.GET("/", func(c *gin.Context) {
		c.String(200, "ola")
	})
	return group
}
