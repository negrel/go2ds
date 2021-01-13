package hashmap

import (
	"github.com/negrel/guds/pkg/elements"
	"github.com/negrel/guds/pkg/maps"
)

var _ maps.Map = &Map{}

// Map is a wrapper around the built-in map type.
type Map struct {
	m map[maps.Key]maps.Value
}

// New instantiate a Map
func New() maps.Map {
	return newMap()
}

func newMap() *Map {
	return &Map{
		m: make(map[maps.Key]maps.Value),
	}
}

// Put insert the given value in the map.
func (m *Map) Set(key maps.Key, value maps.Value) {
	m.m[key] = value
}

// Get return the element mapped to the given key.
// A nil value and false is returned if the map doesn't contain any value
// stored with the given key.
func (m *Map) Get(key maps.Key) (maps.Value, bool) {
	v, ok := m.m[key]
	return v, ok
}

// Delete the given key in the map.
func (m *Map) Delete(key maps.Key) {
	delete(m.m, key)
}

// Keys return all the keys present in the map.
func (m *Map) Keys() []maps.Key {
	keys := make([]maps.Key, len(m.m))
	i := 0
	for key := range m.m {
		keys[i] = key
		i++
	}

	return keys
}

// Empty return if the map is empty.
func (m *Map) Empty() bool {
	return m.Size() == 0
}

// Size return the number of elements stored in the map.
func (m *Map) Size() int {
	return len(m.m)
}

// Clear remove all elements stored in the map.
func (m *Map) Clear() {
	m.m = make(map[maps.Key]maps.Value)
}

// Values return all the values stored in the map.
func (m *Map) Values() []elements.Element {
	result := make([]elements.Element, m.Size())

	i := 0
	for _, value := range m.m {
		result[i] = value
		i++
	}

	return result
}
