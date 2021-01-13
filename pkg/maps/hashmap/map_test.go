package hashmap

import (
	"github.com/negrel/guds/pkg/maps"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap_Set(t *testing.T) {
	m := newMap()

	key := 165
	value := 10
	m.Set(maps.Key(&key), maps.Value(&value))

	actual, ok := m.m[maps.Key(&key)]
	assert.True(t, ok)
	assert.Equal(t, value, *(*int)(actual))
	assert.Equal(t, 1, len(m.m))

	m.Set(maps.Key(&key), maps.Value(&value))
	assert.Equal(t, 1, len(m.m))
}

func TestMap_Get(t *testing.T) {
	m := New()

	key := 178
	value := 10
	m.Set(maps.Key(&key), maps.Value(&value))

	actual, ok := m.Get(maps.Key(&key))
	assert.True(t, ok)
	assert.Equal(t, value, *(*int)(actual))

	actual2, ok := m.Get(maps.Key(&key))
	assert.True(t, ok)
	assert.Equal(t, actual, actual2)
}

func TestMap_Delete(t *testing.T) {
	m := New()

	key := "key"
	value := "value"
	m.Set(maps.Key(&key), maps.Value(&value))

	_, ok := m.Get(maps.Key(&key))
	assert.True(t, ok)

	m.Delete(maps.Key(&key))
	actual, ok := m.Get(maps.Key(&key))
	assert.False(t, ok)
	assert.Equal(t, nil, actual)
}

func TestMap_Keys(t *testing.T) {
	m := New()
	const keyCount = 10
	keys := make([]maps.Key, keyCount)

	// Generates keys
	for i := 0; i < keyCount; i++ {
		key := i
		value := i * 2
		m.Set(maps.Key(&key), maps.Value(&value))

		keys[i] = maps.Key(&key)
	}

	actualKeys := m.Keys()
	assert.Equal(t, keyCount, len(actualKeys))

	keySet := make(map[maps.Key]struct{}, keyCount)
	// Fill the keySet and check that there is no duplicate keys
	for _, actualKey := range actualKeys {
		_, ok := keySet[actualKey]
		assert.False(t, ok)

		keySet[actualKey] = struct{}{}
	}

	// Check that the key set contain the the generated keys
	for _, key := range keys {
		_, ok := keySet[key]
		assert.True(t, ok)
	}
}

func TestMap_Values(t *testing.T) {
	m := New()
	const keyCount = 10
	values := make(map[int]struct{}, keyCount)

	for i := 0; i < keyCount; i++ {
		key := i
		value := i * 2
		m.Set(maps.Key(&key), maps.Value(&value))
		values[value] = struct{}{}
	}

	assert.Equal(t, keyCount, len(m.Values()))

	for _, value := range m.Values() {
		_, ok := values[*(*int)(value)]
		assert.True(t, ok)

		if ok {
			delete(values, *(*int)(value))
		}
	}
}

func TestMap_Empty(t *testing.T) {
	m := New()

	assert.True(t, m.Empty())

	m.Set(nil, nil)
	assert.False(t, m.Empty())

	m.Delete(nil)
	assert.True(t, m.Empty())
}

func TestMap_Size(t *testing.T) {
	m := New()

	m.Set(nil, nil)
	assert.Equal(t, 1, m.Size())

	i := 0
	j := 0
	m.Set(maps.Key(&i), maps.Key(&j))
	assert.Equal(t, 2, m.Size())

	m.Set(maps.Key(&i), maps.Key(&j))
	assert.Equal(t, 2, m.Size())

	m.Delete(maps.Key(&i))
	assert.Equal(t, 1, m.Size())
}

func TestMap_Clear(t *testing.T) {
	m := newMap()

	m.Set(nil, nil)
	assert.Equal(t, 1, m.Size())

	m.Clear()
	m.m = make(map[maps.Key]maps.Value)
	assert.Equal(t, 0, m.Size())
}
