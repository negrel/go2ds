package hashmap

import (
	"github.com/negrel/guds/pkg/elements"
	"github.com/negrel/guds/pkg/maps"
	"sync"
)

// Sync is a thread-safe wrapper around Map.
type Sync struct {
	sync.RWMutex
	*Map
}

// NewSync instantiate a Sync
func NewSync() maps.Map {
	return newCMap()
}

func newCMap() *Sync {
	return &Sync{
		RWMutex: sync.RWMutex{},
		Map:     newMap(),
	}
}

// Set insert the given value in the map.
func (s *Sync) Set(key maps.Key, value maps.Value) {
	s.Lock()
	defer s.Unlock()
	s.Map.Set(key, value)
}

// Get return the element mapped to the given key.
// A nil value and false is returned if the map doesn't contain any value
// stored with the given key.
func (s *Sync) Get(key maps.Key) (maps.Value, bool) {
	s.RLock()
	defer s.RUnlock()
	return s.Map.Get(key)
}

// Delete the given key in the map.
func (s *Sync) Delete(key maps.Key) {
	s.Lock()
	defer s.Unlock()
	s.Map.Delete(key)
}

// Keys return all the keys present in the map.
func (s *Sync) Keys() []maps.Key {
	s.RLock()
	defer s.RUnlock()
	return s.Map.Keys()
}

// Empty return if the map is empty.
func (s *Sync) Empty() bool {
	return s.Size() == 0
}

// Size return the number of elements stored in the map.
func (s *Sync) Size() int {
	s.RLock()
	defer s.RUnlock()
	return s.Map.Size()
}

// Clear remove all elements stored in the map.
func (s *Sync) Clear() {
	s.Lock()
	defer s.Unlock()
	s.Map.Clear()
}

// Values return all the values stored in the map.
func (s *Sync) Values() []elements.Element {
	s.RLock()
	defer s.RUnlock()
	return s.Map.Values()
}
