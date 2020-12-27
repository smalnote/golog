// Package proxy proxies websocket and set timeout to the connection
package proxy

import (
	"os/exec"
	"time"

	"github.com/gin-gonic/gin"
	wp "github.com/koding/websocketproxy"
)

var chromeProvider ChromeProvider

func init() {
	chromeProvider = NewChromeProvider(true)
}

// Service is a websocket proxy service
// proxy request to a chrome debugger port
type Service interface {
	Handle(c *gin.Context)
}

// NewService create a proxy service
func NewService() Service {
	return &serviceImpl{}
}

type serviceImpl struct {
	cmd *exec.Cmd
}

// Handle implement gin.HandleFunc
func (s *serviceImpl) Handle(c *gin.Context) {
	job := NewJob()
	done := make(chan struct{}, 1)
	job.StartTimeout(done, 10*time.Second)
	defer job.CleanUp()
	proxie := wp.WebsocketProxy{
		Backend: job.CreateBackend(c.Request),
	}
	proxie.ServeHTTP(c.Writer, c.Request)
	done <- struct{}{} // job done normally
}
