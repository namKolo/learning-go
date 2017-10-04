package db

import (
	"fmt"
	"sync"
)

type repository struct {
	items map[string]string
	mu    sync.RWMutex
}

func (r *repository) Set(key, data string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.items[key] = data
}

func (r *repository) Get(key string) (string, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	item, ok := r.items[key]
	if !ok {
		return "", fmt.Errorf("The '%s' is not presented", key)
	}
	return item, nil
}

var (
	r    *repository
	once sync.Once
)

func Repository() *repository {
	once.Do(func() {
		r = &repository{
			items: make(map[string]string),
		}
	})

	return r
}
