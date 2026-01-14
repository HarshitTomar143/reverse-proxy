package proxy

import (
	"io"
	"net/http"

	"github.com/harshit/load-balancer/internal/balancer"

)

type ReverseProxy struct {
	pool   *balancer.BackendPool
	client *http.Client
}

func NewReverseProxy(pool *balancer.BackendPool) *ReverseProxy {
	return &ReverseProxy{
		pool:   pool,
		client: &http.Client{},
	}
}

func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	backend, err := p.pool.Next()
	if err != nil {
		http.Error(w, "Service Unavailable", http.StatusServiceUnavailable)
		return
	}

	req, err := http.NewRequest(
		r.Method,
		backend.URL.String()+r.URL.RequestURI(),
		r.Body,
	)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	for k, v := range r.Header {
		for _, val := range v {
			req.Header.Add(k, val)
		}
	}

	req.Host = backend.URL.Host

	resp, err := p.client.Do(req)
	if err != nil {
		p.pool.MarkAlive(backend.URL, false)
		http.Error(w, "Bad Gateway", http.StatusBadGateway)
		return
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		for _, val := range v {
			w.Header().Add(k, val)
		}
	}

	w.WriteHeader(resp.StatusCode)
	io.Copy(w, resp.Body)
}
