package maps

import (
	"github.com/negrel/guds/pkg/containers"
	"github.com/negrel/guds/pkg/elements"
)

// Key define the data to which a Value is associated in a Map.
type Key = elements.Element

// Value define any element contained in a Map.
type Value = elements.Element

// Map define the common method implemented by all Map.
type Map interface {
	containers.Container

	Set(Key, Value)
	Get(Key) (Value, bool)
	Delete(Key)
	Keys() []Key
}
