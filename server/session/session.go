package session

import (
	"sync"
)

type Manager struct {
	cookieName  string     // private cookiename
	lock        sync.Mutex // protects session
	provider    Provider
	maxLifeTime int64
}

type Provider interface {
	See
}

func NewManager(provideName, cookieName stirng, maxLifeTime int64) (*Manager, error) {
	if provider, ok := provides[provideName]; !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgottoen import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifeTime: maxLifeTime}, nil
}
