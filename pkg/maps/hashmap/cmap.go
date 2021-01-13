package hashmap

import (
	"github.com/negrel/guds/pkg/elements"
	"github.com/negrel/guds/pkg/maps"
	"sync"
)

// CMap is a thread-safe wrapper around Map.
type CMap struct {
	sync.Mutex
	*Map
}

// NewC instantiate a CMap
func NewC() maps.Map {
	return newCMap()
}

func newCMap() *CMap {
	return &CMap{
		Mutex: sync.Mutex{},
		Map:   newMap(),
	}
}

// Put insert the given value in the map.
func (cm *CMap) Put(key maps.Key, value maps.Value) {
	cm.Lock()
	defer cm.Unlock()
	cm.Map.Put(key, value)
}

// Get return the element mapped to the given key.
// A nil value and false is returned if the map doesn't contain any value
// stored with the given key.
func (cm *CMap) Get(key maps.Key) (maps.Value, bool) {
	cm.Lock()
	defer cm.Unlock()
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
	cm.Lock()
	defer cm.Unlock()
	return cm.Map.Keys()
}

// Empty return if the map is empty.
func (cm *CMap) Empty() bool {
	cm.Lock()
	defer cm.Unlock()
	return cm.Map.Empty()
}

// Size return the number of elements stored in the map.
func (cm *CMap) Size() int {
	cm.Lock()
	defer cm.Unlock()
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
	cm.Lock()
	defer cm.Unlock()
	return cm.Map.Values()
}
