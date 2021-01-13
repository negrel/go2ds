package hashmap

import (
	"github.com/negrel/guds/pkg/elements"
	"github.com/negrel/guds/pkg/maps"
	"sync"
)

// CMap is a thread-safe wrapper around the hash Map.
type CMap struct {
	sync.Mutex
	*Map
}

func NewC() maps.Map {
	return newCMap()
}

func newCMap() *CMap {
	return &CMap{
		Mutex: sync.Mutex{},
		Map:   newMap(),
	}
}

func (cm *CMap) Put(key maps.Key, value maps.Value) {
	cm.Lock()
	defer cm.Unlock()
	cm.Map.Put(key, value)
}

func (cm *CMap) Get(key maps.Key) (maps.Value, bool) {
	cm.Lock()
	defer cm.Unlock()
	return cm.Map.Get(key)
}

func (cm *CMap) Delete(key maps.Key) {
	cm.Lock()
	defer cm.Unlock()
	cm.Map.Delete(key)
}

func (cm *CMap) Keys() []maps.Key {
	cm.Lock()
	defer cm.Unlock()
	return cm.Map.Keys()
}

func (cm *CMap) Empty() bool {
	cm.Lock()
	defer cm.Unlock()
	return cm.Map.Empty()
}

func (cm *CMap) Size() int {
	cm.Lock()
	defer cm.Unlock()
	return cm.Map.Size()
}

func (cm *CMap) Clear() {
	cm.Lock()
	defer cm.Unlock()
	cm.Map.Clear()
}

func (cm *CMap) Values() []elements.Element {
	cm.Lock()
	defer cm.Unlock()
	return cm.Map.Values()
}
