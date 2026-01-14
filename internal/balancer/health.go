package balancer

import (
	"log"
	"net/http"
	"time"
)

func StartHealthCheck(pool *BackendPool, interval time.Duration){
	client := &http.Client{
		Timeout: 2 * time.Second,
	}

	go func() {
		for { 
			for _, backend := range pool.All(){
				checkBackend(client, pool, backend)
			}

			time.Sleep(interval)
		}
	}()
}


func checkBackend(client *http.Client, pool *BackendPool, backend *Backend){
	resp, err := client.Get(backend.URL.String())
	if err != nil {
		pool.MarkAlive(backend.URL, false)
		log.Println("Backend Down: ", backend.URL)
		return 
	}

	resp.Body.Close()

	if resp.StatusCode >= 200 && resp.StatusCode <300 {
		pool.MarkAlive(backend.URL, true)
	}else{
		pool.MarkAlive(backend.URL, false)
		log.Println("Backend unhealthy: ",backend.URL)
	}
}