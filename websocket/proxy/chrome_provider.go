package proxy

import (
	"log"
	"os"
	"os/signal"
	"syscall"
)

// ChromeProvider provides chrome from preboot swarm or create new as need
type ChromeProvider interface {
	// GetChrome create a chrome debugging port
	GetChrome() (Chrome, error)
}

type chromeProviderImpl struct {
	preboot     bool
	chromeSwarm (chan Chrome)
}

// NewChromeProvider chreate a ChromeProvider
func NewChromeProvider(preboot bool) ChromeProvider {
	cpi := chromeProviderImpl{
		preboot:     preboot,
		chromeSwarm: make(chan Chrome, 2),
	}
	if preboot {
		go cpi.prebootChromeSwarm()
	}
	return &cpi
}

func (p *chromeProviderImpl) prebootChromeSwarm() {
	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGINT, syscall.SIGTERM)
	var chrome Chrome
	for cleaned := false; !cleaned; {
		chrome = p.newChrome()
		select {
		case <-done:
			log.Println("start clean up chrome swarm")
			close(p.chromeSwarm)
			for c := range p.chromeSwarm {
				c.CleanUp()
			}
			cleaned = true
		case p.chromeSwarm <- chrome:
			// preboot a chrome
		}
	}
	chrome.CleanUp()
	log.Println("chrome swarm stopped")
	os.Exit(0)
}

func (p *chromeProviderImpl) newChrome() Chrome {
	chrome, err := NewChrome()
	if err != nil {
		panic("preboot chrome swarm failed: " + err.Error())
	}
	return chrome
}

func (p *chromeProviderImpl) GetChrome() (Chrome, error) {
	if p.preboot {
		select {
		case chrome := <-p.chromeSwarm:
			return chrome, nil
		default:
			// fall back to creating a new one
		}
	}
	return NewChrome()
}
