package runtime

import "sync"

type RunManager struct {
	mu     sync.RWMutex

	runner Runner
}
