package ws

import "sync"

type RwMap struct {
	sync.RWMutex
	m map[string]*Client
}

func NewRwMap() *RwMap {
	return &RwMap{
		m: make(map[string]*Client),
	}
}

func (m *RwMap) Get(k string) (*Client, bool) {
	m.RLock()
	defer m.RUnlock()
	v, existed := m.m[k]
	return v, existed
}

func (m *RwMap) Set(k string, v *Client) {
	m.Lock()
	defer m.Unlock()
	m.m[k] = v
}

func (m *RwMap) Delete(k string) {
	m.Lock()
	defer m.Unlock()
	delete(m.m, k)
}

func (m *RwMap) Len() int {
	m.RLock()
	defer m.RUnlock()
	return len(m.m)
}

func (m *RwMap) Each(f func(k string, v *Client) bool) {
	m.RLock()
	defer m.RUnlock()

	for k, v := range m.m {
		if !f(k, v) {
			return
		}
	}
}
