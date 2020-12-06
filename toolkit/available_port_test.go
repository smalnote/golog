package toolkit

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"testing"
	"time"
)

func TestGetAvailablePort(t *testing.T) {
	port := 8080
	go occupiedPort(port)
	<-time.After(500 * time.Microsecond)

	start := time.Now()
	randomPort, err := getPort(port)
	if err != nil {
		t.Error(err)
		return
	}
	duration := time.Now().Sub(start)
	t.Logf("time cost for get port: %dms", duration.Milliseconds())
	t.Logf("attempt to listen on port: %d", randomPort)

	if port == randomPort {
		t.Error("random port is equal to occupied port")
	}
	<-time.After(1 * time.Second)
}

func occupiedPort(port int) {
	server := http.Server{
		Addr: fmt.Sprintf(":%d", port),
	}
	done := make(chan struct{})
	go func(server *http.Server, done chan<- struct{}) {
		defer func() {
			log.Println("server shutdown normally")
			done <- struct{}{}
		}()
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			log.Fatal(err)
		}
	}(&server, done)

	<-time.After(800 * time.Millisecond)
	err := server.Shutdown(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func getPort(port int) (int, error) {
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		if port > 0 {
			return getPort(0)
		}
		return -1, err
	}
	defer listener.Close()

	return listener.Addr().(*net.TCPAddr).Port, nil
}
