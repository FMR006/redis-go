package storage

import (
	"sync"
	"time"

	"github.com/FMR006/redis-go/internal/resp"
)

type ValueType uint8

const (
	TypeString ValueType = iota
	TypeList
	TypeSet
	TypeHash
)

type Storage struct {
	mu   sync.RWMutex
	data map[string]*Entry
}

type Entry struct {
	Type     ValueType
	ExpireAt time.Time

	Str  []byte
	List [][]byte
	Set  map[string]struct{}
	Hash map[string][]byte
}

func NewStorage() *Storage {
	return &Storage{
		data: make(map[string]*Entry),
	}
}

func (s *Storage) Set(key string, value string, expireAt time.Time) {
	byteValue := resp.ToBytes(value)

	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[key] = &Entry{
		Type:     TypeString,
		ExpireAt: expireAt,
		Str:      byteValue,
	}
}

func (s *Storage) Get(key string) (string, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	entry, ok := s.data[key]

	if !ok || entry.Type != TypeString {
		return resp.NilBulkString(), false
	}
	ok = s.CheckExpired(key)
	if !ok {
		return resp.NilBulkString(), false
	}
	value := resp.ToString(entry.Str)

	return resp.BulkString(value), true
}

func (s *Storage) CheckExpired(key string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	entry, ok := s.data[key]
	if !ok {
		return false
	}
	if entry.ExpireAt.IsZero() {
		return true
	}
	return entry.ExpireAt.After(time.Now())
}

func (s *Storage) RPush(key string, value string, expireAt time.Time) (int, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	v, ok := s.data[key]

	if ok && v.Type != TypeList {
		isExpired := s.CheckExpired(key)
		if isExpired {
			delete(s.data, key)
		} else {
			return 0, false
		}
	}

	bytesValue := resp.ToBytes(value)

	switch ok {
	case false:
		s.data[key] = &Entry{
			Type:     TypeList,
			ExpireAt: expireAt,

			List: [][]byte{bytesValue},
		}
	case true:
		initList := s.data[key].List
		finalList := append(initList, bytesValue)

		s.data[key] = &Entry{
			Type:     TypeList,
			ExpireAt: expireAt,

			List: finalList,
		}
	}
	return len(s.data[key].List), true
}

func (s *Storage) LRange(key string, start int, stop int) ([][]byte, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	val, ok := s.data[key]

	if !ok {
		return [][]byte{}, true
	}
	if val.Type != TypeList {
		return [][]byte{}, false
	}

	if start >= len(val.List) {
		return [][]byte{}, true
	}
	if stop >= len(val.List) {
		stop = len(val.List)
	}

	if start < 0 {
		if -(start) > len(val.List) {
			start = 0
		} else {
			buf := len(val.List) + start
			start = buf
		}
	}
	if stop < 0 {
		if -(stop) > len(val.List) {
			stop = 0
		} else {
			buf := len(val.List) + 1 + stop
			stop = buf
		}
	}
	if start > stop {
		return [][]byte{}, true
	}
	return val.List[start:stop], true
}
