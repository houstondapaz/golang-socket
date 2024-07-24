package routes

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/rs/zerolog/log"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type WebsocketRoute struct{}

func HandleWebsocket(ctx *gin.Context) {
	conn, err := upgrader.Upgrade(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.JSON(500, gin.H{"error": "Failed to upgrade to websocket"})
		return
	}
	defer conn.Close()

	for {
		// Read message from client
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Err(fmt.Errorf("error while reading message: %s", err))
			ctx.AbortWithError(http.StatusInternalServerError, err)
			break
		}

		// Echo message back to client
		err = conn.WriteMessage(messageType, p)
		if err != nil {
			log.Err(fmt.Errorf("error while writing message: %s", err))
			ctx.AbortWithError(http.StatusInternalServerError, err)
			break
		}
	}

	ctx.Writer.WriteString("websocket ok")
}

func (r *WebsocketRoute) Register(router *gin.RouterGroup) {
	router.GET("/websocket", HandleWebsocket)
}

func NewWebsocketRoute() *WebsocketRoute {
	return &WebsocketRoute{}
}
