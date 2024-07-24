package server

import (
	"github.com/houstondapaz/golang-socket/cmd/server/factory"

	"go.uber.org/fx"
)

func NewModule() fx.Option {
	return fx.Options(
		fx.Provide(factory.NewApiConfig),
		fx.Provide(factory.NewServer),
		fx.Provide(fx.Annotate(factory.NewBaseRoute, fx.ResultTags(`name:"main_route"`))),
	)
}
