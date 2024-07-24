package main

import (
	"github.com/houstondapaz/golang-socket/cmd/server"
	"github.com/houstondapaz/golang-socket/internal/websocket"

	"github.com/joho/godotenv"
	"go.uber.org/fx"
)

func main() {
	godotenv.Load()

	app := fx.New(
		server.NewModule(),
		websocket.Module(),
	)

	app.Run()
}
