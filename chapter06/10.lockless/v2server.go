package main

import (
	"fmt"
	"sync"
	"time"
)

type WebServer2 struct {
	config     *Config
	configLock sync.RWMutex
}

func (w *WebServer2) visit() string {
	//w.configLock.RLock()
	//defer w.configLock.RUnlock()
	return w.config.context
}

func (w *WebServer2) reload() {
	//w.configLock.Lock()
	//defer w.configLock.Unlock()
	w.config = &Config{
		context: fmt.Sprintf("%d", time.Now().UnixNano()),
	}
	//fmt.Sprintf("%d", time.Now().UnixNano())
}

func (w *WebServer2) ReloadWork() {
	for true {
		time.Sleep(10 * time.Millisecond)
		w.reload()
	}
}
