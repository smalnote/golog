package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/gin-gonic/gin"
)

// TestWeb test a http server that can be shutdown gracefully
func main() {
	router := gin.New()

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	server.RegisterOnShutdown(func() {
		log.Println("clean up")
	})

	go func(router *gin.Engine) {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal("server startup: ", err)
		}
	}(router)

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// wait for close signal
	<-quit

	ctx := context.Background()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal("server shutdown: ", err)
	}

	log.Println("server is shutdown")
}
