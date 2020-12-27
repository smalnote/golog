package proxy

import (
	"log"
	"net/http"
	"net/url"
	"time"
)

// Job represents a chrome debugging request
type Job interface {
	CreateBackend(req *http.Request) func(req *http.Request) *url.URL
	StartTimeout(done <-chan struct{}, duration time.Duration)
	CleanUp()
}

// NewJob create a chrome debugging job
func NewJob() Job {
	j := jobImpl{
		chromeProvider: chromeProvider,
	}
	return &j
}

type jobImpl struct {
	chromeProvider ChromeProvider
	chrome         Chrome
}

func (j *jobImpl) CreateBackend(req *http.Request) func(req *http.Request) *url.URL {
	return func(req *http.Request) *url.URL {
		c, err := j.chromeProvider.GetChrome()
		if err != nil {
			panic("create chrome failed")
		}
		j.chrome = c
		u, _ := url.Parse(c.GetDebuggerURL())
		query := u.Query()
		for k, vs := range req.URL.Query() {
			for _, v := range vs {
				query.Add(k, v)
			}
		}
		u.RawQuery = query.Encode()
		log.Printf("%s -> %s", req.URL.String(), u.String())
		return u
	}
}

func (j *jobImpl) StartTimeout(done <-chan struct{}, duration time.Duration) {
	go func() {
		select {
		case <-done:
			// normally done
		case t := <-time.After(duration):
			if j.chrome != nil {
				j.chrome.CleanUp()
			}
			log.Println("websocket session timeout at ", t)
		}
	}()
}

func (j *jobImpl) CleanUp() {
	if j.chrome != nil {
		j.chrome.CleanUp()
	}
}
