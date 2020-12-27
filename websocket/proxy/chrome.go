package proxy

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os/exec"
	"sync"
	"time"
)

// Chrome represents a headless chrome instance
type Chrome interface {
	Start() error
	GetDebuggerURL() string
	CleanUp() error
}

type chromeImpl struct {
	started              bool
	cmd                  *exec.Cmd
	port                 int
	WebSocketDebuggerURL string
	cleanUpErr           error
	retry                int
}

// NewChrome create a chrome instance
func NewChrome() (Chrome, error) {
	chrome := chromeImpl{
		retry: 3,
	}
	err := chrome.Start()
	if err != nil {
		return nil, err
	}
	return &chrome, nil
}

type version struct {
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

// Start run a headless chrome and get its debugger URL
func (c *chromeImpl) Start() error {
	if c.started {
		return nil
	}
	start := time.Now()
	port, err := c.getPort()
	if err != nil {
		return err
	}
	c.port = port

	remoteDebuggingPortArg := fmt.Sprintf("--remote-debugging-port=%d", port)

	c.cmd = exec.Command(`C:\Program Files (x86)\Google\Chrome\Application\Chrome.exe`,
		"--headless", " --no-sandbox", "--no-first-run", remoteDebuggingPortArg)

	if err := c.cmd.Start(); err != nil {
		return err
	}
	for i := 0; i < (c.retry + 1); i++ {
		// <-time.After(time.Duration((i+1)*10) * time.Millisecond)
		url, err := c.getDebuggerURL()
		c.WebSocketDebuggerURL = url
		if err == nil {
			break
		}
		if i == c.retry {
			return fmt.Errorf("get debugger url error: %w", err)
		}
	}
	duration := time.Since(start)
	log.Printf("debugger url: %s, cost: %vms", c.WebSocketDebuggerURL, duration.Milliseconds())
	c.started = true
	return nil
}

func (c *chromeImpl) getDebuggerURL() (string, error) {
	versionURL := fmt.Sprintf("http://127.0.0.1:%d/json/version", c.port)
	resp, err := http.Get(versionURL)
	if err != nil {
		return "", err
	}
	version := version{}
	err = json.NewDecoder(resp.Body).Decode(&version)
	if err != nil {
		return "", err
	}
	return version.WebSocketDebuggerURL, nil
}

// randomly generate a available port of localhost
func (c *chromeImpl) getPort() (int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

// GetDebuggerURL return the debugger URL
func (c *chromeImpl) GetDebuggerURL() string {
	return c.WebSocketDebuggerURL
}

// CleanUp kill the chrome process(only execute once)
func (c *chromeImpl) CleanUp() error {
	once := sync.Once{}
	once.Do(c.cleanUpOnce)
	return c.cleanUpErr
}

func (c *chromeImpl) cleanUpOnce() {
	var cleanUpErr error
	if c.cmd != nil && c.cmd.Process != nil {
		if err := c.cmd.Process.Kill(); err != nil {
			cleanUpErr = fmt.Errorf("kill chrome process[%d]: %w", c.port, err)
			log.Printf("kill chrome process[%d]: %s", c.port, err.Error())
		}

		if err := c.cmd.Wait(); err != nil {
			cleanUpErr = fmt.Errorf("wait for cmd[%d]: %w", c.port, err)
		}
		log.Printf("debugging port %d is cleaned", c.port)
	}
	c.cleanUpErr = cleanUpErr
}
