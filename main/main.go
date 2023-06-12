package main

import (
	"net/http/httputil"
	"net/url"
	"sync"
	"sync/atomic"
)

type Backend struct {
	URL          *url.URL
	Alive        bool
	mutex        sync.RWMutex
	ReverseProxy *httputil.ReverseProxy
}

type ServerPool struct {
	backends []*Backend
	current  uint64
}

func (sp *ServerPool) NextIndex() int {
	return int(atomic.AddUint64(&sp.current, uint64(1)) % uint64(len(sp.backends)))
}

func (b *Backend) SetAlive(alive bool) {
	b.mutex.Lock()
	b.Alive = alive
	b.mutex.Unlock()
}

func (b *Backend) IsAlive() (alive bool) {
	b.mutex.RLock()
	alive = b.Alive
	b.mutex.RUnlock()
	return
}
