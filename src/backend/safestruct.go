package main

import "sync"

type SafeMap[T any] struct {
	mu   sync.Mutex
	data map[string]T
}

func (s *SafeMap[T]) Add(key string, value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data[key] = value
}

func (s *SafeMap[T]) Get(key string) (T, bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	value, ok := s.data[key]
	return value, ok
}

func (s *SafeMap[T]) Replace(smap map[string]T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.data = smap
}

type SafeArray[T any] struct {
	mu    sync.Mutex
	array []T
}

func (s *SafeArray[T]) Add(value T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.array = append(s.array, value)
}

func (s *SafeArray[T]) Get() []T {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.array
}

func (s *SafeArray[T]) Set(array []T) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.array = array
}
