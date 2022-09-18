package memory

import (
	"context"
	"errors"
	"github.com/patrickmn/go-cache"
	web "gitlab.xchch.top/zhangsai/go-101/training/week06/web_v9/v1"
	"sync"
	"time"
)

type Store struct {
	// 如果难以保证同一个id不会被多个goroutine操作, 就加上这个
	mutex      sync.RWMutex
	cache      *cache.Cache
	expiration time.Duration
}

func NewStore(expiration time.Duration) *Store {
	return &Store{
		// defaultExpiration: 默认过期时间,  cleanupInterval 默认清理时间
		cache: cache.New(expiration, time.Second*5),
	}
}

func (s *Store) Generate(ctx context.Context, id string) (web.Session, error) {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	res := &Session{
		id:     id,
		values: make(map[string]string, 16),
	}
	s.cache.Set(id, res, s.expiration)
	return res, nil
}

func (s *Store) Remove(ctx context.Context, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.cache.Delete(id)
	return nil
}

func (s *Store) Get(ctx context.Context, id string) (web.Session, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	res, ok := s.cache.Get(id)
	if ok {
		return res.(*Session), nil
	}
	return nil, errors.New("web: key not found")
}

func (s *Store) Refresh(ctx context.Context, id string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	res, ok := s.cache.Get(id)
	if !ok {
		return errors.New("web: key not found")
	}
	s.cache.Set(id, res, s.expiration)
	return nil
}

type Session struct {
	id     string
	mutex  sync.RWMutex
	values map[string]string
}

func (s *Session) Get(ctx context.Context, key string) (string, error) {
	s.mutex.RLock()
	defer s.mutex.RUnlock()

	val, ok := s.values[key]
	if ok {
		return val, nil
	}
	return "", errors.New("web: key not found")
}

func (s *Session) Set(ctx context.Context, key string, val string) error {
	s.mutex.Lock()
	defer s.mutex.Unlock()

	s.values[key] = val
	return nil
}

func (s *Session) ID() string {
	return s.id
}
