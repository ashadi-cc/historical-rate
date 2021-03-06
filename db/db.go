package db

import (
	"fmt"
	"sync"

	"history-rate/db/provider"
)

var (
	providersMu sync.RWMutex
	providers   = make(map[string]provider.Provider)
)

// Register register the provider
func Register(name string, provider provider.Provider) {
	providersMu.Lock()
	defer providersMu.Unlock()
	if provider == nil {
		panic("repo: Register provider is nil")
	}
	if _, dup := providers[name]; dup {
		panic("repo: Register called twice for driver " + name)
	}
	providers[name] = provider
}

func Open(name string) (provider.Provider, error) {
	providersMu.RLock()
	defer providersMu.RUnlock()
	provideri, ok := providers[name]
	if !ok {
		return nil, fmt.Errorf("unknown provider %q (forgotten import?)", name)
	}

	if err := provideri.Init(); err != nil {
		return nil, err
	}

	return provideri, nil
}
