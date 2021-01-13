package maps

import (
	"github.com/negrel/guds/pkg/containers"
	"github.com/negrel/guds/pkg/elements"
)

type Key = elements.Element
type Value = elements.Element

type Map interface {
	containers.Container

	Set(Key, Value)
	Get(Key) (Value, bool)
	Delete(Key)
	Keys() []Key
}
