package cache

import "time"

type Memory struct {
	cache    map[string]interface{}
	timeouts map[string]*time.Time
}

func (m *Memory) Read(key string) (interface{}, error) {
	m.init()
	if m.timeouts[key] == nil || m.cache[key] == nil {
		return nil, nil
	}
	if time.Now().After(*m.timeouts[key]) {
		m.cache[key] = nil
		return nil, nil
	}
	return m.cache[key], nil
}

func (m *Memory) init() {
	if m.cache == nil {
		m.cache = map[string]interface{}{}
		m.timeouts = map[string]*time.Time{}
	}
}

func (m *Memory) Write(key string, obj interface{}, timeoutHours int32) error {
	m.init()
	m.cache[key] = obj
	timeout := time.Now().Add(time.Duration(timeoutHours) * time.Second)
	m.timeouts[key] = &timeout
	return nil
}
