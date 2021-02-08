package resourcepool

import (
	"context"
	"errors"
	"sync"
)

// Resource pooled element
type Resource interface {
	// Close release resource
	Close()
}

// ResourcePool manage resource
type ResourcePool interface {
	GetResource(ctx context.Context) (Resource, error)
	Put(c Resource)
	Close()
}

// New create a ResourcePool
func New() ResourcePool {
	return &pool{
		mu:    &sync.RWMutex{},
		frees: make([]Resource, 0),
		reqs:  make([]chan req, 0),
	}
}

type req struct {
	c Resource
	e error
}

type pool struct {
	mu    *sync.RWMutex
	frees []Resource
	reqs  []chan req
}

func (p *pool) GetResource(ctx context.Context) (Resource, error) {
	p.mu.Lock()

	if len(p.frees) > 0 {
		conn := p.frees[0]
		copy(p.frees, p.frees[1:])
		p.frees = p.frees[:len(p.frees)-1]
		defer p.mu.Unlock()
		return conn, nil
	}

	r := make(chan req, 1)
	p.reqs = append(p.reqs, r)
	p.mu.Unlock()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case ret, ok := <-r:
		if !ok {
			return nil, errors.New("pool closed")
		}
		return ret.c, ret.e
	}
}

func (p *pool) Close() {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, r := range p.reqs {
		close(r)
	}
	p.reqs = nil

	for _, r := range p.frees {
		r.Close()
	}
	p.frees = nil
}

func (p *pool) Put(c Resource) {
	if c == nil {
		panic("resource cannot be nil")
	}
	p.mu.Lock()
	defer p.mu.Unlock()

	if len(p.reqs) > 0 {
		r := p.reqs[0]
		copy(p.reqs, p.reqs[1:])
		p.reqs = p.reqs[:len(p.reqs)-1]
		r <- req{c: c}
		return
	}

	p.frees = append(p.frees, c)
}
