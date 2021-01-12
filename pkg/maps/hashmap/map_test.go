package hashmap

import (
	"github.com/negrel/guds/pkg/maps"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMap_Put(t *testing.T) {
	m := newMap()

	key := 165
	value := 10
	m.Put(maps.Key(&key), maps.Value(&value))

	actual, ok := m.m[maps.Key(&key)]
	assert.True(t, ok)
	assert.Equal(t, *(*int)(actual), value)
}

func TestMap_Get(t *testing.T) {
	m := New()

	key := 178
	value := 10
	m.Put(maps.Key(&key), maps.Value(&value))

	actual, ok := m.Get(maps.Key(&key))
	assert.True(t, ok)
	assert.Equal(t, *(*int)(actual), value)
}

func TestMap_Delete(t *testing.T) {
	m := New()

	key := "key"
	value := "value"
	m.Put(maps.Key(&key), maps.Value(&value))

	_, ok := m.Get(maps.Key(&key))
	assert.True(t, ok)

	m.Delete(maps.Key(&key))
	_, ok = m.Get(maps.Key(&key))
	assert.False(t, ok)
}

func TestMap_Keys(t *testing.T) {
	m := New()
	const keyCount = 10
	keys := make([]maps.Key, keyCount)

	// Generates keys
	for i := 0; i < keyCount; i++ {
		key := i
		value := i * 2
		m.Put(maps.Key(&key), maps.Value(&value))

		keys[i] = maps.Key(&key)
	}

	actualKeys := m.Keys()
	assert.Equal(t, len(actualKeys), keyCount)

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
