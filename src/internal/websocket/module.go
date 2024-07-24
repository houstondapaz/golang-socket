package websocket

import (
	"fmt"
	"github.com/houstondapaz/golang-socket/internal/websocket/routes"
	"github.com/houstondapaz/golang-socket/internal/websocket/shared"
	"github.com/houstondapaz/golang-socket/pkg/api"

	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Params struct {
	fx.In
	Router *gin.RouterGroup `name:"main_route"`
	Routes []api.Router     `group:"websocket_routes"`
}

func Module() fx.Option {
	return fx.Module(shared.PackageName,
		fx.Provide(
			fx.Annotate(
				routes.NewWebsocketRoute,
				fx.As(new(api.Router)),
				fx.ResultTags(fmt.Sprintf(`group:"%s"`, shared.RoutersGroup)),
			),
		),
		fx.Invoke(func(p Params) error {
			for _, router := range p.Routes {
				router.Register(p.Router)
			}

			return nil
		}),
	)
}
