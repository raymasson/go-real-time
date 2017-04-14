package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/olahol/melody"
)

func main() {
	//gin is the package used to create a website in Golang
	r := gin.Default()
	//melody is the package used to handle web sockets
	m := melody.New()

	// '/' route serves the index.html file
	r.GET("/", func(c *gin.Context) {
		http.ServeFile(c.Writer, c.Request, "index.html")
	})

	// '/ws' route handles web socket requests
	r.GET("/ws", func(c *gin.Context) {
		m.HandleRequest(c.Writer, c.Request)
	})

	//Handles the incoming message and broadcast it
	m.HandleMessage(func(s *melody.Session, msg []byte) {
		m.Broadcast(msg)
	})

	//Exposes the website on port 8001
	r.Run(":8001")
}
