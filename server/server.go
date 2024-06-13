package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

const PORT = "8000"

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func StartServer() {
	server := gin.Default()

	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	server.LoadHTMLFiles("index.html")
	server.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", nil)
	})
	server.GET("/chat-data", GetChatData)
	server.POST("/chat-data", PostChatData)
	server.GET("/ws", func(c *gin.Context) {
		WsHandler(c.Writer, c.Request)
	})

	server.Run("0.0.0.0:" + PORT) // listen and serve on 0.0.0.0:8080
}

func WsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering WsHandler")
	conn, err := wsupgrader.Upgrade(w, r, nil)
	fmt.Println("Connection upgraded")
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Failed reading message %+v", err)
			break
		}
		fmt.Println("data %s", string(msg))
		conn.WriteMessage(t, msg)
	}
}
