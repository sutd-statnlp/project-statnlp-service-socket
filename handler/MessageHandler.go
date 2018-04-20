package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// ConfigRequest .
func ConfigRequest(router *gin.Engine) {
	router.GET("/socket", HanldeSocket)
}

// HanldeSocket .
func HanldeSocket(context *gin.Context) {
	wsupgrader := GetUgrader()
	conn, err := wsupgrader.Upgrade(context.Writer, context.Request, nil)
	if err != nil {
		return
	}
	for {
		messageType, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		go HandleMessage(conn, messageType, message)
	}
}

// GetUgrader .
func GetUgrader() websocket.Upgrader {
	return websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}
}

// HandleMessage handles message from client.
func HandleMessage(conn *websocket.Conn, messageType int, message []byte) {
	conn.WriteMessage(messageType, message)
}
