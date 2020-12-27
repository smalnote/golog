package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/smalnote/golog/websocket/proxy"
)

func main() {
	gin := gin.Default()

	proxy := proxy.NewService()

	gin.GET("/api/debugger", proxy.Handle)

	if err := http.ListenAndServe(":8000", gin); err != nil {
		log.Fatal(err)
	}

}
