package handler

import (
	"net"
	"sync"
)

type connStore struct {
	sync.Mutex
	tunnels map[string]*net.Conn
}

var store = connStore{
	tunnels: make(map[string]*net.Conn),
}

func (s *connStore) add(key string, conn *net.Conn) {
	s.Lock()
	defer s.Unlock()
	s.tunnels[key] = conn
}

func (s *connStore) get(key string) *net.Conn {
	s.Lock()
	defer s.Unlock()
	return s.tunnels[key]
}

func (s *connStore) remove(key string) {
	s.Lock()
	defer s.Unlock()
	delete(s.tunnels, key)
}
