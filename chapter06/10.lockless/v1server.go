package main

import (
	"fmt"
	"sync"
	"time"
)

type WebServer struct {
	config     Config
	configLock sync.RWMutex
}

func (w *WebServer) visit() string {
	w.configLock.RLock()
	defer w.configLock.RUnlock()
	return w.config.context
}

func (w *WebServer) reload() {
	w.configLock.Lock()
	defer w.configLock.Unlock()
	w.config.context = fmt.Sprintf("%d", time.Now().UnixNano())
}

func (w *WebServer) ReloadWork() {
	for true {
		time.Sleep(10 * time.Millisecond)
		w.reload()
	}
}
