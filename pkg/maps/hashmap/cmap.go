package hashmap

import (
	"github.com/negrel/guds/pkg/elements"
	"github.com/negrel/guds/pkg/maps"
	"sync"
)

// CMap is a thread-safe wrapper around Map.
type CMap struct {
	sync.RWMutex
	*Map
}

// NewC instantiate a CMap
func NewC() maps.Map {
	return newCMap()
}

func newCMap() *CMap {
	return &CMap{
		RWMutex: sync.RWMutex{},
		Map:     newMap(),
	}
}

// Set insert the given value in the map.
func (cm *CMap) Set(key maps.Key, value maps.Value) {
	cm.Lock()
	defer cm.Unlock()
	cm.Map.Set(key, value)
}

// Get return the element mapped to the given key.
// A nil value and false is returned if the map doesn't contain any value
// stored with the given key.
func (cm *CMap) Get(key maps.Key) (maps.Value, bool) {
	cm.RLock()
	defer cm.RUnlock()
	return cm.Map.Get(key)
}

// Delete the given key in the map.
func (cm *CMap) Delete(key maps.Key) {
	cm.Lock()
	defer cm.Unlock()
	cm.Map.Delete(key)
}

// Keys return all the keys present in the map.
func (cm *CMap) Keys() []maps.Key {
	cm.RLock()
	defer cm.RUnlock()
	return cm.Map.Keys()
}

// Empty return if the map is empty.
func (cm *CMap) Empty() bool {
	return cm.Size() == 0
}

// Size return the number of elements stored in the map.
func (cm *CMap) Size() int {
	cm.RLock()
	defer cm.RUnlock()
	return cm.Map.Size()
}

// Clear remove all elements stored in the map.
func (cm *CMap) Clear() {
	cm.Lock()
	defer cm.Unlock()
	cm.Map.Clear()
}

// Values return all the values stored in the map.
func (cm *CMap) Values() []elements.Element {
	cm.RLock()
	defer cm.RUnlock()
	return cm.Map.Values()
}
