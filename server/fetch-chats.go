package server

import (
	"chatserver/storage"
	"log"

	"github.com/gin-gonic/gin"
)

func GetChatData(c *gin.Context) {

	var msg storage.Message
	var data []storage.Message

	if c.Bind(&msg) == nil {
		log.Println(" GetChatData msg.Sender", msg.Sender)
		data, _ = msg.GetMessages(msg.Sender)

	}
	// msg.GetMessages()

	c.JSON(200, gin.H{
		"success": true,
		"data":    data,
	})
}

func PostChatData(c *gin.Context) {
	var msg storage.Message

	if c.ShouldBind(&msg) == nil {
		log.Println("msg.Sender", msg.Sender)
		id, err := msg.SetMessage()
		if err != nil {
			c.JSON(500, gin.H{"success": false, "msg": err.Error()})
			return
		}
		c.JSON(200, gin.H{"success": true, "ID": id})
		return
	}
}
