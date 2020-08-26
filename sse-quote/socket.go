package ssequote

import (
	"fmt"
	"github.com/gin-gonic/gin"
	socketio "github.com/googollee/go-socket.io"
	"log"
)

var (
	server *socketio.Server
	err    error
)

func GinMiddleware(allowOrigin string) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", allowOrigin)
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, Content-Length, X-CSRF-Token, Token, session, Origin, Host, Connection, Accept-Encoding, Accept-Language, X-Requested-With")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Request.Header.Del("Origin")

		c.Next()
	}
}

func SocketIO() {
	gin.SetMode(Conf.GinMode)
	router := gin.New()
	server, err = socketio.NewServer(nil)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		_ = server.Close()
	}()

	server.OnConnect("/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})
	server.OnEvent("/", "sse-request", func(s socketio.Conn, msg string) {
		sseHandler(s, msg)
	})
	server.OnError("/", func(s socketio.Conn, e error) {
		fmt.Println("error:", e)
	})
	server.OnDisconnect("/", func(s socketio.Conn, reason string) {
		fmt.Println("closed", reason)
	})

	go func() {
		_ = server.Serve()
	}()

	router.Use(GinMiddleware("*"))
	router.GET("/sse/*any", gin.WrapH(server))
	router.POST("/sse/*any", gin.WrapH(server))

	_ = router.Run(Conf.Socket.GetHost())
}
