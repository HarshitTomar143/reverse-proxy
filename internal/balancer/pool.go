package balancer

import (
	"errors"
	"net/url"
	"sync"
)

type Backend struct {
	URL   *url.URL
	Alive bool
}

type BackendPool struct {
	backends []*Backend
	index    int
	mu       sync.RWMutex
}

func NewBackendPool(urls []string) (*BackendPool, error) {
	var backends []*Backend

	for _, raw := range urls {
		u, err := url.Parse(raw)
		if err != nil {
			return nil, err
		}
		backends = append(backends, &Backend{
			URL:   u,
			Alive: true,
		})
	}

	return &BackendPool{backends: backends}, nil
}

func (p *BackendPool) Next() (*Backend, error) {
	p.mu.Lock()
	defer p.mu.Unlock()

	n := len(p.backends)
	for i := 0; i < n; i++ {
		backend := p.backends[p.index]
		p.index = (p.index + 1) % n

		if backend.Alive {
			return backend, nil
		}
	}

	return nil, errors.New("no healthy backends")
}

func (p *BackendPool) MarkAlive(url *url.URL, alive bool) {
	p.mu.Lock()
	defer p.mu.Unlock()

	for _, b := range p.backends {
		if b.URL.String() == url.String() {
			b.Alive = alive
			return
		}
	}
}

// ✅ THIS WAS MISSING
func (p *BackendPool) All() []*Backend {
	p.mu.RLock()
	defer p.mu.RUnlock()
	return p.backends
}
