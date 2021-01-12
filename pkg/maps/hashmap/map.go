package hashmap

import (
	"github.com/negrel/guds/pkg/elements"
	"github.com/negrel/guds/pkg/maps"
)

var _ maps.Map = &Map{}

type Map struct {
	m map[maps.Key]maps.Value
}

func New() maps.Map {
	return newMap()
}

func newMap() *Map {
	return &Map{
		m: make(map[maps.Key]maps.Value),
	}
}

func (m *Map) Put(key maps.Key, value maps.Value) {
	m.m[key] = value
}

func (m *Map) Get(key maps.Key) (maps.Value, bool) {
	v, ok := m.m[key]
	return v, ok
}

func (m *Map) Delete(key maps.Key) {
	delete(m.m, key)
}

func (m *Map) Keys() []maps.Key {
	keys := make([]maps.Key, len(m.m))
	i := 0
	for key, _ := range m.m {
		keys[i] = key
		i++
	}

	return keys
}

func (m *Map) Empty() bool {
	return m.Size() == 0
}

func (m *Map) Size() int {
	return len(m.m)
}

func (m *Map) Clear() {
	m.m = make(map[maps.Key]maps.Value)
	//for _, value := range m.Keys() {
	//	m.Delete(value)
	//}
}

func (m *Map) Values() []elements.Element {
	result := make([]elements.Element, m.Size())

	i := 0
	for _, value := range m.m {
		result[i] = value
		i++
	}

	return result
}
