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
}

func NewChrome() (Chrome, error) {
	chrome := chromeImpl{}
	err := chrome.Start()
	if err != nil {
		return nil, err
	}
	return &chrome, nil
}

type version struct {
	WebSocketDebuggerURL string `json:"webSocketDebuggerUrl"`
}

func (c *chromeImpl) Start() error {
	if c.started {
		return nil
	}
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
	<-time.After(1000 * time.Millisecond)
	versionURL := fmt.Sprintf("http://127.0.0.1:%d/json/version", port)
	resp, err := http.Get(versionURL)
	if err != nil {
		return err
	}

	version := version{}
	err = json.NewDecoder(resp.Body).Decode(&version)
	if err != nil {
		return err
	}
	c.WebSocketDebuggerURL = version.WebSocketDebuggerURL
	log.Printf("debugger url: %s", version.WebSocketDebuggerURL)
	return nil
}

func (c *chromeImpl) getPort() (int, error) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	if err != nil {
		return 0, err
	}
	defer listener.Close()
	return listener.Addr().(*net.TCPAddr).Port, nil
}

func (c *chromeImpl) GetDebuggerURL() string {
	return c.WebSocketDebuggerURL
}

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
