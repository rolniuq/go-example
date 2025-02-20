package eventbus

import "sync"

type HandleFunc func(interface{})

type EventBus struct {
	subscribers map[string][]HandleFunc
	mu          sync.Mutex
}
