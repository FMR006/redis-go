package storage

import (
	"sync"
	"time"
)

type Storage struct {
	storage  map[string]string
	expireAt map[string]time.Time
	mu       *sync.RWMutex
}

func (s *Storage) NewStorage() *Storage {
	return &Storage{
		storage:  make(map[string]string),
		expireAt: make(map[string]time.Time),
		mu:       &sync.RWMutex{},
	}
}
