package handler

import (
	"github.com/gin-gonic/gin"
)

func ping(c *gin.Context) {
	messages := []string{}
	messages = append(messages, "test1")
	messages = append(messages, "test2")
	c.JSON(200, gin.H{
		"messages": messages,
	})
}
