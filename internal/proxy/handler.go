package proxy

import (
	"io"
	"log"
	"net/http"
	"net/url"
)

type ReverseProxy struct{
	backend *url.URL
	client *http.Client
}

// creates A new reverseProxy instance
func NewReverseProxy(backend string) *ReverseProxy{
	backendURL, err := url.Parse(backend)
	if err != nil {
		log.Fatalf("invalid backend URL: %v",err)
	}

	return &ReverseProxy{
		backend: backendURL,
		client: &http.Client{},
	}
}

// ServeHTTP makes ReverseProxy an http.Handler
func (p *ReverseProxy) ServeHTTP(w http.ResponseWriter, r *http.Request){

}