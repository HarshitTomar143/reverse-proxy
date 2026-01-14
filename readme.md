# HTTP Reverse Proxy & Load Balancer (Go)

A **production-style HTTP reverse proxy and Layer-7 load balancer** built from scratch in Go to understand how real-world systems like **Nginx, Envoy, and AWS ALB** work internally.

This project focuses on **systems fundamentals**: HTTP internals, concurrency, health checks, and fault tolerance — not frameworks or shortcuts.

---

## 🎯 Project Goals

- Understand how HTTP reverse proxies work internally
- Implement Layer-7 load balancing (round-robin)
- Handle backend failures gracefully
- Learn Go concurrency using goroutines and mutexes
- Build a clean, production-inspired architecture

---

## ✨ Features

### ✅ Reverse Proxy
- Accepts client HTTP requests
- Forwards requests to backend servers
- Streams request and response bodies (no buffering)

### ⚖️ Round-Robin Load Balancing (Layer 7)
- Distributes requests evenly across multiple backends
- Thread-safe backend selection
- One request → one backend

### ❤️ Active Health Checks
- Periodic background health probes (`GET /health`)
- Automatically marks unhealthy backends as DOWN
- Re-adds backends when they recover (self-healing)

### 🚀 Concurrency Safe
- Handles multiple concurrent clients
- Uses goroutines and mutexes
- No race conditions



