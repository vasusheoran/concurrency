package safemap

import "sync"

// SafeMap acts as visited set
type SafeMap interface {
	Exist(key interface{}) bool
	Add(key interface{})
}

type safeMap struct {
	data  map[interface{}]interface{}
	mutex sync.RWMutex
}

func (sm *safeMap) Add(key interface{}) {
	sm.mutex.Lock()
	defer sm.mutex.Unlock()

	sm.data[key] = true
}

func (sm *safeMap) Exist(key interface{}) bool {
	sm.mutex.RLock()
	defer sm.mutex.RUnlock()

	_, ok := sm.data[key]
	return ok
}

func NewSafeMap() SafeMap {
	return &safeMap{
		data:  map[interface{}]interface{}{},
		mutex: sync.RWMutex{},
	}
}
