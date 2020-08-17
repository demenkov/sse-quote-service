package main

import (
	"./sse-quote"
	"github.com/joho/godotenv"
	"log"
)

var conf *ssequote.Config

func init() {
	if err := godotenv.Load(); err != nil {
		log.Print("No .env file found")
	}
	conf = ssequote.NewConfig()
	ssequote.InitPackage(conf)
}

func main() {
	ssequote.LogInfo("Go SSE Quote Service started!!!")
	ssequote.SocketIO()
}
